package entities

// Cart represents a Moltin cart: https://docs.moltin.com/api/carts-and-checkout/carts
type Cart struct {
	ID    string   `json:"id,omitempty"`
	Type  string   `json:"type"`
	Links Links    `json:"links,omitempty"`
	Meta  CartMeta `json:"meta,omitempty"`
}

// CartMeta represents the Meta object for a Moltin cart
type CartMeta struct {
	DisplayPrice DisplayPriceWrapper `json:"display_price,omitempty"`
	Timestamps   Timestamps          `json:"timestamps,omitempty"`
}

// SetType sets the resource type on the struct
func (c *Cart) SetType() {
	c.Type = cartType
}
