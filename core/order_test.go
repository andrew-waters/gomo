package core

import "testing"

func TestOrderType(t *testing.T) {
	o := Order{}
	o.SetType()

	if(o.Type != "order") {
		t.Errorf("Order did not return correct type: `%s`", o.Type)
	}
}
