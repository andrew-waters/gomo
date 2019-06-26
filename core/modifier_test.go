package core

import "testing"

func TestModifierType(t *testing.T) {
	m := Modifier{}
	m.SetType()

	if m.Type != "modifier" {
		t.Errorf("Modifier did not return correct type: `%s`", m.Type)
	}
}
