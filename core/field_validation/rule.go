package field_validation

import (
	"reflect"
	"regexp"
	"strings"
)

type Rule interface {
	setType(string)
}

var typeRE = regexp.MustCompile(`[A-Z][a-z]+`)

func parseType(s string) (string, string) {
	splitB := typeRE.FindAll([]byte(s), -1)
	split := make([]string, len(splitB))
	for i, b := range splitB {
		split[i] = strings.ToLower(string(b))
	}
	return split[0], strings.Join(split[1:], "-")
}

func ForType(r Rule) string {
	forType, _ := parseType(reflect.ValueOf(r).Elem().Type().Name())
	return forType
}

func ValidationType(r Rule) string {
	_, validationType := parseType(reflect.ValueOf(r).Elem().Type().Name())
	if validationType != "plain" {
		return validationType
	}
	return reflect.ValueOf(r).Elem().FieldByName("Type").String()
}

func SetType(r Rule) {
	vt := ValidationType(r)
	r.setType(vt)
}

func TypeFor(fieldType, validationType string) Rule {
	var rule Rule
	factory := map[string]map[string]func(){
		"string": map[string]func(){
			"enum": func() { rule = &StringEnum{} },
		},
		"integer": map[string]func(){
			"between": func() { rule = &IntegerBetween{} },
			"enum":    func() { rule = &IntegerEnum{} },
		},
		"float": map[string]func(){
			"between": func() { rule = &FloatBetween{} },
			"enum":    func() { rule = &FloatEnum{} },
		},
		"date": map[string]func(){
			"enum": func() { rule = &DateEnum{} },
		},
		"relationship": map[string]func(){
			"one-to-many": func() {
				rule = &RelationshipOneToMany{}
			},
			"one-to-one": func() {
				rule = &RelationshipOneToOne{}
			},
		},
	}
	typeValidations, ok := factory[fieldType]
	if !ok {
		return nil
	}
	rfunc, ok := typeValidations[validationType]
	if ok {
		rfunc()
	} else {
		rule = &StringPlain{}
	}
	return rule
}
