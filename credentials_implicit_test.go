package gomo

import "testing"

func TestNewImplicitCredentials(t *testing.T) {
	creds := NewImplicitCredentials("abc")
	if creds.clientID != "abc" {
		t.Error("Incorrect client id")
	}
	if creds.grantType() != implicitGrantType {
		t.Error("Incorrect grant type")
	}
}
