package gomo

import (
	"log"
	"os"
	"testing"
)

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
