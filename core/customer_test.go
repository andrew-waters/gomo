package core

import "testing"

func TestCustomerType(t *testing.T) {
	c := Customer{}
	c.SetType()
	if c.Type != "customer" {
		t.Errorf("Customer did not return correct type: `%s`", c.Type)
	}
}
