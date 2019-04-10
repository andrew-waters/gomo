package entities

import "testing"

func TestCardConnectType(t *testing.T) {
	g := CardConnectGateway{}
	g.SetType()
	if g.Type != "gateway" {
		t.Errorf("CardConnect did not return correct type: `%s`", g.Type)
	}
}
