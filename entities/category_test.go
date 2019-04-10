package entities

import "testing"

func TestCategoryType(t *testing.T) {
	c := Category{}
	c.SetType()
	if c.Type != "category" {
		t.Errorf("Category did not return correct type: `%s`", c.Type)
	}
}
