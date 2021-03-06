package client

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/visit1985/atlasgo/common"
    "github.com/visit1985/atlasgo/common/credentials"
    "net/http"
    "errors"
)

type stubProvider struct {
    creds   credentials.Value
    expired bool
    err     error
}

func (s *stubProvider) Retrieve() (credentials.Value, error) {
    s.expired = false
    s.creds.ProviderName = "stubProvider"
    return s.creds, s.err
}

func (s *stubProvider) IsExpired() bool {
    return s.expired
}

func TestClient(t *testing.T) {
    creds := credentials.NewCredentials(&stubProvider{
        creds: credentials.Value{
            Username:    "username",
            AccessKey:    "secret",
        },
        expired: true,
    })

    client := New().WithCredentials(creds).Init()

    assert.Nil(t, client.Error, "Expect no error")
    assert.Equal(t, common.DefaultEndpoint, client.Endpoint, "Expect endpoint to match")
    assert.IsType(t, &http.Client{}, client.HTTPClient, "Expected http.Client to match")
}

func TestClientNoCreds(t *testing.T) {
    creds := credentials.NewCredentials(&stubProvider{
        err: errors.New("credentials error"),
        expired: true,
    })

    client := New().WithCredentials(creds).Init()

    assert.Equal(t, "credentials error", client.Error.Error(), "Expect credentials error")
}
