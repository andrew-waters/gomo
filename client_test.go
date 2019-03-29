package gomo

import (
	"bytes"
	"log"
	"net/http"
	"strings"
	"testing"
)

func testClient() Client {
	return Client{
		credentials: NewClientCredentials("abc", "def"),
		APIVersion:  defaultAPIVersion,
		Endpoint:    defaultEndpoint,
		Debug:       false,
		httpClient:  &http.Client{},
		Logger:      defaultLogger,
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

func TestCustomLogger(t *testing.T) {
	client := NewClient(
		NewClientCredentials(
			"abc",
			"def",
		),
	)
	client.EnableDebug()

	logHit := false
	client.CustomLogger(func(c *Client, msg interface{}) {
		logHit = true
	})
	err := client.Authenticate()
	client.Log(err)

	if logHit == false {
		t.Errorf("log not hit")
	}
}

func TestDefaultLogger(t *testing.T) {
	client := NewClient(
		NewClientCredentials(
			"abc",
			"def",
		),
	)
	client.EnableDebug()

	var debugOnOut bytes.Buffer
	log.SetOutput(&debugOnOut)

	client.Log("a test")

	// the log time is included eg "2019/03/29 23:44:59 a test"
	// remove the date and new line then test the rest of the string
	o := strings.TrimSuffix(debugOnOut.String()[20:], "\n")

	if o != "a test" {
		t.Errorf("Logger did not work: `%s` (expected `%s`)", o, "a test")
	}

	client.DisableDebug()

	var debugOffOut bytes.Buffer
	log.SetOutput(&debugOffOut)

	client.Log("a test")
	if debugOffOut.String() != "" {
		t.Errorf("Logger did not work - did not expect any output: `%s`", debugOffOut.String())
	}
}
