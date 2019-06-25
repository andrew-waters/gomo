package entities

import "testing"

func TestOptionType(t *testing.T) {
	o := Option{}
	o.SetType()

	if o.Type != "option" {
		t.Errorf("Option did not return correct type: `%s`", o.Type)
	}
}
