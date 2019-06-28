package core

import (
	"encoding/json"
)

// Variation is a Moltin variation: https://docs.moltin.com/api/catalog/product-variations
type Variation struct {
	ID            string             `json:"id,omitempty"`
	Name          string             `json:"name"`
	Type          string             `json:"type"`
	Relationships *Relationships     `json:"relationships,omitempty"`
	Included      *VariationIncluded `json:"included,omitempty"`
}

type VariationIncluded struct {
	Options   []VariationIncludedOption `json:"options"`
	Modifiers []Modifier                `json:"modifiers"`
}

type VariationIncludedOption struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Relationships *Relationships `json:"relationships,omitempty"`
}

type VariationMatrix struct {
	Children map[string]*VariationMatrix
	Products map[string]string
}

func (m VariationMatrix) toMap() map[string]interface{} {
	out := make(map[string]interface{})
	for option, child := range m.Children {
		out[option] = child.toMap()
	}
	for option, product := range m.Products {
		out[option] = product
	}
	return out
}

func (m VariationMatrix) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.toMap())
}

func (m *VariationMatrix) UnmarshalJSON(b []byte) error {
	var array []interface{}
	err := json.Unmarshal(b, &array)
	if err == nil {
		// We receive [] if there is no variation matrix
		return nil
	}

	var mx map[string]json.RawMessage
	err = json.Unmarshal(b, &mx)
	if err != nil {
		return err
	}
	for option, jsonMsg := range mx {
		var product string
		if err := json.Unmarshal(jsonMsg, &product); err == nil {
			if m.Products == nil {
				m.Products = make(map[string]string)
			}
			m.Products[option] = product
			continue
		}
		var child VariationMatrix
		err := json.Unmarshal(jsonMsg, &child)
		if err != nil {
			return err
		}
		if m.Children == nil {
			m.Children = make(map[string]*VariationMatrix)
		}
		m.Children[option] = &child
	}
	return nil
}

// SetType sets the resource type on the struct
func (f *Variation) SetType() {
	f.Type = variationType
}
