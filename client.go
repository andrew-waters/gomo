package gomo

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

const (
	defaultAPIVersion = "v2"
	defaultEndpoint   = "https://api.moltin.com"
	defaultUserAgent  = "gomo"
)

// Client is the main client struct
type Client struct {
	credentials credentials
	APIVersion  string
	Endpoint    string
	AccessToken string
	Debug       bool
	Logs        []interface{}
	httpClient  *http.Client
}

// NewClient creates a new client for you to make requests with
func NewClient(c credentials) Client {
	return Client{
		credentials: c,
		APIVersion:  defaultAPIVersion,
		Endpoint:    defaultEndpoint,
		Debug:       false,
		httpClient:  &http.Client{},
	}
}

// NewClientWithCustomEndpoint creates a new client for you to make requests with to a different endpoint
func NewClientWithCustomEndpoint(c credentials, e string) Client {
	return Client{
		credentials: c,
		APIVersion:  defaultAPIVersion,
		Endpoint:    e,
		Debug:       false,
		httpClient:  &http.Client{},
	}
}

// GrantType returns the string value of the current crednetials grant type
func (c Client) GrantType() string {
	return c.credentials.grantType()
}

// CustomEndpoint overrides the endpoints that the client accesses
func (c *Client) CustomEndpoint(e string) {
	c.Endpoint = e
}

// EnableDebug logs debugging info from the API calls
func (c *Client) EnableDebug() {
	c.Debug = true
}

// DisableDebug stops logs form API calls
func (c *Client) DisableDebug() {
	c.Debug = false
}

// Log will dump debug info onto stdout
func (c *Client) Log(msgs ...interface{}) {
	for _, msg := range msgs {
		c.Logs = append(c.Logs, msg)
		if c.Debug {
			spew.Dump(msg)
		}
	}
}
