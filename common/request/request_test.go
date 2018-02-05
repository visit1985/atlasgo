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

func newStubRequest(response string) *Request {
    ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, response)
    }))
    tc := ts.Client()
    c := client.New("group").WithHTTPClient(tc).Init()
    r := New(c, &Operation{}, nil, nil, &Handlers{})
    r.HTTPRequest.URL, r.Error= url.Parse(ts.URL)
    return r
}

func TestNew(t *testing.T) {
    c := client.New("group")
    r := New(c, &Operation{}, nil, nil, &Handlers{})
    assert.Nil(t, r.Error, "Expect no error")
}

func TestRequest_Send(t *testing.T) {
    r := newStubRequest("response")
    r.Send()
    assert.Nil(t, r.Error, "Expect no error")
    assert.Equal(t, "response", string(r.Body))
}

func TestRequest_Paginate(t *testing.T) {

}
