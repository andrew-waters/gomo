package entities

import "testing"

func TestCartType(t *testing.T) {
	c := Cart{}
	c.SetType()
	if c.Type != "cart" {
		t.Errorf("Cart did not return correct type: `%s`", c.Type)
	}
}