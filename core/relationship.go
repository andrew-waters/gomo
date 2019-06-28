package core

import (
	"encoding/json"
	"strings"
)

// Relationship is the shorthand relationship object
type Relationship struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// RelationshipContainer contains one-to-one or on-to-many relationships.
// One-to-one relationships should have exactly one Relationship and
// OneToOne set to true.
type RelationshipContainer struct {
	Relationships []Relationship
	OneToOne      bool
}

// Relationships represents a relaionships object
type Relationships map[string]RelationshipContainer

// MarshalJSON marshals Relationships into JSON
func (r Relationships) MarshalJSON() ([]byte, error) {
	out := make(map[string]interface{})
	for key, rel := range r {
		if rel.OneToOne {
			if len(rel.Relationships) == 0 {
				continue
			}
			out[key] = struct {
				Data Relationship `json:"data"`
			}{
				Data: rel.Relationships[0],
			}
		} else {
			out[key] = struct {
				Data []Relationship `json:"data"`
			}{
				Data: rel.Relationships,
			}
		}
	}
	return json.Marshal(out)
}

func isJSONArray(j []byte) bool {
	s := string(j)
	return strings.HasPrefix(s, "[")
}

// UnmarshalJSON unmarshals Relationships from JSON
func (r *Relationships) UnmarshalJSON(b []byte) error {
	var rels map[string]struct {
		Data json.RawMessage `json:"data"`
	}
	err := json.Unmarshal(b, &rels)
	if err != nil {
		return err
	}
	for key, data := range rels {
		var rc RelationshipContainer
		if isJSONArray(data.Data) {
			var relationships []Relationship
			err = json.Unmarshal(data.Data, &relationships)
			if err != nil {
				return err
			}
			rc.Relationships = relationships
			rc.OneToOne = false
		} else {
			var relationship Relationship
			err := json.Unmarshal(data.Data, &relationship)
			if err != nil {
				return err
			}
			rc.Relationships = []Relationship{relationship}
			rc.OneToOne = true
		}
		if *r == nil {
			*r = make(Relationships)
		}
		(*r)[key] = rc
	}
	return nil
}
