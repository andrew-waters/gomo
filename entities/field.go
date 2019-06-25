package entities

// Field is a Moltin Field - https://docs.moltin.com/advanced/custom-data/fields
type Field struct {
	ID              string        `json:"id,omitempty"`
	Type            string        `json:"type"`
	Name            string        `json:"name"`
	Slug            string        `json:"slug"`
	Description     string        `json:"description"`
	FieldType       string        `json:"field_type"`
	ValidationRules []interface{} `json:"validation_rules,omitempty"`
	Enabled         bool          `json:"enabled"`
	Required        bool          `json:"required"`
	Order           int32         `json:"order"`
	Unique          bool          `json:"unique"`
	Default         interface{}   `json:"default,omitempty"`
	OmitNull        bool          `json:"omit_null,omitempty"`
	Meta            struct {
		Timestamps Timestamps `json:"timestamps,omitempty"`
	}
	Relationships FieldToFlowRelationship `json:"relationships,omitempty"`
}

// FieldToFlowRelationship is the container for flow relationship data
type FieldToFlowRelationship struct {
	Flow RelationshipData `json:"flow,omitempty"`
}

// SetType sets the resource type on the struct
func (f *Field) SetType() {
	f.Type = fieldType
}
