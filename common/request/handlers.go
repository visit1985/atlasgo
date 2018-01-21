package request

import (
	"net/http"
	"encoding/json"
)

type Handlers struct {
	RequestHandler func(*interface{}, *http.Request) error
	ResponseHandler func(*http.Response, *interface{}) error
}

func ResponseHandler(response *http.Response, output *interface{}) error {
	return json.NewDecoder(response.Body).Decode(&output)
}

func ListResponseHandler(response *http.Response, output *interface{}) error {
	var objmap map[string]*json.RawMessage
	err := json.NewDecoder(response.Body).Decode(&objmap)
	if err != nil {
		return err
	}
	return json.Unmarshal(*objmap["results"], &output)
}
