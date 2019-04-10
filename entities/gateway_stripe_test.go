package entities

import "testing"

func TestStripeType(t *testing.T) {
	g := StripeGateway{}
	g.SetType()
	if g.Type != "gateway" {
		t.Errorf("Stripe did not return correct type: `%s`", g.Type)
	}
}
