package core

import "testing"

func TestFieldType(t *testing.T) {
	f := Field{}
	f.SetType()
	if f.Type != "field" {
		t.Errorf("Field did not return correct type: `%s`", f.Type)
	}
}
