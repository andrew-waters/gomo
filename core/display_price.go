package core

// DisplayPriceWrapper is the format that price display responses are returned by Moltin
type DisplayPriceWrapper struct {
	WithTax    DisplayPrice `json:"with_tax"`
	WithoutTax DisplayPrice `json:"without_tax"`
}

// DisplayPrice is the content of a price for display purposes
type DisplayPrice struct {
	Amount    int    `json:"amount"`
	Currency  string `json:"currency"`
	Formatted string `json:"formatted"`
}
