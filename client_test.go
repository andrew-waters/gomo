package gomo

import (
	"net/http"
	"testing"
)

func testClient() Client {
	return Client{
		credentials: NewClientCredentials("abc", "def"),
		APIVersion:  defaultAPIVersion,
		Endpoint:    defaultEndpoint,
		Debug:       false,
		httpClient:  &http.Client{},
	}
}

func TestClientDefaults(t *testing.T) {
	c := testClient()

	// test the default endpoint
	expectedEndpoint := "https://api.moltin.com"
	if c.Endpoint != expectedEndpoint {
		t.Errorf("Incorrect default endpoint: %s (expected %s)", c.Endpoint, expectedEndpoint)
	}

	// test the API version
	expectedAPIVersion := "v2"
	if c.APIVersion != expectedAPIVersion {
		t.Errorf("Incorrect API version: %s (expected %s)", c.APIVersion, expectedAPIVersion)
	}

	// test the defaul debug status
	if c.Debug != false {
		t.Errorf("Incorrect debug val: %t (expected %t)", c.Debug, false)
	}
}

func TestClientDebug(t *testing.T) {
	c := testClient()

	// test enabling debug
	c.EnableDebug()
	if c.Debug != true {
		t.Errorf("Incorrect debug val: %t (expected %t)", c.Debug, true)
	}

	// test disabling debug
	c.DisableDebug()
	if c.Debug != false {
		t.Errorf("Incorrect debug val: %t (expected %t)", c.Debug, false)
	}
}

func TestCustomEndpoint(t *testing.T) {
	endpoint := "https://yourdomain.com"

	c := testClient()

	c.CustomEndpoint(endpoint)

	if c.Endpoint != endpoint {
		t.Errorf("Incorrect debug val: %s (expected %s)", c.Endpoint, endpoint)
	}

}

func TestNewClientWithCustomEndpoint(t *testing.T) {
	endpoint := "https://yourdomain.com"
	c := NewClientWithCustomEndpoint(
		NewClientCredentials("abc", "def"),
		endpoint,
	)

	if c.Endpoint != endpoint {
		t.Errorf("Incorrect endpoint: %s (expected %s)", c.Endpoint, endpoint)
	}
}

func TestClientGrantTypeImplicit(t *testing.T) {
	client := NewClient(
		NewImplicitCredentials(
			"abc",
		),
	)

	if client.GrantType() != "implicit" {
		t.Error("Implicit Credentials do not return implicit grant type")
	}
}

func TestClientGrantTypeClientCredentials(t *testing.T) {
	client := NewClient(
		NewClientCredentials(
			"abc",
			"def",
		),
	)

	if client.GrantType() != "client_credentials" {
		t.Error("Client Credentials do not return client_credentials grant type")
	}
}
