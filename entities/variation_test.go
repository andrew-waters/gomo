package entities

import "testing"

func TestVariationType(t *testing.T) {
	v := Variation{}
	v.SetType()
	if v.Type != "product-variation" {
		t.Errorf("Variation did not return correct type: `%s`", v.Type)
	}
}
