package request

import (
    "encoding/json"
)

// A Handlers provides a collection of handlers for various stages of handling requests.
type Handlers struct {
    RequestHandler  func(*Request, *interface{}) error
    ResponseHandler func(*Request, *interface{}) error
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
