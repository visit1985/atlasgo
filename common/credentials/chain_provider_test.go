package credentials

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "errors"
)

type secondStubProvider struct {
    creds   Value
    expired bool
    err     error
}

func (s *secondStubProvider) Retrieve() (Value, error) {
    s.expired = false
    s.creds.ProviderName = "secondStubProvider"
    return s.creds, s.err
}

func (s *secondStubProvider) IsExpired() bool {
    return s.expired
}

func TestChainProviderWithNames(t *testing.T) {
    p := &ChainProvider{
        Providers: []Provider{
            &stubProvider{err: errors.New("first provider error")},
            &stubProvider{err: errors.New("second provider error")},
            &secondStubProvider{
                creds: Value{
                    Username:  "USERNAME",
                    AccessKey: "SECRET",
                },
            },
            &stubProvider{
                creds: Value{
                    Username:  "USERNAME",
                    AccessKey: "SECRET",
                },
            },
        },
    }

    creds, err := p.Retrieve()
    assert.Nil(t, err, "Expect no error")
    assert.Equal(t, "secondStubProvider", creds.ProviderName, "Expect provider name to match")
    assert.Equal(t, "USERNAME", creds.Username, "Expect username to match")
    assert.Equal(t, "SECRET", creds.AccessKey, "Expect access key to match")
}

func TestChainProviderGet(t *testing.T) {
    p := &ChainProvider{
        Providers: []Provider{
            &stubProvider{err: errors.New("first provider error")},
            &stubProvider{err: errors.New("second provider error")},
            &stubProvider{
                creds: Value{
                    Username:  "USERNAME",
                    AccessKey: "SECRET",
                },
            },
        },
    }

    creds, err := p.Retrieve()
    assert.Nil(t, err, "Expect no error")
    assert.Equal(t, "USERNAME", creds.Username, "Expect username to match")
    assert.Equal(t, "SECRET", creds.AccessKey, "Expect access key to match")
}

func TestChainProviderIsExpired(t *testing.T) {
    stubProvider := &stubProvider{expired: true}
    p := &ChainProvider{
        Providers: []Provider{
            stubProvider,
        },
    }

    assert.True(t, p.IsExpired(), "Expect expired to be true before any Retrieve")
    _, err := p.Retrieve()
    assert.Nil(t, err, "Expect no error")
    assert.False(t, p.IsExpired(), "Expect not expired after retrieve")

    stubProvider.expired = true
    assert.True(t, p.IsExpired(), "Expect return of expired provider")

    _, err = p.Retrieve()
    assert.False(t, p.IsExpired(), "Expect not expired after retrieve")
}

func TestChainProviderWithNoProvider(t *testing.T) {
    p := &ChainProvider{
        Providers: []Provider{},
    }

    assert.True(t, p.IsExpired(), "Expect expired with no providers")
    _, err := p.Retrieve()
    assert.Equal(t,
        ErrNoValidProvidersFoundInChain,
        err,
        "Expect no providers error returned")
}

func TestChainProviderWithNoValidProvider(t *testing.T) {
    errs := []error{
        errors.New("first provider error"),
        errors.New("second provider error"),
    }
    p := &ChainProvider{
        Providers: []Provider{
            &stubProvider{err: errs[0]},
            &stubProvider{err: errs[1]},
        },
    }

    assert.True(t, p.IsExpired(), "Expect expired with no providers")
    _, err := p.Retrieve()

    assert.Equal(t,
        ErrNoValidProvidersFoundInChain,
        err,
        "Expect no providers error returned")
}

func TestChainProviderWithNoValidProviderWithVerboseEnabled(t *testing.T) {
    errs := []error{
        errors.New("first provider error"),
        errors.New("second provider error"),
    }
    p := &ChainProvider{
        VerboseErrors: true,
        Providers: []Provider{
            &stubProvider{err: errs[0]},
            &stubProvider{err: errs[1]},
        },
    }

    assert.True(t, p.IsExpired(), "Expect expired with no providers")
    _, err := p.Retrieve()

    assert.Equal(t,
        errors.New("no valid providers in chain"),
        err,
        "Expect no providers error returned")
}
