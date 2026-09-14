package encoder

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"net"
)

// TrojanRecord handles Trojan protocol encoding/decoding.
// Wire format: [AuthHash:56][Command:1][Port:2][AddrType:1][Addr:...][Payload]
type TrojanRecord struct {
	AuthHash [56]byte // SHA256 of token
	Command  byte
	Port     uint16
	AddrType byte
	Address  string
	Payload  []byte
}

const (
	TrojanCmdTCP byte = 0x01
	TrojanCmdUDP byte = 0x02
)

// ComputeTrojanAuth derives the auth hash from a token (full SHA256 = 32 bytes, padded to 56).
func ComputeTrojanAuth(token string) [56]byte {
	hash := sha256.Sum256([]byte(token))
	var result [56]byte
	// Copy 32 bytes of hash, rest is zero
	copy(result[:], hash[:])
	return result
}

// BuildTrojanRequest builds a Trojan request for dialing.
func BuildTrojanRequest(authHash [56]byte, addr string, port uint16, cmd byte) ([]byte, error) {
	req := make([]byte, 0, 256)
	req = append(req, authHash[:]...)
	req = append(req, cmd)
	req = append(req, byte(port>>8), byte(port&0xFF))

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
		if len(addr) > 255 {
			return nil, fmt.Errorf("trojan domain too long: %s", addr)
		}
		req = append(req, AddrDomain, byte(len(addr)))
		req = append(req, []byte(addr)...)
	}

	return req, nil
}

// ParseTrojanRequest parses a Trojan request from bytes.
func ParseTrojanRequest(data []byte) (*TrojanRecord, int, error) {
	if len(data) < 60 {
		return nil, 0, fmt.Errorf("trojan request too short: %d < 60", len(data))
	}

	rec := &TrojanRecord{}
	copy(rec.AuthHash[:], data[:56])
	rec.Command = data[56]
	rec.Port = binary.BigEndian.Uint16(data[57:59])

	idx := 59
	rec.AddrType = data[idx]
	idx++

	switch rec.AddrType {
	case AddrIPv4:
		if idx+4 > len(data) {
			return nil, 0, fmt.Errorf("trojan ipv4 truncated")
		}
		rec.Address = net.IP(data[idx : idx+4]).String()
		idx += 4

	case AddrIPv6:
		if idx+16 > len(data) {
			return nil, 0, fmt.Errorf("trojan ipv6 truncated")
		}
		rec.Address = net.IP(data[idx : idx+16]).String()
		idx += 16

	case AddrDomain:
		if idx >= len(data) {
			return nil, 0, fmt.Errorf("trojan domain length missing")
		}
		domainLen := int(data[idx])
		idx++
		if idx+domainLen > len(data) {
			return nil, 0, fmt.Errorf("trojan domain truncated")
		}
		rec.Address = string(data[idx : idx+domainLen])
		idx += domainLen

	default:
		return nil, 0, fmt.Errorf("trojan unknown address type: %d", rec.AddrType)
	}

	rec.Payload = data[idx:]
	return rec, idx, nil
}
