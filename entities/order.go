package entities

// Order represents a Moltin order: https://docs.moltin.com/api/orders-and-customers/orders
type Order struct {
	ID              string        `json:"id,omitempty"`
	Type            string        `json:"type"`
	Status          string        `json:"status"`
	Payment         string        `json:"payment"`
	Shipping        string        `json:"shipping"`
	Customer        OrderCustomer `json:"customer"`
	ShippingAddress Address       `json:"shipping_address"`
	BillingAddress  Address       `json:"billing_address"`
	Links           Links         `json:"links"`
	Meta            struct {
		DisplayPrice DisplayPriceWrapper `json:"display_price"`
		Timestamps   Timestamps          `json:"timestamps,omitempty"`
	} `json:"meta,omitempty"`
	Relationships interface{} `json:"relationships"`
}

type OrderCustomer struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// SetType sets the resource type on the struct
func (o *Event) SetType() {
	o.Type = orderType
}
