package core

import "testing"

func TestProductType(t *testing.T) {
	p := Product{}
	p.SetType()
	if p.Type != "product" {
		t.Errorf("Product did not return correct type: `%s`", p.Type)
	}
}
