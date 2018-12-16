package entities

type DisplayPriceWrapper struct {
	WithTax    DisplayPrice `json:"with_tax"`
	WithoutTax DisplayPrice `json:"without_tax"`
}

type DisplayPrice struct {
	Amount    int    `json:"amount"`
	Currency  string `json:"currency"`
	Formatted string `json:"formatted"`
}
