package entities

// Customer represents a Moltin Customer Entity: https://docs.moltin.com/api/orders-and-customers/customers
type Customer struct {
	ID       string      `json:"id,omitempty"`
	Type     string      `json:"type"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Password interface{} `json:"password"`
}

// SetType sets the resource type on the struct
func (c *Customer) SetType() {
	c.Type = customerType
}
