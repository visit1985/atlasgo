package digest

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "net/http"
)

func TestNewTransport(t *testing.T) {
    tp := NewTransport("username", "password")
    assert.IsType(t, &Transport{}, tp, "Expect *digest.Transport type")
    c, err := tp.Client()
    assert.Nil(t, err, "Expect no error")
    assert.IsType(t, &http.Client{}, c, "Expect *http.Client type")
    // TODO: mock http server and check headers sent
}
