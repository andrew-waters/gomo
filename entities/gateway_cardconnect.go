package entities

// CardConnectGateway is a Moltin CardConnectGateway - https://docs.moltin.com/payments/gateways/configure-cardconnect
type CardConnectGateway struct {
	Type       string `json:"type"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	Enabled    bool   `json:"enabled"`
	Test       bool   `json:"test"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	MerchantID string `json:"merchant_id"`
}

// SetType sets the resource type on the struct
func (c *CardConnectGateway) SetType() {
	c.Type = gatewayType
}
