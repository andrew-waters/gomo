package core

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestVariationMatrixMarshalJSON(t *testing.T) {
	for _, test := range []struct {
		name     string
		matrix   *VariationMatrix
		expected string
	}{
		{
			"recursive",
			&VariationMatrix{
				Children: map[string]*VariationMatrix{
					"f2f501ad-5829-4eec-bcb6-b35a48a2ea93": &VariationMatrix{
						Products: map[string]string{
							"e9af3f88-1bae-4069-a3e0-51b0abe24931": "751b294d-1b1b-4de2-9d91-6f240eb921c4",
						},
					},
				},
				Products: map[string]string{
					"90c1dd19-ff71-4eeb-b1a0-2e12a82ba357": "19f9672c-5c37-4f63-abe6-c06ccb71d999",
				},
			},
			`{"90c1dd19-ff71-4eeb-b1a0-2e12a82ba357":"19f9672c-5c37-4f63-abe6-c06ccb71d999","f2f501ad-5829-4eec-bcb6-b35a48a2ea93":{"e9af3f88-1bae-4069-a3e0-51b0abe24931":"751b294d-1b1b-4de2-9d91-6f240eb921c4"}}`,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			data, err := json.Marshal(test.matrix)
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

func TestVariationMatrixUnmarshalJSON(t *testing.T) {
	for _, test := range []struct {
		name     string
		json     string
		expected *VariationMatrix
	}{
		{
			"recursive",
			`{"90c1dd19-ff71-4eeb-b1a0-2e12a82ba357":"19f9672c-5c37-4f63-abe6-c06ccb71d999","f2f501ad-5829-4eec-bcb6-b35a48a2ea93":{"e9af3f88-1bae-4069-a3e0-51b0abe24931":"751b294d-1b1b-4de2-9d91-6f240eb921c4"}}`,
			&VariationMatrix{
				Children: map[string]*VariationMatrix{
					"f2f501ad-5829-4eec-bcb6-b35a48a2ea93": &VariationMatrix{
						Products: map[string]string{
							"e9af3f88-1bae-4069-a3e0-51b0abe24931": "751b294d-1b1b-4de2-9d91-6f240eb921c4",
						},
					},
				},
				Products: map[string]string{
					"90c1dd19-ff71-4eeb-b1a0-2e12a82ba357": "19f9672c-5c37-4f63-abe6-c06ccb71d999",
				},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			var matrix *VariationMatrix
			err := json.Unmarshal([]byte(test.json), &matrix)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(matrix, test.expected) {
				expectedS := spew.Sprintf("%#v", test.expected)
				matrixS := spew.Sprintf("%#v", matrix)
				t.Errorf(
					"\nexpected: %#v\ngot:      %#v\n",
					expectedS,
					matrixS,
				)
			}
		})
	}
}
func TestVariationType(t *testing.T) {
	v := Variation{}
	v.SetType()
	if v.Type != "product-variation" {
		t.Errorf("Variation did not return correct type: `%s`", v.Type)
	}
}
