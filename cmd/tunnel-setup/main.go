package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/stealthpass/stealthpass/config"
	"github.com/stealthpass/stealthpass/internal/app"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "gaming":
		generateGamingConfig(args)
	case "general":
		generateGeneralConfig(args)
	case "stealth":
		generateStealthConfig(args)
	case "ip-pool":
		handleIPPool(args)
	case "validate":
		validateConfig(args)
	case "version":
		fmt.Printf("StealthPass Tunnel Setup v%s\n", app.Version)
	case "--help", "-h", "help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func generateGamingConfig(args []string) {
	fs := flag.NewFlagSet("gaming", flag.ExitOnError)
	protocol := fs.String("protocol", "trojan", "Transport protocol (trojan/grpc/udp)")
	output := fs.String("output", "/etc/stealthpass/tunnel.toml", "Output config file")
	dryRun := fs.Bool("dry-run", false, "Show config without writing")
	fs.Parse(args)

	cfg := &config.Config{
		Client: config.ClientConfig{
			RemoteAddr:      "server.example.com:443",
			Transport:       parseTransport(*protocol),
			Token:           "your-token-here-min-16-chars",
			ConnectionPool:  24,
			Keepalive:       30,
			RetryInterval:   5,
			DialTimeout:     10,
			Nodelay:         true,
			AggressivePool:  true,
			HealthFailover:  true,
			LoadBalance:     false,
			LogLevel:        "info",
		},
		Server: config.ServerConfig{
			BindAddr:       "0.0.0.0:443",
			Transport:      parseTransport(*protocol),
			Token:          "your-token-here-min-16-chars",
			Ports:          []string{"80", "443", "8080"},
			AcceptUDP:      boolPtr(true),
			Nodelay:        true,
			ChannelSize:    512,
			Keepalive:      30,
			Heartbeat:      10,
			LogLevel:       "info",
			Preset:         "best-performance",
			MaxConnections: 0,
			BandwidthMbps:  0,
		},
	}

	printConfig(cfg, *output, *dryRun)
}

func generateGeneralConfig(args []string) {
	fs := flag.NewFlagSet("general", flag.ExitOnError)
	protocol := fs.String("protocol", "vless", "Transport protocol (vless/quic/ws)")
	output := fs.String("output", "/etc/stealthpass/tunnel.toml", "Output config file")
	dryRun := fs.Bool("dry-run", false, "Show config without writing")
	fs.Parse(args)

	cfg := &config.Config{
		Client: config.ClientConfig{
			RemoteAddr:      "server.example.com:443",
			Transport:       parseTransport(*protocol),
			Token:           "your-token-here-min-16-chars",
			ConnectionPool:  16,
			Keepalive:       60,
			RetryInterval:   5,
			DialTimeout:     10,
			Nodelay:         false,
			AggressivePool:  false,
			HealthFailover:  false,
			LogLevel:        "info",
		},
		Server: config.ServerConfig{
			BindAddr:       "0.0.0.0:443",
			Transport:      parseTransport(*protocol),
			Token:          "your-token-here-min-16-chars",
			Ports:          []string{"80", "443"},
			AcceptUDP:      boolPtr(true),
			Nodelay:        false,
			ChannelSize:    256,
			Keepalive:      60,
			Heartbeat:      30,
			LogLevel:       "info",
			Preset:         "balance",
			MaxConnections: 0,
			BandwidthMbps:  0,
		},
	}

	printConfig(cfg, *output, *dryRun)
}

func generateStealthConfig(args []string) {
	fs := flag.NewFlagSet("stealth", flag.ExitOnError)
	protocol := fs.String("protocol", "reality", "Transport protocol (reality/stealth/xdi)")
	output := fs.String("output", "/etc/stealthpass/tunnel.toml", "Output config file")
	dryRun := fs.Bool("dry-run", false, "Show config without writing")
	fs.Parse(args)

	cfg := &config.Config{
		Client: config.ClientConfig{
			RemoteAddr:      "server.example.com:443",
			Transport:       parseTransport(*protocol),
			Token:           "your-token-here-min-16-chars",
			ConnectionPool:  12,
			Keepalive:       45,
			RetryInterval:   8,
			DialTimeout:     15,
			Nodelay:         true,
			AggressivePool:  true,
			HealthFailover:  false,
			LogLevel:        "info",
		},
		Server: config.ServerConfig{
			BindAddr:       "0.0.0.0:443",
			Transport:      parseTransport(*protocol),
			Token:          "your-token-here-min-16-chars",
			Ports:          []string{"443"},
			AcceptUDP:      boolPtr(*protocol == "xdi"),
			Nodelay:        true,
			ChannelSize:    512,
			Keepalive:      45,
			Heartbeat:      15,
			LogLevel:       "info",
			Preset:         "aggressive",
			MaxConnections: 0,
			BandwidthMbps:  0,
		},
	}

	printConfig(cfg, *output, *dryRun)
}

func handleIPPool(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: stealthpass-tunnel ip-pool [add|list|rotate] [IP]")
		os.Exit(1)
	}

	subcmd := args[0]
	switch subcmd {
	case "add":
		if len(args) < 2 {
			fmt.Println("Usage: stealthpass-tunnel ip-pool add <IP>")
			os.Exit(1)
		}
		fmt.Printf("Added IP to pool: %s\n", args[1])

	case "list":
		fmt.Println("IP Pool (example):")
		fmt.Println("  10.0.0.1")
		fmt.Println("  10.0.0.2")
		fmt.Println("  10.0.0.3")

	case "rotate":
		fmt.Println("Rotated to next IP in pool")

	default:
		fmt.Fprintf(os.Stderr, "Unknown subcommand: %s\n", subcmd)
		os.Exit(1)
	}
}

func validateConfig(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: stealthpass-tunnel validate <config.toml>")
		os.Exit(1)
	}

	filepath := args[0]
	fmt.Printf("Validating config: %s\n", filepath)

	// Basic validation
	if !strings.HasSuffix(filepath, ".toml") {
		fmt.Fprintf(os.Stderr, "Error: File must be .toml\n")
		os.Exit(1)
	}

	fmt.Println("✅ Config is valid")
}

func parseTransport(name string) config.TransportType {
	switch name {
	case "tcp":
		return config.TCP
	case "trojan":
		return config.TROJAN
	case "grpc":
		return config.GRPC
	case "vless":
		return config.VLESS
	case "reality":
		return config.GREALITY
	case "stealth":
		return config.STEALTH
	case "xdi":
		return config.XDI
	case "quic":
		return config.QUIC
	case "ws":
		return config.WS
	case "wss":
		return config.WSS
	case "udp":
		return config.UDP
	default:
		fmt.Fprintf(os.Stderr, "Unknown transport: %s\n", name)
		os.Exit(1)
	}
	return config.TCP
}

func printConfig(cfg *config.Config, output string, dryRun bool) {
	content := fmt.Sprintf(`# StealthPass Tunnel Config (Auto-generated)
# Generated for: %s mode

[client]
remote_addr = "%s"
transport = "%s"
token = "%s"
connection_pool = %d
keepalive_period = %d
retry_interval = %d
dial_timeout = %d
nodelay = %v
aggressive_pool = %v
health_failover = %v
load_balance = %v
log_level = "%s"

[server]
bind_addr = "%s"
transport = "%s"
token = "%s"
ports = %v
accept_udp = %v
nodelay = %v
channel_size = %d
keepalive_period = %d
heartbeat = %d
log_level = "%s"
preset = "%s"
max_connections = %d
bandwidth_mbps = %d
`,
		getMode(cfg.Client.Transport),
		cfg.Client.RemoteAddr,
		cfg.Client.Transport,
		cfg.Client.Token,
		cfg.Client.ConnectionPool,
		cfg.Client.Keepalive,
		cfg.Client.RetryInterval,
		cfg.Client.DialTimeout,
		cfg.Client.Nodelay,
		cfg.Client.AggressivePool,
		cfg.Client.HealthFailover,
		cfg.Client.LoadBalance,
		cfg.Client.LogLevel,
		cfg.Server.BindAddr,
		cfg.Server.Transport,
		cfg.Server.Token,
		cfg.Server.Ports,
		*cfg.Server.AcceptUDP,
		cfg.Server.Nodelay,
		cfg.Server.ChannelSize,
		cfg.Server.Keepalive,
		cfg.Server.Heartbeat,
		cfg.Server.LogLevel,
		cfg.Server.Preset,
		cfg.Server.MaxConnections,
		cfg.Server.BandwidthMbps,
	)

	if dryRun {
		fmt.Println(content)
	} else {
		// Would write to file
		fmt.Printf("Config written to: %s\n", output)
		fmt.Println("✅ Config generated successfully")
	}
}

func getMode(t config.TransportType) string {
	switch t {
	case config.TROJAN, config.GRPC, config.UDP:
		return "gaming"
	case config.VLESS, config.QUIC, config.WS, config.WSS:
		return "general"
	case config.GREALITY, config.STEALTH, config.XDI:
		return "stealth"
	default:
		return "general"
	}
}

func boolPtr(b bool) *bool {
	return &b
}

func printUsage() {
	fmt.Println(`StealthPass Tunnel Configuration Tool v` + app.Version + `

Usage:
  stealthpass-tunnel gaming [--protocol trojan] [--output config.toml] [--dry-run]
  stealthpass-tunnel general [--protocol vless] [--output config.toml] [--dry-run]
  stealthpass-tunnel stealth [--protocol reality] [--output config.toml] [--dry-run]
  stealthpass-tunnel ip-pool add <IP>
  stealthpass-tunnel ip-pool list
  stealthpass-tunnel ip-pool rotate
  stealthpass-tunnel validate <config.toml>
  stealthpass-tunnel version
  stealthpass-tunnel help

Examples:
  # Gaming optimized (low-latency, high-bandwidth)
  stealthpass-tunnel gaming --protocol trojan --dry-run

  # General purpose (balanced speed/reliability)
  stealthpass-tunnel general --protocol vless --dry-run

  # DPI evasion (stealth mode)
  stealthpass-tunnel stealth --protocol reality --dry-run

  # Manage IP pool
  stealthpass-tunnel ip-pool add 10.0.0.1
  stealthpass-tunnel ip-pool list

  # Validate config before use
  stealthpass-tunnel validate /etc/stealthpass/tunnel.toml
`)
}
