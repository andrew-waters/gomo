package entities

// Event represents a Moltin event (not an event payload): https://docs.moltin.com/api/advanced/events
type Event struct {
	ID              string   `json:"id,omitempty"`
	Type            string   `json:"type"`
	Observes        []string `json:"observes"`
	Enabled         bool     `json:"enabled,omitempty"`
	IntegrationType string   `json:"integration_type"`
	Description     string   `json:"description"`
	Name            string   `json:"name"`
	Configuration   struct {
		URL       string `json:"url"`
		SecretKey string `json:"secret_key,omitempty"`
	} `json:"configuration"`
}

// SetType sets the resource type on the struct
func (e *Event) SetType() {
	e.Type = eventType
}
