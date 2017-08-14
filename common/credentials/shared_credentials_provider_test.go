package credentials

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharedCredentialsProvider(t *testing.T) {
	os.Clearenv()

	p := SharedCredentialsProvider{Filename: "example.ini", Profile: ""}
	creds, err := p.Retrieve()
	assert.Nil(t, err, "Expect no error")

	assert.Equal(t, "groupid", creds.GroupID, "Expect group ID to match")
	assert.Equal(t, "username", creds.Username, "Expect username to match")
	assert.Equal(t, "secret", creds.AccessKey, "Expect access key to match")
}

func TestSharedCredentialsProviderIsExpired(t *testing.T) {
	os.Clearenv()

	p := SharedCredentialsProvider{Filename: "example.ini", Profile: ""}

	assert.True(t, p.IsExpired(), "Expect creds to be expired before retrieve")

	_, err := p.Retrieve()
	assert.Nil(t, err, "Expect no error")

	assert.False(t, p.IsExpired(), "Expect creds to not be expired after retrieve")
}

func TestSharedCredentialsProviderWithATLAS_SHARED_CREDENTIALS_FILE(t *testing.T) {
	os.Clearenv()
	os.Setenv("ATLAS_SHARED_CREDENTIALS_FILE", "example.ini")
	p := SharedCredentialsProvider{}
	creds, err := p.Retrieve()

	assert.Nil(t, err, "Expect no error")

	assert.Equal(t, "groupid", creds.GroupID, "Expect group ID to match")
	assert.Equal(t, "username", creds.Username, "Expect username to match")
	assert.Equal(t, "secret", creds.AccessKey, "Expect access key to match")
}

func TestSharedCredentialsProviderWithATLAS_SHARED_CREDENTIALS_FILEAbsPath(t *testing.T) {
	os.Clearenv()
	wd, err := os.Getwd()
	assert.NoError(t, err)
	os.Setenv("ATLAS_SHARED_CREDENTIALS_FILE", filepath.Join(wd, "example.ini"))
	p := SharedCredentialsProvider{}
	creds, err := p.Retrieve()
	assert.Nil(t, err, "Expect no error")

	assert.Equal(t, "groupid", creds.GroupID, "Expect group ID to match")
	assert.Equal(t, "username", creds.Username, "Expect username to match")
	assert.Equal(t, "secret", creds.AccessKey, "Expect access key to match")
}

func TestSharedCredentialsProviderWithATLAS_PROFILE(t *testing.T) {
	os.Clearenv()
	os.Setenv("ATLAS_PROFILE", "default")

	p := SharedCredentialsProvider{Filename: "example.ini", Profile: ""}
	creds, err := p.Retrieve()
	assert.Nil(t, err, "Expect no error")

	assert.Equal(t, "groupid", creds.GroupID, "Expect group ID to match")
	assert.Equal(t, "username", creds.Username, "Expect username to match")
	assert.Equal(t, "secret", creds.AccessKey, "Expect access key to match")
}

func TestSharedCredentialsProviderColonInCredFile(t *testing.T) {
	os.Clearenv()

	p := SharedCredentialsProvider{Filename: "example.ini", Profile: "with_colon"}
	creds, err := p.Retrieve()
	assert.Nil(t, err, "Expect no error")

	assert.Equal(t, "groupid", creds.GroupID, "Expect group ID to match")
	assert.Equal(t, "username", creds.Username, "Expect username to match")
	assert.Equal(t, "secret", creds.AccessKey, "Expect access key to match")
}

func BenchmarkSharedCredentialsProvider(b *testing.B) {
	os.Clearenv()

	p := SharedCredentialsProvider{Filename: "example.ini", Profile: ""}
	_, err := p.Retrieve()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := p.Retrieve()
		if err != nil {
			b.Fatal(err)
		}
	}
}
