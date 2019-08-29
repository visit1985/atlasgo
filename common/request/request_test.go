package request

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "github.com/visit1985/atlasgo/common/client"
    "net/http"
    "net/http/httptest"
    "net/url"
    "testing"
)

func newStubRequest(response string) *Request {
    ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        _, _ = fmt.Fprint(w, response)
    }))
    tc := ts.Client()
    c := client.New().WithHTTPClient(tc).Init()
    r := New(c, &Operation{}, nil, nil, &Handlers{})
    r.HTTPRequest.URL, r.Error= url.Parse(ts.URL)
    return r
}

func TestNew(t *testing.T) {
    c := client.New()
    r := New(c, &Operation{}, nil, nil, &Handlers{})
    assert.Nil(t, r.Error, "Expect no error")
}

func TestRequest_Send(t *testing.T) {
    r := newStubRequest("response")
    _ = r.Send()
    assert.Nil(t, r.Error, "Expect no error")
    assert.Equal(t, "response", string(r.Body))
}

func TestRequest_Paginate(t *testing.T) {
    ts2 := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        _, _ = fmt.Fprint(w, `{"results": [ {"test": "response2"} ]}`)
    }))
    r := newStubRequest(`{"results": [ {"test": "response"} ], "links": [ {"href": "` + ts2.URL + `", "rel": "next"} ]}`)
    o := &testListResponse{}
    r.Output = o
    r.Handlers = &Handlers{ResponseHandler: ListResponseHandler}
    _ = r.Paginate()
    assert.Nil(t, r.Error, "Expect no error")
    assert.Equal(t, "response", (*o)[0].Test)
    assert.Equal(t, "response2", (*o)[1].Test)
}
