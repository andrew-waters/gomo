package core

// Relationship is the shorthand relationship object
type Relationship struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// SingleRelationship represents the data element of a relationship
type SingleRelationship struct {
	Data Relationship `json:"data"`
}

// MulitpleRelationship represents the data element of a relationships
type MultipleRelationship struct {
	Data []Relationship `json:"data"`
}
