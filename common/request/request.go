package request

import (
	"net/http"
	"net/url"
	"errors"
)

type Request struct {
	Client                 *http.Client
	Operation              *Operation
	HTTPRequest            *http.Request
	Handlers               *Handlers
	HTTPResponse           *http.Response
	Input                  interface{}
	Error                  error
	Output                 interface{}
}

type Operation struct {
	Name       string
	HTTPMethod string
	HTTPPath   string
}

func New(client *http.Client, endpoint string, operation *Operation, input interface{}, output interface{}, handlers *Handlers) *Request {
	method := operation.HTTPMethod
	if method == "" {
		method = "GET"
	}

	httpReq, _ := http.NewRequest(method, "", nil)

	var err error
	httpReq.URL, err = url.Parse(endpoint + operation.HTTPPath)

	if err != nil {
		httpReq.URL = &url.URL{}
		err = errors.New("invalid endpoint uri")
	}

	r := &Request{
		Client:		client,
		Operation:	operation,
		HTTPRequest:	httpReq,
		Handlers:       handlers,
		Input:		input,
		Error:		err,
		Output:		output,
	}

	return r
}

func (r *Request) Send() error {
	var err error

	// prepare input data for http request
	if r.Handlers.RequestHandler != nil {
		r.Handlers.RequestHandler(&r.Input, r.HTTPRequest)
	}

	// do the request
	r.HTTPResponse, err = r.Client.Do(r.HTTPRequest)

	// transform http response
	r.Handlers.ReponseHandler(r.HTTPResponse, &r.Output)

	return err
}
