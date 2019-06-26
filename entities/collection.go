package entities

// Collection is a Moltin Collection - https://docs.moltin.com/catalog/collections
type Collection struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Meta        struct {
		Timestamps Timestamps `json:"timestamps,omitempty"`
	} `json:"meta,omitempty"`
	Relationships struct {
		Products []Relationship `json:"products,omitempty"`
	} `json:"relationships,omitempty"`
}

// CollectionIncludes is possible includes for a Collection
type CollectionIncludes struct {
	Products []Product `json:"products"`
}

// SetType sets the resource type on the struct
func (c *Collection) SetType() {
	c.Type = collectionType
}
