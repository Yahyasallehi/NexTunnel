package config

// Spoof-carrier mode resolution, kept in one place so every consumer (the
// server, the client, and the startup validation) agrees on what "relay mode"
// means and where it forwards — and so the legacy spoof_pipe keys keep working
// without that alias logic being scattered across the dispatch sites.

// Relay mode — a bare datagram relay to a local UDP socket, instead of a
// tunnel — was a shape of the reverse spoof transport, and went with it. The
// direct tunnel carries a whole private network, which is what the relay was
// reached for: an inner transport that brings its own reliability (WireGuard,
// most often) is routed over the tunnel rather than piped through it.

// SNISpoofConfig holds the SNI rotation settings for the spoof transport.
// When sni_rotate is true, each new connection picks a different domain from
// sni_domains (or the built-in Iranian list when that is empty), so the tunnel
// does not sit on one SNI a filter can learn to block.
type SNISpoofConfig struct {
	// SNIRotate enables automatic SNI rotation per connection.
	SNIRotate bool `toml:"sni_rotate"`
	// SNIRandom picks domains at random instead of cycling in order.
	SNIRandom bool `toml:"sni_random"`
	// SNIDomains is the list of domains to rotate through.
	// Empty means use the built-in Iranian domain list.
	SNIDomains []string `toml:"sni_domains"`
}
