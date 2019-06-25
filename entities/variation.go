package entities

// Variation is a Moltin variation: https://docs.moltin.com/api/catalog/product-variations
type Variation struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Type          string                 `json:type`
	Relationships VariationRelationships `json:"relationships"`
	Included      interface{}            `json:"included"`
}

// VariationRelationships hold the relationships between a variations and its options
type VariationRelationships struct {
	Data []VariationOptionRelationship `json:"data"`
}

// VariationOptionRelationship represents an option which is related to a variation
type VariationOptionRelationship struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// SetType sets the resource type on the struct
func (f *Variation) SetType() {
	f.Type = variationType
}
