package core

import (
	"encoding/json"
	"fmt"
)

// Modifier represents a modifier for a variation option:
type Modifier struct {
	ID           string        `json:"id,omitempty"`
	Type         string        `json:"type"`
	ModifierType string        `json:"modifier_type"`
	Value        ModifierValue `json:"value"`
}

// ModifierValue repesents a number of possible structures depending
// on the modifier type
type ModifierValue struct {
	ModifierValuePlain
	ModifierValueBuilder
	ModifierValuePrice
}

// ModifierValuePlain is used for modifier types: description_equals,
// description_prepend, description_append, commoditytype, name_equals,
// name_prepend, name_append, slug_equals, slug_prepend, slug_append,
// sku_equals, sku_prepend, sku_append, status
type ModifierValuePlain struct {
	Value string
}

// ModifierValueBuilder is used for modifier types: sku_builder, slug_builder
type ModifierValueBuilder struct {
	Seek string `json:"seek"`
	Set  string `json:"set"`
}

// ModifierValuePrice is used for modifier types: price_increment,
// price_decrement, price_equals
type ModifierValuePrice struct {
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	IncludesTax bool   `json:"includes_tax"`
}

// SetType sets the resource type on the struct
func (m *Modifier) SetType() {
	m.Type = modifierType
}

type modifierValueType int

const (
	modifierValueTypeInvalid = iota
	modifierValueTypePlain
	modifierValueTypeBuilder
	modifierValueTypePrice
)

func parseModifierType(mtype string) (t modifierValueType) {
	switch mtype {
	case
		"description_equals",
		"description_prepend",
		"description_append",
		"commoditytype",
		"name_equals",
		"name_prepend",
		"name_append",
		"slug_equals",
		"slug_prepend",
		"slug_append",
		"sku_equals",
		"sku_prepend",
		"sku_append",
		"status":
		t = modifierValueTypePlain
	case
		"sku_builder",
		"slug_builder":
		t = modifierValueTypeBuilder
	case
		"price_increment",
		"price_decrement",
		"price_equals":
		t = modifierValueTypePrice
	default:
		t = modifierValueTypeInvalid
	}
	return
}

// MarshalJSON marshals a Modifier into JSON
func (m Modifier) MarshalJSON() ([]byte, error) {
	mtype := m.ModifierType
	mod := struct {
		ID           string      `json:"id,omitempty"`
		Type         string      `json:"type"`
		ModifierType string      `json:"modifier_type"`
		Value        interface{} `json:"value"`
	}{
		ID:           m.ID,
		Type:         m.Type,
		ModifierType: mtype,
	}
	switch parseModifierType(mtype) {
	case modifierValueTypePlain:
		mod.Value = m.Value.ModifierValuePlain.Value
	case modifierValueTypeBuilder:
		mod.Value = m.Value.ModifierValueBuilder
	case modifierValueTypePrice:
		mod.Value = m.Value.ModifierValuePrice
	default:
		return nil, fmt.Errorf("unknown modifier type: %s", mtype)
	}
	return json.Marshal(mod)
}

// UnmarshalJSON unmarshals a Modifier from JSON
func (m *Modifier) UnmarshalJSON(b []byte) error {
	mod := struct {
		ID           string `json:"id,omitempty"`
		Type         string `json:"type"`
		ModifierType string `json:"modifier_type"`
	}{}
	err := json.Unmarshal(b, &mod)
	if err != nil {
		return err
	}
	mtype := mod.ModifierType
	m.ID = mod.ID
	m.Type = mod.Type
	m.ModifierType = mtype
	switch parseModifierType(mtype) {
	case modifierValueTypePlain:
		value := struct {
			Value string `json:"value"`
		}{}
		err := json.Unmarshal(b, &value)
		if err != nil {
			return err
		}
		m.Value.ModifierValuePlain.Value = value.Value
	case modifierValueTypeBuilder:
		value := struct {
			Value ModifierValueBuilder `json:"value"`
		}{}
		err := json.Unmarshal(b, &value)
		if err != nil {
			return err
		}
		m.Value.ModifierValueBuilder = value.Value
	case modifierValueTypePrice:
		value := struct {
			Value ModifierValuePrice `json:"value"`
		}{}
		err := json.Unmarshal(b, &value)
		if err != nil {
			return err
		}
		m.Value.ModifierValuePrice = value.Value
	default:
		return fmt.Errorf("unknown modifier type: %s", mtype)
	}
	return nil
}
