package gomo

import (
	"log"
	"net/http"
)

const (
	defaultAPIVersion = "v2"
	defaultEndpoint   = "https://api.moltin.com"
	defaultUserAgent  = "gomo"
)

var defaultLogger = func(c *Client, msg interface{}) {
	if c.Debug {
		log.Println(msg)
	}
}

// Client is the main client struct
type Client struct {
	credentials credentials
	APIVersion  string
	Endpoint    string
	AccessToken string
	Debug       bool
	Logs        []interface{}
	httpClient  *http.Client
	Logger      func(*Client, interface{})
}

// NewClient creates a new client for you to make requests with
func NewClient(c credentials) Client {
	return Client{
		credentials: c,
		APIVersion:  defaultAPIVersion,
		Endpoint:    defaultEndpoint,
		Debug:       false,
		httpClient:  &http.Client{},
		Logger:      defaultLogger,
	}
}

// GrantType returns the string value of the current crednetials grant type
func (c *Client) GrantType() string {
	return c.credentials.grantType()
}

// CustomEndpoint overrides the endpoints that the client accesses
func (c *Client) CustomEndpoint(e string) {
	c.Endpoint = e
}

// CustomHTTPClient overrides the default http.Client that the client uses
func (c *Client) CustomHTTPClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

// EnableDebug logs debugging info from the API calls
func (c *Client) EnableDebug() {
	c.Debug = true
}

// DisableDebug stops logs form API calls
func (c *Client) DisableDebug() {
	c.Debug = false
}

// CustomLogger allows you to pass in a custom log function to override the client default
func (c *Client) CustomLogger(l func(*Client, interface{})) {
	c.Logger = l
}

// Log will dump debug info onto stdout
func (c *Client) Log(msgs ...interface{}) {
	for _, msg := range msgs {
		c.Logs = append(c.Logs, msg)
		c.Logger(c, msg)
	}
}
