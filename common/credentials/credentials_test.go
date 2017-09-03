package credentials

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"
)

type stubProvider struct {
	creds   Value
	expired bool
	err     error
}

func (s *stubProvider) Retrieve() (Value, error) {
	s.expired = false
	s.creds.ProviderName = "stubProvider"
	return s.creds, s.err
}

func (s *stubProvider) IsExpired() bool {
	return s.expired
}

func TestCredentialsGet(t *testing.T) {
	c := NewCredentials(&stubProvider{
		creds: Value{
			Username:	"username",
			AccessKey:	"secret",
		},
		expired: true,
	})

	creds, err := c.Get()
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, "username", creds.Username, "Expect username to match")
	assert.Equal(t, "secret", creds.AccessKey, "Expect access key to match")
}

func TestCredentialsGetWithError(t *testing.T) {
	c := NewCredentials(&stubProvider{err: errors.New("provider error"), expired: true})

	_, err := c.Get()
	assert.Equal(t, "provider error", err.Error(), "Expected provider error")
}

func TestCredentialsExpire(t *testing.T) {
	stub := &stubProvider{}
	c := NewCredentials(stub)

	stub.expired = false
	assert.True(t, c.IsExpired(), "Expected to start out expired")
	c.Expire()
	assert.True(t, c.IsExpired(), "Expected to be expired")

	c.forceRefresh = false
	assert.False(t, c.IsExpired(), "Expected not to be expired")

	stub.expired = true
	assert.True(t, c.IsExpired(), "Expected to be expired")
}

func TestCredentialsGetWithProviderName(t *testing.T) {
	stub := &stubProvider{}

	c := NewCredentials(stub)

	creds, err := c.Get()
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, creds.ProviderName, "stubProvider", "Expected provider name to match")
}
