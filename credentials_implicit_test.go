package gomo

import "testing"

func TestImplicitCredentials(t *testing.T) {
	creds := implicitCredentials{}
	if creds.grantType() != implicitGrantType {
		t.Error("Incorrect grant type")
	}
}
