package entities

import "testing"

func TestCartItemType(t *testing.T) {
	ci := CartItem{}
	ci.SetType()
	if ci.Type != "cart_item" {
		t.Errorf("Cart did not return correct type: `%s`", ci.Type)
	}
}
