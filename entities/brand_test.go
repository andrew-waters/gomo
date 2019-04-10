package entities

import "testing"

func TestBrandType(t *testing.T) {
	b := Brand{}
	b.SetType()
	if b.Type != "brand" {
		t.Errorf("Brand did not return correct type: `%s`", b.Type)
	}
}
