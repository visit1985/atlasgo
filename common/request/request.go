package request

import (
    "net/http"
    "net/url"
    "errors"
    "github.com/visit1985/atlasgo/common/client"
    "encoding/json"
    "io/ioutil"
    "reflect"
)

// A Request is the service request to be made.
type Request struct {
    Client       *client.Client
    Operation    *Operation
    HTTPRequest  *http.Request
    Handlers     *Handlers
    HTTPResponse *http.Response
    Body         []byte
    Input        interface{}
    Error        error
    Output       interface{}
}

// An Operation is the service API operation to be made.
type Operation struct {
    Name       string
    HTTPMethod string
    HTTPPath   string
}

// A jsonError is the JSON structure a service API error returns.
// https://docs.atlas.mongodb.com/api/#errors
type jsonError struct {
    Detail string `json:"detail"`
    Error  int    `json:"error"`
    Reason string `json:"reason"`
}

// paginationLinks is the JSON structure of the envelope around a API list response.
// https://docs.atlas.mongodb.com/api/#lists
type paginationLinks struct {
    Links []struct {
        Href string `json:"href"`
        Rel  string `json:"rel"`
    } `json:"links"`
}

// New returns a new Request pointer for the service API operation and parameters.
func New(client *client.Client, operation *Operation, input interface{}, output interface{}, handlers *Handlers) *Request {
    method := operation.HTTPMethod
    if method == "" {
        method = "GET"
    }

    httpReq, _ := http.NewRequest(method, "", nil)

    var err error
    httpReq.URL, err = url.Parse(client.Endpoint + operation.HTTPPath)

    if err != nil {
        httpReq.URL = &url.URL{}
        err = errors.New("invalid endpoint uri")
    }

    r := &Request{
        Client:      client,
        Operation:   operation,
        HTTPRequest: httpReq,
        Handlers:    handlers,
        Input:       input,
        Error:       err,
        Output:      output,
    }

    return r
}

// Copy returns a new Request pointer for the same service API operation and parameters.
// The copied Request contains a new HTTPRequest object with an un-initialized URL.
func (r *Request) Copy() *Request {
    req := &Request{}
    *req = *r
    httpReq, _ := http.NewRequest(r.HTTPRequest.Method, "", nil)
    req.HTTPRequest = httpReq
    req.HTTPResponse = nil
    req.Body = nil
    req.Output = reflect.New(reflect.TypeOf(r.Output).Elem()).Interface()
    return req
}

// Send will send the request returning error if errors are encountered.
func (r *Request) Send() error {
    if r.Client.Error != nil {
        r.Error = r.Client.Error
        return r.Error
    }

    // prepare input data for http request
    if r.Handlers.RequestHandler != nil {
        r.Error = r.Handlers.RequestHandler(r, &r.Input)
    }

    // do the request
    r.HTTPResponse, r.Error = r.Client.HTTPClient.Do(r.HTTPRequest)
    if r.Error != nil {
        return r.Error
    }

    // read the response
    defer r.HTTPResponse.Body.Close()
    r.Body, r.Error = ioutil.ReadAll(r.HTTPResponse.Body)

    // handle http errors
    if r.HTTPResponse.StatusCode >= 400 {
        var err jsonError
        r.Error = json.Unmarshal(r.Body, &err)
        if r.Error != nil {
            return r.Error
        }
        r.Error = errors.New(err.Detail)
        return r.Error
    }

    // transform http response
    if r.Handlers.ResponseHandler != nil {
        r.Error = r.Handlers.ResponseHandler(r, &r.Output)
    }

    return r.Error
}

// NextPage will attempt to retrieve a new Request pointer for the next page of the API operation.
// It will return nil if the page cannot be retrieved, or there are no more pages.
func (r *Request) NextPage() *Request {
    var attr paginationLinks
    r.Error = json.Unmarshal(r.Body, &attr)

    for i := range attr.Links {
        if attr.Links[i].Rel == "next" {
            req := r.Copy()
            req.HTTPRequest.URL, req.Error = url.Parse(attr.Links[i].Href)
            return req
        }
    }

    return nil
}

// Paginate iterates over each page of a paginated Request object and merges responses and errors
// back to the initial request.
//
// The type of the Requests Output structure must be []struct.
func (r *Request) Paginate() error {
    var err error

    if err = r.Send(); err != nil {
        r.Error = err
        return err
    }

    for page := r.NextPage(); page != nil; page = page.NextPage() {
        if err = page.Send(); err != nil {
            r.Error = err
            return err
        }

        // add page output to original output
        src := reflect.ValueOf(page.Output).Elem()
        dst := reflect.ValueOf(r.Output).Elem()
        if src.Kind() == reflect.Slice && dst.Kind() == reflect.Slice {
            for i := 0; i < src.Len(); i++ {
                dst.Set(reflect.Append(dst, src.Index(i)))
            }
        }
    }

    return nil
}
