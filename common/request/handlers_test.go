package request

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

type testResponse struct {
    Test string `json:"test"`
}

func TestResponseHandler(t *testing.T) {
    r := newStubRequest(`{"test": "response"}`)
    o := &testResponse{}
    r.Output = o
    r.Handlers = &Handlers{ResponseHandler: ResponseHandler}
    r.Send()
    assert.Nil(t, r.Error, "Expect no error")
    assert.Equal(t, "response", o.Test)
}

type testListResponse []struct {
    Test string `json:"test"`
}

func TestListResponseHandler(t *testing.T) {
    r := newStubRequest(`{"results": [ {"test": "response"}, {"test": "response2"} ]}`)
    o := &testListResponse{}
    r.Output = o
    r.Handlers = &Handlers{ResponseHandler: ListResponseHandler}
    r.Send()
    assert.Nil(t, r.Error, "Expect no error")
    assert.Equal(t, "response", (*o)[0].Test)
    assert.Equal(t, "response2", (*o)[1].Test)
}
