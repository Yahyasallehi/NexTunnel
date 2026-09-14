# StealthPass Implementation Guide

## Phase 2: New Transports

### 1. gRPC Transport
**Files to create:**
- `internal/client/transport/grpc.go` (500 lines)
- `internal/server/transport/grpc.go` (500 lines)

**Pattern to follow:** Copy from `internal/client/transport/tcp.go` and `internal/server/transport/quic.go`

**Key differences:**
- Use `google.golang.org/grpc` bidirectional streams instead of raw TCP/QUIC
- Token in gRPC metadata header, not handshake
- Add dependency: `require github.com/google.golang.org/grpc v1.60.0`

**Add to `internal/client/client.go`:**
```go
case config.GRPC:
    grpcConfig := &transport.GrpcConfig{...}
    grpcClient := transport.NewGrpcClient(c.ctx, grpcConfig, c.logger)
    go grpcClient.Start()
```

**Add to `internal/server/server.go`:**
```go
case config.GRPC:
    grpcConfig := &transport.GrpcConfig{...}
    grpcServer := transport.NewGrpcServer(s.ctx, grpcConfig, s.logger)
    go grpcServer.Start()
```

---

### 2. Reality TLS Transport
**Files to create:**
- `internal/snispoof/reality.go` (200 lines)

**Pattern:** Extend from existing uTLS usage

**Key logic:**
- Use `github.com/refraction-networking/utls` with Chrome 120 fingerprint
- Build ClientHello that exactly matches browser
- Server validates SNI against `RealityDomain` from config

**Add to `internal/client/client.go`:**
```go
case config.GREALITY:
    realityConfig := &transport.RealityConfig{
        Domain: c.config.RealityDomain,
        Rotator: snispoof.NewRotator(c.config.SNIDomains, c.config.SNIRandom),
    }
    realityClient := transport.NewRealityClient(c.ctx, realityConfig, c.logger)
    go realityClient.Start()
```

---

### 3. VLESS Transport
**Files to create:**
- `internal/utils/encoder/vless.go` (300 lines) - Protocol encoder
- `internal/client/transport/vless.go` (400 lines)
- `internal/server/transport/vless.go` (400 lines)

**VLESS wire format (simplified):**
```
[Version:1][UserHash:16][CommandType:1][IPv4/Domain:...][Port:2][Payload]
```

**Add to dispatch cases.**

---

### 4. Trojan Transport
**Files to create:**
- `internal/utils/encoder/trojan.go` (200 lines) - Protocol encoder
- `internal/client/transport/trojan.go` (400 lines)
- `internal/server/transport/trojan.go` (400 lines)

**Trojan wire format:**
```
[AuthHash:56][Command:1][Port:2][Type:1][Domain/IPv4:...][Payload]
```

---

## Phase 3: Configuration Tool (`stealthpass-tunnel`)

### Directory Structure
```
cmd/tunnel-setup/
├── main.go                    (CLI entry point)
└── handler.go                 (subcommands)

internal/tunnel-setup/
├── setup.go                   (mode selection)
├── gaming.go                  (gaming presets)
├── general.go                 (general presets)
├── stealth.go                 (stealth/spoofing presets)
├── ippool.go                  (manual IP management)
├── spoof_profiles.go          (UDP/TCP/ICMP profiles)
└── validator.go               (validation)
```

### CLI Commands
```bash
stealthpass-tunnel                    # Interactive mode
stealthpass-tunnel gaming            # Gaming preset
stealthpass-tunnel general           # General preset
stealthpass-tunnel stealth           # Stealth preset
stealthpass-tunnel ip-pool add IP    # Add IP to rotation pool
stealthpass-tunnel ip-pool list      # List IPs
stealthpass-tunnel validate config.toml  # Validate config
```

### Key Validation Rules
- Transport type exists in config constants
- Token >= 16 characters
- Ports in 1-65535 range
- SNI domains not empty if rotation enabled
- TLS certs exist if required for transport
- All preset fields compatible with transport

---

## Build & Test Checklist

- [ ] `go mod tidy` (resolve new dependencies)
- [ ] `GOOS=linux GOARCH=amd64 go build -o stealthpass .`
- [ ] Build succeeds with no errors
- [ ] `./stealthpass -v` outputs version
- [ ] Sample config loads without errors
- [ ] All 15 transports dispatch correctly
- [ ] `stealthpass-tunnel` binary builds
- [ ] CLI help text works
- [ ] Config generation works
- [ ] IP pool commands work
- [ ] Validation catches errors

---

## Token-Efficient Implementation Order

1. **Start with VLESS** (simplest protocol, Xray wire-format documented)
2. **Then Trojan** (similar to VLESS, well-documented)
3. **Then gRPC** (leverage grpc library, less code needed)
4. **Then Reality** (build on uTLS, most flexible)
5. **Finally tool** (integrates all transports)

Each transport should follow the existing pattern:
- Config struct in config.go ✅
- Client transport handler
- Server transport handler
- Add case to client.go dispatch
- Add case to server.go dispatch
- Test with sample config

---

## Files Already Modified
- ✅ `config/config.go` - Added TransportType constants and config structs
- ✅ `internal/app/app.go` - Version bumped to v2.0.0
- ✅ `go.mod` - Module renamed to stealthpass/stealthpass

## Files Ready to Modify
- `internal/client/client.go` - Add transport dispatch cases
- `internal/server/server.go` - Add transport dispatch cases

---

## Debugging Tips

**Build fails with "undefined reference":**
- Check that transport case in client.go/server.go matches TransportType constant
- Verify NewXxxClient() and NewXxxServer() functions exist

**Config loads but tunnel doesn't start:**
- Check logger output for transport-specific errors
- Verify TLS certs exist if required
- Verify token format matches protocol expectations

**Protocol handshake fails:**
- Most common: token mismatch (hash calculations)
- Check endianness of binary fields
- Verify both ends using same preset
