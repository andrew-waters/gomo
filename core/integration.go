package core

// Integration is a Moltin Integration - https://docs.moltin.com/advanced/events
type Integration struct {
	ID              string                   `json:"id,omitempty"`
	Type            string                   `json:"type"`
	Name            string                   `json:"name"`
	Description     string                   `json:"description"`
	Enabled         bool                     `json:"enabled"`
	IntegrationType string                   `json:"integration_type"`
	Observes        []string                 `json:"observes"`
	Configuration   IntegrationConfiguration `json:"configuration"`
	Meta            *struct {
		Timestamps Timestamps `json:"timestamps,omitempty"`
	} `json:"meta,omitempty"`
}

// IntegrationConfiguration specifies a webhook configuration object
type IntegrationConfiguration struct {
	URL       string `json:"url"`
	SecretKey string `json:"secret_key,omitempty"`
}

// SetType sets the resource type on the struct
func (i *Integration) SetType() {
	i.Type = integrationType
}
