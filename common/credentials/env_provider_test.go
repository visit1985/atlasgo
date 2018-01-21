package credentials

import (
    "github.com/stretchr/testify/assert"
    "os"
    "testing"
)

func TestEnvProviderRetrieve(t *testing.T) {
    os.Clearenv()
    os.Setenv("ATLAS_USERNAME", "username")
    os.Setenv("ATLAS_ACCESS_KEY", "secret")

    e := EnvProvider{}
    creds, err := e.Retrieve()
    assert.Nil(t, err, "Expect no error")

    assert.Equal(t, "username", creds.Username, "Expect username to match")
    assert.Equal(t, "secret", creds.AccessKey, "Expect access key to match")
}

func TestEnvProviderIsExpired(t *testing.T) {
    os.Clearenv()
    os.Setenv("ATLAS_USERNAME", "username")
    os.Setenv("ATLAS_ACCESS_KEY", "secret")

    e := EnvProvider{}

    assert.True(t, e.IsExpired(), "Expect creds to be expired before retrieve.")

    _, err := e.Retrieve()
    assert.Nil(t, err, "Expect no error")

    assert.False(t, e.IsExpired(), "Expect creds to not be expired after retrieve.")
}

func TestEnvProviderNoUsername(t *testing.T) {
    os.Clearenv()
    os.Setenv("ATLAS_ACCESS_KEY", "secret")

    e := EnvProvider{}
    creds, err := e.Retrieve()
    assert.Equal(t, ErrUsernameNotFound, err, "ErrUsernameNotFound expected, but was %#v error: %#v", creds, err)
}

func TestEnvProviderNoAccessKey(t *testing.T) {
    os.Clearenv()
    os.Setenv("ATLAS_USERNAME", "username")

    e := EnvProvider{}
    creds, err := e.Retrieve()
    assert.Equal(t, ErrAccessKeyNotFound, err, "ErrAccessKeyNotFound expected, but was %#v error: %#v", creds, err)
}
