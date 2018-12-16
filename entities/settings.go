package entities

// Settings is a Moltin Setting Configuration - https://docs.moltin.com/catalog/brands
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
