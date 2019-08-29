package client

import (
    "net/http"

    "github.com/visit1985/atlasgo/common/credentials"
    "github.com/visit1985/atlasgo/common/digest"
    "github.com/visit1985/atlasgo/common"
)

// A Client implements the base client request and response handling used by all service clients.
type Client struct {
    Credentials *credentials.Credentials
    HTTPClient  *http.Client
    Endpoint    string
    Error       error
}

// New will return a pointer to a new un-initialized service client.
func New() *Client {
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
    }
}

// WithCredentials will override the default Credentials Provider with the given one.
func (c *Client) WithCredentials(credentials *credentials.Credentials) *Client {
    c.Credentials = credentials
    return c
}

// WithEndpoint will override the default Endpoint with the given one.
func (c *Client) WithEndpoint(endpoint string) *Client {
    c.Endpoint = endpoint
    return c
}

// WithHTTPClient will override the default http.Client with the given one.
func (c *Client) WithHTTPClient(client *http.Client) *Client {
    c.HTTPClient = client
    return c
}

// Init will initialized the service client.
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
