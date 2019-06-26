package core

import "testing"

func TestIntegrationType(t *testing.T) {
	i := Integration{}
	i.SetType()
	if i.Type != "integration" {
		t.Errorf("Integration did not return correct type: `%s`", i.Type)
	}
}
