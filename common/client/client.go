package client

import (
	"net/http"

	"github.com/visit1985/atlasgo/common/credentials"
	"github.com/visit1985/atlasgo/common/digest"
	"github.com/visit1985/atlasgo/common"
)

type UnderlineString string

type Client struct {
	Credentials	*credentials.Credentials
	HTTPClient	*http.Client
	Endpoint	string
}

func New() *Client {
	return &Client{}
}

func NewClient() (*Client, error) {
	return New().Init()
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

func (c *Client) Init() (*Client, error) {
	if c.Endpoint == "" {
		c.Endpoint = common.DefaultEndpoint
	}

	if c.Credentials == nil {
		c.Credentials = credentials.NewCredentials(
			&credentials.ChainProvider{
				Providers: []credentials.Provider{
					&credentials.EnvProvider{},
					&credentials.SharedCredentialsProvider{},
				},
			},
		)
	}

	creds, err := c.Credentials.Get()
	if err != nil {
		return nil, err
	}

	client, err := digest.NewTransport(creds.Username, creds.AccessKey).Client()
	if err != nil {
		return nil, err
	}

	if c.HTTPClient == nil {
		c.HTTPClient = client
	}

	return c, err
}
