package request

import (
	"net/http"
	"encoding/json"
)

type Handlers struct {
	RequestHandler func(*interface{}, *http.Request) error
	ReponseHandler func(*http.Response, *interface{}) error
}

func ListReponseHandler(response *http.Response, output *interface{}) error {
	var objmap map[string]*json.RawMessage
	err := json.NewDecoder(response.Body).Decode(&objmap)
	if err != nil {
		return err
	}
	return json.Unmarshal(*objmap["results"], &output)
}
