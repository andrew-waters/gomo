package entities

import "testing"

func TestCartItemType(t *testing.T) {
	ci := CartItem{}
	ci.SetType()
	if ci.Type != "cart_item" {
		t.Errorf("CartItem did not return correct type: `%s`", ci.Type)
	}
}
