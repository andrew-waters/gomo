package entities

// Brand is a Moltin Brand - https://docs.moltin.com/catalog/brands
type Brand struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Meta        struct {
		Timestamps Timestamps `json:"timestamps,omitempty"`
	}
	Relationships struct {
		Products []Relationship `json:"products,omitempty"`
	} `json:"relationships,omitempty"`
}

// BrandIncludes is possible includes for a Brand
type BrandIncludes struct {
	Products []Product `json:"products"`
}

// SetType sets the resource type on the struct
func (b *Brand) SetType() {
	b.Type = brandType
}
