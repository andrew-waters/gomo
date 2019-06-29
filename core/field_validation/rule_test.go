package field_validation_test

import (
	"testing"

	fv "github.com/moltin/gomo/core/field_validation"
)

func TestForType(t *testing.T) {
	for _, test := range []struct {
		name     string
		rule     fv.Rule
		expected string
	}{
		{
			"integer/enum",
			&fv.IntegerEnum{
				Type:    "enum",
				Options: []int{1, 2, 3},
			},
			"integer",
		},
		{
			"relationship/one-to-many",
			&fv.RelationshipOneToMany{
				Type: "one-to-many",
				To:   "foo",
			},
			"relationship",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			forType := fv.ForType(test.rule)
			if forType != test.expected {
				t.Errorf(
					"expected: %s, got %s",
					test.expected,
					forType,
				)
			}
		})
	}
}

func TestValidationType(t *testing.T) {
	for _, test := range []struct {
		name     string
		rule     fv.Rule
		expected string
	}{
		{
			"integer/enum",
			&fv.IntegerEnum{
				Options: []int{1, 2, 3},
			},
			"enum",
		},
		{
			"string/slug",
			&fv.StringPlain{
				Type: "slug",
			},
			"slug",
		},
		{
			"relationship/one-to-many",
			&fv.RelationshipOneToMany{
				To: "foo",
			},
			"one-to-many",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			forType := fv.ValidationType(test.rule)
			if forType != test.expected {
				t.Errorf(
					"expected: %s, got %s",
					test.expected,
					forType,
				)
			}
		})
	}
}
