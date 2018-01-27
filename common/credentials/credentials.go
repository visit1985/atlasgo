package credentials

import (
    "sync"
    "time"
)

// A Value is the credentials value for individual credential fields.
type Value struct {
    Username     string
    AccessKey    string
    ProviderName string
}

// A Provider is the interface for any component which will provide credentials Value.
type Provider interface {
    Retrieve()  (Value, error)
    IsExpired() bool
}

// An ErrorProvider is a stub credentials provider that always returns an error.
type ErrorProvider struct {
    Err          error
    ProviderName string
}

// Retrieve will always return the error that the ErrorProvider was created with.
func (p ErrorProvider) Retrieve() (Value, error) {
    return Value{ProviderName: p.ProviderName}, p.Err
}

// IsExpired will always return not expired.
func (p ErrorProvider) IsExpired() bool {
    return false
}

// A Expiry provides shared expiration logic to be used by credentials providers to implement expiry functionality.
type Expiry struct {
    expiration  time.Time
    CurrentTime func() time.Time
}

// SetExpiration sets the expiration IsExpired will check when called.
func (e *Expiry) SetExpiration(expiration time.Time, window time.Duration) {
    e.expiration = expiration
    if window > 0 {
        e.expiration = e.expiration.Add(-window)
    }
}

// IsExpired returns if the credentials are expired.
func (e *Expiry) IsExpired() bool {
    if e.CurrentTime == nil {
        e.CurrentTime = time.Now
    }
    return e.expiration.Before(e.CurrentTime())
}

// A Credentials provides synchronous safe retrieval of credentials Value.
type Credentials struct {
    creds        Value
    forceRefresh bool
    m            sync.Mutex

    provider Provider
}

// NewCredentials returns a pointer to a new Credentials with the provider set.
func NewCredentials(provider Provider) *Credentials {
    return &Credentials{
        provider:     provider,
        forceRefresh: true,
    }
}

// Get returns the credentials value, or error if the credentials Value failed to be retrieved.
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

// Expire expires the credentials and forces them to be retrieved on the next call to Get().
func (c *Credentials) Expire() {
    c.m.Lock()
    defer c.m.Unlock()

    c.forceRefresh = true
}

// IsExpired returns if the credentials are no longer valid, and need to be retrieved.
func (c *Credentials) IsExpired() bool {
    c.m.Lock()
    defer c.m.Unlock()

    return c.isExpired()
}

// isExpired helper method wrapping the definition of expired credentials.
func (c *Credentials) isExpired() bool {
    return c.forceRefresh || c.provider.IsExpired()
}
