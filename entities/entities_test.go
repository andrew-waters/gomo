package entities

import (
	"reflect"
	"testing"
)

var typeTests = []struct {
	typeName string
	expected string
}{
	{"Brand", "brand"},
	{"Category", "category"},
	{"Collection", "collection"},
	{"Currency", "currency"},
	{"Field", "field"},
	{"File", "file"},
	{"Flow", "flow"},
	{"Integration", "integration"},
	{"Product", "product"},
	{"Settings", "settings"},
}

func TestTypes(t *testing.T) {
	var e interface{}
	for _, test := range typeTests {
		switch test.typeName {
		case "Brand":
			e = Brand{}
		case "Category":
			e = Category{}
		case "Collection":
			e = Collection{}
		case "Currency":
			e = Currency{}
		case "Field":
			e = Field{}
		case "File":
			e = File{}
		case "Flow":
			e = Flow{}
		case "Integration":
			e = Integration{}
		case "Product":
			e = Product{}
		case "Settings":
			e = Settings{}
		}
		// set the resource type if the entity has the SetType method
		if e, ok := e.(interface{ SetType() }); ok {
			e.SetType()
			val := reflect.ValueOf(e)
			if !val.IsValid() {
				t.Errorf("%s if not valid", val)
			}
			typeName := val.Elem().FieldByName("Type")
			if !typeName.IsValid() {
				t.Errorf("%s did not return correct type", test.typeName)
			}
			if !ok {
				t.Errorf("%s does not have `SetType` func", test.typeName)
			}
		}
	}
}
