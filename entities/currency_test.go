package entities

import "testing"

func TestCurrencyType(t *testing.T) {
	c := Currency{}
	c.SetType()
	if c.Type != "currency" {
		t.Errorf("Currency did not return correct type: `%s`", c.Type)
	}
}
