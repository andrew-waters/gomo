package core

import (
	"encoding/json"

	fv "github.com/moltin/gomo/core/field_validation"
)

// Field is a Moltin Field - https://docs.moltin.com/advanced/custom-data/fields
type Field struct {
	ID              string      `json:"id,omitempty"`
	Type            string      `json:"type"`
	Name            string      `json:"name"`
	Slug            string      `json:"slug"`
	Description     string      `json:"description"`
	FieldType       string      `json:"field_type"`
	ValidationRules []fv.Rule   `json:"validation_rules,omitempty"`
	Enabled         bool        `json:"enabled"`
	Required        bool        `json:"required"`
	Order           int32       `json:"order"`
	Unique          bool        `json:"unique"`
	Default         interface{} `json:"default,omitempty"`
	OmitNull        bool        `json:"omit_null,omitempty"`
	Meta            *struct {
		Timestamps Timestamps `json:"timestamps,omitempty"`
	} `json:"meta,omitempty"`
	Relationships *Relationships `json:"relationships,omitempty"`
}

func (f Field) MarshalJSON() ([]byte, error) {
	rules := make([]fv.Rule, 0, len(f.ValidationRules))
	for _, rule := range f.ValidationRules {
		if fv.ForType(rule) != f.FieldType {
			continue
		}
		fv.SetType(rule)
		rules = append(rules, rule)
	}
	f.ValidationRules = rules
	type Alias Field
	return json.Marshal(Alias(f))
}

func (f *Field) UnmarshalJSON(b []byte) error {
	type Alias Field
	field := struct {
		*Alias
		ValidationRules []json.RawMessage `json:"validation_rules,omitempty"`
	}{Alias: (*Alias)(f)}
	err := json.Unmarshal(b, &field)
	if err != nil {
		return err
	}
	f.ValidationRules = make([]fv.Rule, len(field.ValidationRules))
	for i, data := range field.ValidationRules {
		var vtype struct {
			Type string `json:"type"`
		}
		err := json.Unmarshal(data, &vtype)
		if err != nil {
			return err
		}
		rule := fv.TypeFor(f.FieldType, vtype.Type)
		err = json.Unmarshal(data, rule)
		if err != nil {
			return err
		}
		f.ValidationRules[i] = rule
	}
	return nil
}

// SetType sets the resource type on the struct
func (f *Field) SetType() {
	f.Type = fieldType
}
