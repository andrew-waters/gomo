package entities

import "testing"

func TestAdyenType(t *testing.T) {
	g := AdyenGateway{}
	g.SetType()
	if g.Type != "gateway" {
		t.Errorf("Adyen did not return correct type: `%s`", g.Type)
	}
}
