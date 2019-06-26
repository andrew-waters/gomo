package core

import "testing"

func TestAddressType(t *testing.T) {
	a := Address{}
	a.SetType()

	if(a.Type != "address") {
		t.Errorf("Address did not return correct type: `%s`", a.Type)
	}
}
