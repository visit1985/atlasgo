package request

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
)

// A Handlers provides a collection of handlers for various stages of handling requests.
type Handlers struct {
    RequestHandler  func(*Request, *interface{}) error
    ResponseHandler func(*Request, *interface{}) error
}

// RequestHandler encodes a structure into a JSON string
func RequestHandler(request *Request, input *interface{}) error {
    jsonstr, err := json.Marshal(&input)
    request.HTTPRequest.Body = ioutil.NopCloser(bytes.NewBuffer(jsonstr))
    return err
}

// ResponseHandler decodes a JSON string into a structure.
func ResponseHandler(request *Request, output *interface{}) error {
    return json.Unmarshal(request.Body, &output)
}

// ListResponseHandler extracts results from a JSON envelope and decodes them into a structure.
// https://docs.atlas.mongodb.com/api/#lists
func ListResponseHandler(request *Request, output *interface{}) error {
    var objmap map[string]*json.RawMessage
    err := json.Unmarshal(request.Body, &objmap)
    if err != nil {
        return err
    }
    return json.Unmarshal(*objmap["results"], &output)
}
