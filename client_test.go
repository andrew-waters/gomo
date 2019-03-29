package gomo

import (
	"log"
	"net/http"
	"os"
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

func TestClientAuthenticatesWithClientCredentials(t *testing.T) {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	if clientID == "" {
		log.Println("No client id, skipping test")
		return
	}
	if clientSecret == "" {
		log.Println("No client secret, skipping test")
		return
	}

	client := NewClient(
		NewClientCredentials(
			clientID,
			clientSecret,
		),
	)

	if err := client.Authenticate(); err != nil {
		t.Errorf("Could not authenticate with client credentials: %s", err)
	}
}

func TestClientAuthenticatesImplicitly(t *testing.T) {
	clientID := os.Getenv("CLIENT_ID")

	if clientID == "" {
		log.Println("No client id, skipping test")
		return
	}

	client := NewClient(
		NewImplicitCredentials(
			clientID,
		),
	)

	if err := client.Authenticate(); err != nil {
		t.Errorf("Could not authenticate implicitly: %s", err)
	}
}
