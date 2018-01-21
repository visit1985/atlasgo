package request

import (
    "encoding/json"
)

type Handlers struct {
    RequestHandler  func(*Request, *interface{}) error
    ResponseHandler func(*Request, *interface{}) error
}

func ResponseHandler(request *Request, output *interface{}) error {
    return json.Unmarshal(request.Body, &output)
}

func ListResponseHandler(request *Request, output *interface{}) error {
    var objmap map[string]*json.RawMessage
    err := json.Unmarshal(request.Body, &objmap)
    if err != nil {
        return err
    }
    return json.Unmarshal(*objmap["results"], &output)
}
