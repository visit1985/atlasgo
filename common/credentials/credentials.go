package credentials

import (
    "sync"
    "time"
)


type Value struct {
    Username     string
    AccessKey    string
    ProviderName string
}

type Provider interface {
    Retrieve()  (Value, error)
    IsExpired() bool
}

type ErrorProvider struct {
    Err          error
    ProviderName string
}

func (p ErrorProvider) Retrieve() (Value, error) {
    return Value{ProviderName: p.ProviderName}, p.Err
}

func (p ErrorProvider) IsExpired() bool {
    return false
}

type Expiry struct {
    expiration  time.Time
    CurrentTime func() time.Time
}

func (e *Expiry) SetExpiration(expiration time.Time, window time.Duration) {
    e.expiration = expiration
    if window > 0 {
        e.expiration = e.expiration.Add(-window)
    }
}

func (e *Expiry) IsExpired() bool {
    if e.CurrentTime == nil {
        e.CurrentTime = time.Now
    }
    return e.expiration.Before(e.CurrentTime())
}

type Credentials struct {
    creds        Value
    forceRefresh bool
    m            sync.Mutex

    provider Provider
}

func NewCredentials(provider Provider) *Credentials {
    return &Credentials{
        provider:     provider,
        forceRefresh: true,
    }
}

func (c *Credentials) Get() (Value, error) {
    c.m.Lock()
    defer c.m.Unlock()

    if c.isExpired() {
        creds, err := c.provider.Retrieve()
        if err != nil {
            return Value{}, err
        }
        c.creds = creds
        c.forceRefresh = false
    }

    return c.creds, nil
}

func (c *Credentials) Expire() {
    c.m.Lock()
    defer c.m.Unlock()

    c.forceRefresh = true
}

func (c *Credentials) IsExpired() bool {
    c.m.Lock()
    defer c.m.Unlock()

    return c.isExpired()
}

func (c *Credentials) isExpired() bool {
    return c.forceRefresh || c.provider.IsExpired()
}
