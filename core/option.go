package core

// Option represents an option for selection for a single Moltin product-variation: https://docs.moltin.com/api/catalog/product-variations/options
type Option struct {
	ID          string     `json:"id,omitempty"`
	Type        string     `json:"type"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Modifiers   []Modifier `json:"modifiers"`
}

// SetType sets the resource type on the struct
func (o *Option) SetType() {
	o.Type = optionType
}
