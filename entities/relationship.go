package entities

// Relationship is the shorthand relationship object
type Relationship struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type RelationshipData struct {
	Data interface{} `json:"data"`
}
