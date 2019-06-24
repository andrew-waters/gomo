package gomo

import (
	"net/http"
	"os"
)

// ClientCredentials configures a client with client credentials
func ClientCredentials(clientID string, clientSecret string) func(*Client) {
	return func(c *Client) {
		c.credentials = clientCredentials{
			clientID:     os.Getenv(clientID),
			clientSecret: os.Getenv(clientSecret),
		}
	}
}

// ImplicitCredentials configures a client with implicit credentials
func ImplicitCredentials(clientID string, clientSecret string) func(*Client) {
	return func(c *Client) {
		c.credentials = implicitCredentials{
			clientID: os.Getenv(clientID),
		}
	}
}

// APIVersion configures the API version for the client
func APIVersion(apiVersion string) func(*Client) {
	return func(c *Client) {
		c.APIVersion = apiVersion
	}
}

// Endpoint configures the API endpoint for the client
func Endpoint(endpoint string) func(*Client) {
	return func(c *Client) {
		c.Endpoint = endpoint
	}
}

// Debug turns of client debugging
func Debug(c *Client) {
	c.Debug = true
}

// HTTPClient configures a client to use an http.Client
func HTTPClient(client *http.Client) func(*Client) {
	return func(c *Client) {
		c.httpClient = client
	}
}

// Logger configures a client to use a logger
func Logger(logger func(*Client, interface{})) func(*Client) {
	return func(c *Client) {
		c.Logger = logger
	}
}
