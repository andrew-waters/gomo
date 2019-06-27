package core

// Variation is a Moltin variation: https://docs.moltin.com/api/catalog/product-variations
type Variation struct {
	ID            string                  `json:"id,omitempty"`
	Name          string                  `json:"name"`
	Type          string                  `json:"type"`
	Relationships *VariationRelationships `json:"relationships,omitempty"`
	Included      *interface{}            `json:"included,omitempty"`
}

// VariationRelationships hold the relationships between a variations and its options
type VariationRelationships struct {
	Options MultipleRelationship `json:"options"`
}

// SetType sets the resource type on the struct
func (f *Variation) SetType() {
	f.Type = variationType
}
