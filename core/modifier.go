package core

// Modifier represents a modifier for a variation option: https://docs.moltin.com/api/catalog/product-variations/modifiers
type Modifier struct {
	ID           string      `json:"id,omitempty"`
	Type         string      `json:"type"`
	ModifierType string      `json:"modifier_type"`
	Value        interface{} `json:"value"`
}

// SetType sets the resource type on the struct
func (m *Modifier) SetType() {
	m.Type = modifierType
}
