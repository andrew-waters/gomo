package gomo

import (
	"net/http"
)

// ClientCredentials configures a client with client credentials
func ClientCredentials(clientID string, clientSecret string) ClientOption {
	return func(c *Client) {
		c.credentials = clientCredentials{
			clientID:     clientID,
			clientSecret: clientSecret,
		}
	}
}

// ImplicitCredentials configures a client with implicit credentials
func ImplicitCredentials(clientID string) ClientOption {
	return func(c *Client) {
		c.credentials = implicitCredentials{
			clientID: clientID,
		}
	}
}

// APIVersion configures the API version for the client
func APIVersion(apiVersion string) ClientOption {
	return func(c *Client) {
		c.APIVersion = apiVersion
	}
}

// Endpoint configures the API endpoint for the client
func Endpoint(endpoint string) ClientOption {
	return func(c *Client) {
		c.Endpoint = endpoint
	}
}

// Debug turns on client debugging, which is off by default
func Debug() ClientOption {
	return func(c *Client) {
		c.Debug = true
	}
}

// HTTPClient configures a client to use an http.Client
func HTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

// Logger configures a client to use a logger
func Logger(logger func(*Client, interface{})) ClientOption {
	return func(c *Client) {
		c.Logger = logger
	}
}
