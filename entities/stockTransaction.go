package entities

// StockTransaction represents a Moltin inventory transaction: https://docs.moltin.com/api/catalog/inventory/stock-transactions
type StockTransaction struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Action    string `json:"action"`
	ProductID string `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}

// SetType sets the resource type on the struct
func (st *StockTransaction) SetType() {
	st.Type = stockTransactionType
}
