package entities

import "testing"

func TestFlowType(t *testing.T) {
	f := Flow{}
	f.SetType()
	if f.Type != "flow" {
		t.Errorf("Flow did not return correct type: `%s`", f.Type)
	}
}
