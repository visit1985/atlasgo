package credentials

import "errors"

var (
    ErrNoValidProvidersFoundInChain = errors.New(`no valid providers in chain`)
)

type ChainProvider struct {
    Providers     []Provider
    curr          Provider
    VerboseErrors bool
}

func NewChainCredentials(providers []Provider) *Credentials {
    return NewCredentials(&ChainProvider{
        Providers: append([]Provider{}, providers...),
    })
}

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

func (c *ChainProvider) IsExpired() bool {
    if c.curr != nil {
        return c.curr.IsExpired()
    }

    return true
}
