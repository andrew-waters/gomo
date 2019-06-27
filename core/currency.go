package core

// Currency is a Moltin Currency - https://docs.moltin.com/advanced/currencies
type Currency struct {
	ID                string  `json:"id,omitempty"`
	Type              string  `json:"type"`
	Code              string  `json:"code"`
	ExchangeRate      float64 `json:"exchange_rate"`
	Format            string  `json:"format"`
	ThousandSeparator string  `json:"thousand_separator"`
	DecimalPoint      string  `json:"decimal_point"`
	DecimalPlaces     int64   `json:"decimal_places"`
	Default           bool    `json:"default"`
	Enabled           bool    `json:"enabled"`
	Links             *Links  `json:"links"`
	Meta              *struct {
		Timestamps Timestamps `json:"timestamps,omitempty"`
	} `json:"meta,omitempty"`
}

// SetType sets the resource type on the struct
func (c *Currency) SetType() {
	c.Type = currencyType
}
