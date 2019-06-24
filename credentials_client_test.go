package gomo

import "testing"

func TestClientCredentials(t *testing.T) {
	creds := clientCredentials{}
	if creds.grantType() != clientCredentialsGrantType {
		t.Error("Incorrect grant type")
	}
}
