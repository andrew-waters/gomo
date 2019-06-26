package core

import "testing"

func TestPromotionType(t *testing.T) {
	p := Promotion{}
	p.SetType()
	if p.Type != "promotion" {
		t.Errorf("Promotion did not return correct type: `%s`", p.Type)
	}
}
