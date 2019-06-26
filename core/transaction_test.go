package core

import "testing"

func TestTransactionType(t *testing.T) {
	ts := Transaction{}
	ts.SetType()
	if ts.Type != "transaction" {
		t.Errorf("Transaction did not return correct type: `%s`", ts.Type)
	}
}
