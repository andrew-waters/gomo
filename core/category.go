package core

// Category is a Moltin Category - https://docs.moltin.com/catalog/categories
type Category struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Meta        *struct {
		Timestamps Timestamps `json:"timestamps,omitempty"`
	} `json:"meta,omitempty"`
	Relationships *Relationships `json:"relationships,omitempty"`
}

// CategoryIncludes is possible includes for a Category
type CategoryIncludes struct {
	Products []Product `json:"products"`
}

// SetType sets the resource type on the struct
func (c *Category) SetType() {
	c.Type = categoryType
}
