package client

import (
	"net/http"

	"github.com/visit1985/atlasgo/common/credentials"
	"github.com/visit1985/atlasgo/common/digest"
	"github.com/visit1985/atlasgo/common/request"
	"github.com/visit1985/atlasgo/common"
)

type UnderlineString string

type Client struct {
	Credentials	*credentials.Credentials
	GroupID		string
	HTTPClient	*http.Client
	Endpoint	string
	Error		error
}

func New(gid string) *Client {
	return &Client{

		Credentials: credentials.NewCredentials(
			&credentials.ChainProvider{
				Providers: []credentials.Provider{
					&credentials.EnvProvider{},
					&credentials.SharedCredentialsProvider{},
				},
			},
		),
		Endpoint: common.DefaultEndpoint,
		GroupID: gid,
	}
}

func (c *Client) WithCredentials(credentials *credentials.Credentials) *Client {
	c.Credentials = credentials
	return c
}

func (c *Client) WithEndpoint(endpoint string) *Client {
	c.Endpoint = endpoint
	return c
}

func (c *Client) WithHTTPClient(client *http.Client) *Client {
	c.HTTPClient = client
	return c
}

func (c *Client) Init() *Client {
	creds, err := c.Credentials.Get()
	if err != nil {
		c.Error = err
		return c
	}

	if c.HTTPClient == nil {
		c.HTTPClient, err = digest.NewTransport(creds.Username, creds.AccessKey).Client()
	}
	return c
}

func (c *Client) NewRequest(operation *request.Operation, input interface{}, output interface{}, handlers *request.Handlers) *request.Request {
	return request.New(c.HTTPClient, c.Endpoint, operation, input, output, handlers)
}
