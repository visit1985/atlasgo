package credentials

import "errors"

var (
    ErrNoValidProvidersFoundInChain = errors.New(`no valid providers in chain`)
)

// A ChainProvider will search for a provider which returns credentials.
type ChainProvider struct {
    Providers     []Provider
    curr          Provider
    VerboseErrors bool
}

// NewChainCredentials returns a pointer to a new Credentials object wrapping a chain of providers.
func NewChainCredentials(providers []Provider) *Credentials {
    return NewCredentials(&ChainProvider{
        Providers: append([]Provider{}, providers...),
    })
}

// Retrieve returns the credentials value or error if no provider returned without error.
func (c *ChainProvider) Retrieve() (Value, error) {
    var errs []error
    for _, p := range c.Providers {
        creds, err := p.Retrieve()
        if err == nil {
            c.curr = p
            return creds, nil
        }
        errs = append(errs, err)
    }
    c.curr = nil

    err := ErrNoValidProvidersFoundInChain
    return Value{}, err
}

// IsExpired will returned the expired state of the currently cached provider if there is one.
func (c *ChainProvider) IsExpired() bool {
    if c.curr != nil {
        return c.curr.IsExpired()
    }

    return true
}
