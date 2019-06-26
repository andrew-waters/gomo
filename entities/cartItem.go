package entities

// CartItem represents a Moltin cart item: https://docs.moltin.com/api/carts-and-checkout/carts/cart-items
type CartItem struct {
	ID          string            `json:"id,omitempty"`
	Type        string            `json:"type"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Sku         string            `json:"sku"`
	Image       CartItemImage     `json:"image,omitempty"`
	Quantity    int32             `json:"quantity"`
	UnitPrice   CartItemUnitPrice `json:"unit_price"`
	Value       CartItemValue     `json:"value"`
	Links       Links             `json:"links,omitempty"`
	Meta        CartItemMeta      `json:"meta,omitempty"`
}

// CartItemImage represents the image object for a Moltin cart item
type CartItemImage struct {
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
	Href     string `json:"href"`
}

// CartItemUnitPrice represents the unit price object for a Moltin cart item
type CartItemUnitPrice struct {
	Amount      int32  `json:"amount"`
	Currency    string `json:"currenct"`
	IncludesTax bool   `json:"includes_tax"`
}

// CartItemValue  represents the value object for a Moltin cart item
type CartItemValue struct {
	Amount      int32  `json:"amount"`
	Currency    string `json:"currenct"`
	IncludesTax bool   `json:"includes_tax"`
}

// CartItemMeta represents the Meta object for a Moltin cart item
type CartItemMeta struct {
	DisplayPrice DisplayPriceWrapper `json:"display_price,omitempty"`
	Timestamps   Timestamps          `json:"timestamps,omitempty"`
}

// SetType sets the resource type on the struct
func (ci *CartItem) SetType() {
	ci.Type = cartItemType
}
