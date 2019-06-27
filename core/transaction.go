package core

type Transaction struct {
	ID              string       `json:"id,omitempty"`
	Type            string       `json:"type"`
	Reference       string       `json:"reference"`
	Gateway         string       `json:"gateway"`
	Amount          int32        `json:"amount"`
	Currency        string       `json:"currency"`
	TransactionType string       `json:"transaction_type"`
	Status          string       `json:"status"`
	Relationships   *interface{} `json:"relationships,omitempty"`
}

// SetType sets the resource type on the struct
func (t *Transaction) SetType() {
	t.Type = transactionType
}
