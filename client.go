package gomo

import (
	"log"
	"net/http"
	"os"
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

// NewClient creates a new client for you to make requests with. It is
// configured by passing in list of option functions.
func NewClient(options ...func(*Client)) Client {
	client := Client{
		credentials: defaultCredentials(),
		APIVersion:  defaultAPIVersion,
		Endpoint:    defaultEndpoint,
		Debug:       false,
		httpClient:  &http.Client{},
		Logger:      defaultLogger,
	}
	for _, option := range options {
		option(&client)
	}
	return client
}

// GrantType returns the string value of the current crednetials grant type
func (c *Client) GrantType() string {
	return c.credentials.grantType()
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
		c.Logger(c, msg)
	}
}

func defaultCredentials() credentials {
	return clientCredentials{
		clientID:     os.Getenv("MOLTIN_CLIENT_ID"),
		clientSecret: os.Getenv("MOLTIN_CLIENT_SECRET"),
	}
}
