package core

// Flow is a Moltin Flow - https://docs.moltin.com/advanced/custom-data/flows
type Flow struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Meta        *struct {
		Timestamps Timestamps `json:"timestamps,omitempty"`
	} `json:"meta,omitempty"`
	Relationships *struct {
		Fields []Field `json:"fields,omitempty"`
	} `json:"relationships,omitempty"`
}

// SetType sets the resource type on the struct
func (f *Flow) SetType() {
	f.Type = flowType
}
