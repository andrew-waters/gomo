package entities

type Event struct {
	ID              string   `json:"id,omitempty"`
	Type            string   `json:"type"`
	Observes        []string `json:"observes"`
	Enabled         bool     `json:"enabled,omitempty"`
	IntegrationType string   `json:"integration_type"`
	Description     string   `json:"description"`
	Name            string   `json:"name"`
	Configuration   struct {
		Url       string `json:"url"`
		SecretKey string `json:"secret_key,omitempty"`
	} `json:"configuration"`
}

func (e *Event) SetType() {
	e.Type = eventType
}
