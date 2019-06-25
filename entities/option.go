package entities

// Option represents an option for selection for a single product-variation:
type Option struct {
	ID          string     `json:"id"`
	Type        string     `json:"type"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Modifiers   []Modifier `json:"modifiers"`
}

// SetType sets the resource type on the struct
func (o *Option) SetType() {
	o.Type = optionType
}
