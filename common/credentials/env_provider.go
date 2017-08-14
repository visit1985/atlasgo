package credentials

import (
	"os"
	"errors"
)

const EnvProviderName = "EnvProvider"

var (
	ErrGroupIDNotFound = errors.New("ATLAS_GROUP_ID not found in environment")
	ErrUsernameNotFound = errors.New("ATLAS_USERNAME not found in environment")
	ErrAccessKeyNotFound = errors.New("ATLAS_ACCESS_KEY not found in environment")
)

type EnvProvider struct {
	retrieved bool
}

func NewEnvCredentials() *Credentials {
	return NewCredentials(&EnvProvider{})
}

func (e *EnvProvider) Retrieve() (Value, error) {
	e.retrieved = false

	group_id := os.Getenv("ATLAS_GROUP_ID")

	username := os.Getenv("ATLAS_USERNAME")

	access_key := os.Getenv("ATLAS_ACCESS_KEY")

	if group_id == "" {
		return Value{ProviderName: EnvProviderName}, ErrGroupIDNotFound
	}

	if username == "" {
		return Value{ProviderName: EnvProviderName}, ErrUsernameNotFound
	}

	if access_key == "" {
		return Value{ProviderName: EnvProviderName}, ErrAccessKeyNotFound
	}

	e.retrieved = true
	return Value{
		GroupID:	group_id,
		Username:	username,
		AccessKey:	access_key,
		ProviderName:	EnvProviderName,
	}, nil
}

func (e *EnvProvider) IsExpired() bool {
	return !e.retrieved
}
