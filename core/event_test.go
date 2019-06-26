package core

import "testing"

func TestEventType(t *testing.T) {
	e := Event{}
	e.SetType()

	if(e.Type != "event") {
		t.Errorf("Event did not return correct type: `%s`", e.Type)
	}
}
