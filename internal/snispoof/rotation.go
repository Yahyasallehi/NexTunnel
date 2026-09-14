package snispoof

import (
	"math/rand"
	"sync"
	"time"
)

// IranianDomains is a curated list of domains that pass through Iranian
// filters. These are well-known domestic domains a middlebox is unlikely to block.
var IranianDomains = []string{
	"mci.ir",
	"irancell.ir",
	"shaparak.ir",
	"aparat.com",
	"digikala.com",
	"snapp.ir",
	"tap.ir",
	"cafebazaar.ir",
	"divar.ir",
	"hamrahcard.ir",
	"telewebion.com",
	"filimo.com",
	"namava.ir",
	"eitaa.com",
	"rubika.ir",
}

// Rotator picks a domain from a list on every call.
// In random mode each pick is random; otherwise it cycles in order.
type Rotator struct {
	mu      sync.Mutex
	domains []string
	index   int
	random  bool
	rng     *rand.Rand
}

// NewRotator builds a Rotator from the given list.
// If domains is empty, IranianDomains is used.
func NewRotator(domains []string, random bool) *Rotator {
	if len(domains) == 0 {
		domains = IranianDomains
	}
	return &Rotator{
		domains: domains,
		random:  random,
		rng:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Next returns the next domain to use as SNI.
func (r *Rotator) Next() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	if len(r.domains) == 1 {
		return r.domains[0]
	}
	if r.random {
		return r.domains[r.rng.Intn(len(r.domains))]
	}
	d := r.domains[r.index]
	r.index = (r.index + 1) % len(r.domains)
	return d
}

// DomainOrRotate returns the configured domain if set,
// otherwise picks the next domain from the rotator.
func DomainOrRotate(domain string, r *Rotator) string {
	if domain != "" {
		return domain
	}
	if r != nil {
		return r.Next()
	}
	return DefaultDomain
}
