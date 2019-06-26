package core

import "testing"

func TestCollectionType(t *testing.T) {
	c := Collection{}
	c.SetType()
	if c.Type != "collection" {
		t.Errorf("Collection did not return correct type: `%s`", c.Type)
	}
}
