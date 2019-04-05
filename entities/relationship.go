package entities

// Relationship is the shorthand relationship object
type Relationship struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// RelationshipData represents the data element of a relationship
type RelationshipData struct {
	Data interface{} `json:"data"`
}
