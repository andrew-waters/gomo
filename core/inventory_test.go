package core

import "testing"

func TestInventoryType(t *testing.T) {
	i := Inventory{}
	i.SetType()

	if i.Type != "inventory" {
		t.Errorf("Inventory did not return correct type: `%s`", i.Type)
	}
}
