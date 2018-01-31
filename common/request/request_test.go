package request

import (
    "testing"
    "github.com/visit1985/atlasgo/common/client"
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    "fmt"
    "net/url"
)

func TestNew(t *testing.T) {
    cli := client.New("test")
    op := &Operation{
        Name: "test",
        HTTPMethod: "GET",
        HTTPPath: "/test",
    }
    handlers := &Handlers{}
    req := New(cli, op, nil, nil, handlers)
    assert.Nil(t, req.Error, "Expect no error")
}

func TestRequest_Send(t *testing.T) {
    ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "test")
    }))
    tc := ts.Client()
    cli := client.New("test").WithHTTPClient(tc).Init()
    op := &Operation{
        Name: "test",
        HTTPMethod: "GET",
        HTTPPath: "/test",
    }
    handlers := &Handlers{}
    req := New(cli, op, nil, nil, handlers)
    assert.Nil(t, req.Error, "Expect no error")
    req.HTTPRequest.URL, req.Error= url.Parse(ts.URL)
    req.Send()
    assert.Nil(t, req.Error, "Expect no error")
    assert.Equal(t, "test\n", string(req.Body))
}
