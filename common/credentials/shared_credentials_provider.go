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

// A SharedCredentialsProvider retrieves credentials from the current user's home directory,
// and keeps track if those credentials are expired.
//
// Profile ini file example: $HOME/.atlas/credentials
type SharedCredentialsProvider struct {
    Filename  string
    Profile   string
    retrieved bool
}

// NewSharedCredentials returns a pointer to a new Credentials object
// wrapping the Profile file provider.
func NewSharedCredentials(filename, profile string) *Credentials {
    return NewCredentials(&SharedCredentialsProvider{
        Filename: filename,
        Profile:  profile,
    })
}

// Retrieve reads and extracts the shared credentials from the current users home directory.
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

// IsExpired returns if the shared credentials have expired.
func (p *SharedCredentialsProvider) IsExpired() bool {
    return !p.retrieved
}

// loadProfiles loads from the file pointed to by shared credentials filename for profile.
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
        Username:    username.String(),
        AccessKey:    access_key.String(),
        ProviderName:    SharedCredsProviderName,
    }, nil
}

// filename returns the filename to use to read shared credentials.
// Will return an error if the user's home directory path cannot be found.
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

// profile returns the shared credentials profile.
// If empty will read environment variable "ATLAS_PROFILE".
// If that is not set profile will return "default".
func (p *SharedCredentialsProvider) profile() string {
    if p.Profile == "" {
        p.Profile = os.Getenv("ATLAS_PROFILE")
    }
    if p.Profile == "" {
        p.Profile = "default"
    }

    return p.Profile
}

// UserHomeDir returns the home directory for the user the process is running under.
func UserHomeDir() string {
    if runtime.GOOS == "windows" { // Windows
        return os.Getenv("USERPROFILE")
    }

    return os.Getenv("HOME")
}

// SharedCredentialsFilename returns the SDK's default file path for the shared credentials file.
func SharedCredentialsFilename() string {
    return filepath.Join(UserHomeDir(), ".atlas", "credentials")
}
