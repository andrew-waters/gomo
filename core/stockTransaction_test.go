package core

import "testing"

func TestStockTransactionType(t *testing.T) {
	st := StockTransaction{}
	st.SetType()
	if st.Type != "stock-transaction" {
		t.Errorf("StockTransaction did not return correct type: `%s`", st.Type)
	}
}
