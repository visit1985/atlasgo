package credentials

import (
    "os"
    "errors"
)

const EnvProviderName = "EnvProvider"

var (
    ErrUsernameNotFound = errors.New("ATLAS_USERNAME not found in environment")
    ErrAccessKeyNotFound = errors.New("ATLAS_ACCESS_KEY not found in environment")
)

// A EnvProvider retrieves credentials from the environment variables of the running process.
// Environment credentials never expire.
type EnvProvider struct {
    retrieved bool
}

// NewEnvCredentials returns a pointer to a new Credentials object wrapping the environment variable provider.
func NewEnvCredentials() *Credentials {
    return NewCredentials(&EnvProvider{})
}

// Retrieve retrieves the keys from the environment.
func (e *EnvProvider) Retrieve() (Value, error) {
    e.retrieved = false

    username := os.Getenv("ATLAS_USERNAME")

    access_key := os.Getenv("ATLAS_ACCESS_KEY")

    if username == "" {
        return Value{ProviderName: EnvProviderName}, ErrUsernameNotFound
    }

    if access_key == "" {
        return Value{ProviderName: EnvProviderName}, ErrAccessKeyNotFound
    }

    e.retrieved = true
    return Value{
        Username:     username,
        AccessKey:    access_key,
        ProviderName: EnvProviderName,
    }, nil
}

// IsExpired returns if the credentials have been retrieved.
func (e *EnvProvider) IsExpired() bool {
    return !e.retrieved
}
