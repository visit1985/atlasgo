package credentials

import (
	"errors"
	"fmt"
	"os"
	"github.com/go-ini/ini"
	"path/filepath"
	"runtime"
)

const SharedCredsProviderName = "SharedCredentialsProvider"

var (
	ErrSharedCredentialsHomeNotFound = errors.New("user home directory not found.")
)

func UserHomeDir() string {
	if runtime.GOOS == "windows" { // Windows
		return os.Getenv("USERPROFILE")
	}

	return os.Getenv("HOME")
}

func SharedCredentialsFilename() string {
	return filepath.Join(UserHomeDir(), ".atlas", "credentials")
}

type SharedCredentialsProvider struct {
	Filename string
	Profile string
	retrieved bool
}

func NewSharedCredentials(filename, profile string) *Credentials {
	return NewCredentials(&SharedCredentialsProvider{
		Filename: filename,
		Profile:  profile,
	})
}

func (p *SharedCredentialsProvider) Retrieve() (Value, error) {
	p.retrieved = false

	filename, err := p.filename()
	if err != nil {
		return Value{ProviderName: SharedCredsProviderName}, err
	}

	creds, err := loadProfile(filename, p.profile())
	if err != nil {
		return Value{ProviderName: SharedCredsProviderName}, err
	}

	p.retrieved = true
	return creds, nil
}

func (p *SharedCredentialsProvider) IsExpired() bool {
	return !p.retrieved
}

func loadProfile(filename, profile string) (Value, error) {
	config, err := ini.Load(filename)
	if err != nil {
		return Value{ProviderName: SharedCredsProviderName}, errors.New(
			"failed to load shared credentials file")
	}
	iniProfile, err := config.GetSection(profile)
	if err != nil {
		return Value{ProviderName: SharedCredsProviderName}, errors.New(
			"failed to get profile")
	}

	group_id, err := iniProfile.GetKey("atlas_group_id")
	if err != nil {
		return Value{ProviderName: SharedCredsProviderName}, errors.New(
			fmt.Sprintf("shared credentials %s in %s did not contain atlas_group_id", profile, filename))
	}

	username, err := iniProfile.GetKey("atlas_username")
	if err != nil {
		return Value{ProviderName: SharedCredsProviderName}, errors.New(
			fmt.Sprintf("shared credentials %s in %s did not contain atlas_username", profile, filename))
	}

	access_key, err := iniProfile.GetKey("atlas_access_key")
	if err != nil {
		return Value{ProviderName: SharedCredsProviderName}, errors.New(
			fmt.Sprintf("shared credentials %s in %s did not contain atlas_access_key", profile, filename))
	}

	return Value{
		GroupID:	group_id.String(),
		Username:	username.String(),
		AccessKey:	access_key.String(),
		ProviderName:	SharedCredsProviderName,
	}, nil
}

func (p *SharedCredentialsProvider) filename() (string, error) {
	if len(p.Filename) != 0 {
		return p.Filename, nil
	}

	if p.Filename = os.Getenv("ATLAS_SHARED_CREDENTIALS_FILE"); len(p.Filename) != 0 {
		return p.Filename, nil
	}

	if home := UserHomeDir(); len(home) == 0 {
		// Backwards compatibility of home directly not found error being returned.
		// This error is too verbose, failure when opening the file would of been
		// a better error to return.
		return "", ErrSharedCredentialsHomeNotFound
	}

	p.Filename = SharedCredentialsFilename()

	return p.Filename, nil
}

func (p *SharedCredentialsProvider) profile() string {
	if p.Profile == "" {
		p.Profile = os.Getenv("ATLAS_PROFILE")
	}
	if p.Profile == "" {
		p.Profile = "default"
	}

	return p.Profile
}
