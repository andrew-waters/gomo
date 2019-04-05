package entities

// Settings is a Moltin Setting Configuration - https://docs.moltin.com/advanced/settings
type Settings struct {
	Type                string   `json:"type"`
	PageLength          int      `json:"page_length"`
	ListChildProducts   bool     `json:"list_child_products"`
	AdditionalLanguages []string `json:"additional_languages"`
}

// SetType sets the resource type on the struct
func (s *Settings) SetType() {
	s.Type = settingsType
}

// SupportedLanguages returns the languages that are supported on the additional_languages field
func (s *Settings) SupportedLanguages() []string {
	return []string{"ab", "af", "sq", "hy", "av", "ae", "ay", "az", "eu", "be", "bs", "bg", "ca", "co", "cs", "da", "nl", "en", "et", "fj", "fi", "fr", "de", "gd", "ga", "el", "ht", "hr", "hu", "is", "it", "lv", "lt", "lb", "no", "fa", "pl", "pt", "ro", "ru", "sk", "sl", "sm", "es", "sc", "sr", "sv", "uk", "cy"}
}
