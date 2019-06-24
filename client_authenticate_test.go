package gomo

import (
	"os"
	"testing"
)

func TestClientAuthenticatesWithClientCredentials(t *testing.T) {
	clientID := os.Getenv("MOLTIN_CLIENT_ID")
	clientSecret := os.Getenv("MOLTIN_CLIENT_SECRET")

	if clientID == "" {
		t.Skip("No client id, skipping test")
	}
	if clientSecret == "" {
		t.Skip("No client secret, skipping test")
	}

	client := NewClient(
		ClientCredentials(
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
		t.Skip("No client id, skipping test")
	}

	client := NewClient(
		ImplicitCredentials(
			clientID,
		),
	)

	if err := client.Authenticate(); err != nil {
		t.Errorf("Could not authenticate implicitly: %s", err)
	}
}
