package entities

import (
	"reflect"
	"testing"
)

func TestSettingsType(t *testing.T) {
	s := Settings{}
	s.SetType()
	if s.Type != "settings" {
		t.Errorf("Settings did not return correct type: `%s`", s.Type)
	}
}

func TestSupportedLanguages(t *testing.T) {
	s := Settings{}
	expected := []string{"ab", "af", "sq", "hy", "av", "ae", "ay", "az", "eu", "be", "bs", "bg", "ca", "co", "cs", "da", "nl", "en", "et", "fj", "fi", "fr", "de", "gd", "ga", "el", "ht", "hr", "hu", "is", "it", "lv", "lt", "lb", "no", "fa", "pl", "pt", "ro", "ru", "sk", "sl", "sm", "es", "sc", "sr", "sv", "uk", "cy"}
	if !reflect.DeepEqual(s.SupportedLanguages(), expected) {
		t.Errorf("Settings did not return the correct supported languages. Got: `%s`", s.SupportedLanguages())
	}
}
