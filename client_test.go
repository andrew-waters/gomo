package gomo

import (
	"net/http"
	"testing"
)

func testClient() Client {
	return Client{
		credentials: NewClientCredentials("abc", "def"),
		APIVersion:  DefaultAPIVersion,
		Endpoint:    DefaultEndpoint,
		Debug:       false,
		httpClient:  &http.Client{},
	}
}

func TestClientDefaults(t *testing.T) {
	c := testClient()
	expectedEndpoint := "https://api.moltin.com"
	if c.Endpoint != expectedEndpoint {
		t.Errorf("Incorrect default endpoint: %s (expected %s)", c.Endpoint, expectedEndpoint)
	}
	expectedAPIVersion := "v2"
	if c.APIVersion != expectedAPIVersion {
		t.Errorf("Incorrect API version: %s (expected %s)", c.APIVersion, expectedAPIVersion)
	}
}