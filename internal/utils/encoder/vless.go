package encoder

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"net"
)

// VLESSRecord handles VLESS protocol encoding/decoding.
// Wire format: [UserHash:16][Command:1][Port:2][AddrType:1][Addr:...][Payload]
type VLESSRecord struct {
	UserHash  [16]byte
	Command   byte
	Port      uint16
	AddrType  byte
	Address   string
	Payload   []byte
}

const (
	CmdTCP   = 0x01
	CmdUDP   = 0x02
	AddrIPv4 = 0x01
	AddrIPv6 = 0x02
	AddrDomain = 0x03
)

// ComputeUserHash derives the user hash from a token (first 16 bytes of SHA256).
func ComputeUserHash(token string) [16]byte {
	hash := sha256.Sum256([]byte(token))
	var result [16]byte
	copy(result[:], hash[:16])
	return result
}

// BuildRequest builds a VLESS request for dialing a remote address.
func BuildRequest(userHash [16]byte, addr string, port uint16, cmd byte) ([]byte, error) {
	req := make([]byte, 0, 256)
	req = append(req, userHash[:]...)
	req = append(req, cmd)
	req = append(req, byte(port>>8), byte(port&0xFF))

	// Parse address
	ip := net.ParseIP(addr)
	if ip != nil {
		if ip4 := ip.To4(); ip4 != nil {
			req = append(req, AddrIPv4)
			req = append(req, ip4...)
		} else {
			req = append(req, AddrIPv6)
			req = append(req, ip...)
		}
	} else {
		// Domain name
		if len(addr) > 255 {
			return nil, fmt.Errorf("domain too long: %s", addr)
		}
		req = append(req, AddrDomain, byte(len(addr)))
		req = append(req, []byte(addr)...)
	}

	return req, nil
}

// ParseRequest parses a VLESS request from bytes.
func ParseRequest(data []byte) (*VLESSRecord, int, error) {
	if len(data) < 19 {
		return nil, 0, fmt.Errorf("vless request too short: %d < 19", len(data))
	}

	rec := &VLESSRecord{}
	copy(rec.UserHash[:], data[:16])
	rec.Command = data[16]
	rec.Port = binary.BigEndian.Uint16(data[17:19])

	idx := 19
	rec.AddrType = data[idx]
	idx++

	switch rec.AddrType {
	case AddrIPv4:
		if idx+4 > len(data) {
			return nil, 0, fmt.Errorf("vless ipv4 truncated")
		}
		rec.Address = net.IP(data[idx : idx+4]).String()
		idx += 4

	case AddrIPv6:
		if idx+16 > len(data) {
			return nil, 0, fmt.Errorf("vless ipv6 truncated")
		}
		rec.Address = net.IP(data[idx : idx+16]).String()
		idx += 16

	case AddrDomain:
		if idx >= len(data) {
			return nil, 0, fmt.Errorf("vless domain length missing")
		}
		domainLen := int(data[idx])
		idx++
		if idx+domainLen > len(data) {
			return nil, 0, fmt.Errorf("vless domain truncated")
		}
		rec.Address = string(data[idx : idx+domainLen])
		idx += domainLen

	default:
		return nil, 0, fmt.Errorf("vless unknown address type: %d", rec.AddrType)
	}

	rec.Payload = data[idx:]
	return rec, idx, nil
}
