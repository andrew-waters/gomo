package entities

import "testing"

func TestBraintreeType(t *testing.T) {
	g := BraintreeGateway{}
	g.SetType()
	if g.Type != "gateway" {
		t.Errorf("Braintree did not return correct type: `%s`", g.Type)
	}
}
