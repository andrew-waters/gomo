package core

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"

	fv "github.com/moltin/gomo/core/field_validation"
)

func TestFieldMarshalJSON(t *testing.T) {
	for _, test := range []struct {
		name     string
		field    Field
		expected string
	}{
		{
			"integer/enum",
			Field{
				FieldType: "integer",
				ValidationRules: []fv.Rule{
					&fv.StringEnum{
						Options: []string{"one", "two"},
					},
					&fv.IntegerEnum{
						Options: []int{1, 2},
					},
				},
			},
			`{"type":"","name":"","slug":"","description":"","field_type":"integer","validation_rules":[{"type":"enum","options":[1,2]}],"enabled":false,"required":false,"order":0,"unique":false}`,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			data, err := json.Marshal(test.field)
			if err != nil {
				t.Fatal(err)
			}
			if string(data) != test.expected {
				t.Errorf(
					"\nexpected: %s\ngot:       %s\n",
					test.expected,
					string(data),
				)
			}
		})
	}
}

func TestFieldUnmarshalJSON(t *testing.T) {
	for _, test := range []struct {
		name     string
		json     string
		expected Field
	}{
		{
			"integer/enum",
			`{"type":"","name":"","slug":"","description":"","field_type":"integer","validation_rules":[{"type":"enum","options":[1,2]}],"enabled":false,"required":false,"order":0,"unique":false}`,
			Field{
				FieldType: "integer",
				ValidationRules: []fv.Rule{
					&fv.IntegerEnum{
						Type:    "enum",
						Options: []int{1, 2},
					},
				},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			var field Field
			err := json.Unmarshal([]byte(test.json), &field)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(field, test.expected) {
				t.Errorf(
					"\nexpected: %s\ngot:       %s\n",
					spew.Sdump(test.expected),
					spew.Sdump(field),
				)
			}
		})
	}
}

func TestFieldType(t *testing.T) {
	f := Field{}
	f.SetType()
	if f.Type != "field" {
		t.Errorf("Field did not return correct type: `%s`", f.Type)
	}
}
