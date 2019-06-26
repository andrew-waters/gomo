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
	Meta            OrderMeta     `json:"meta,omitempty"`
	Relationships   interface{}   `json:"relationships"`
}

// OrderMeta represents the meta object for a Moltin order
type OrderMeta struct {
	DisplayPrice DisplayPriceWrapper `json:"display_price"`
	Timestamps   Timestamps          `json:"timestamps,omitempty"`
}

// OrderCustomer represents a Moltin customer object for a Moltin order (can be ID or Name and Email)
type OrderCustomer struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// SetType sets the resource type on the struct
func (o *Order) SetType() {
	o.Type = orderType
}
