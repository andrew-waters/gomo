package core

// BraintreeGateway is a Moltin BraintreeGateway - https://docs.moltin.com/payments/gateways/configure-braintree
type BraintreeGateway struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Enabled     bool   `json:"enabled"`
	Environment string `json:"environment"`
	PublicKey   string `json:"public_key"`
	PrivateKey  string `json:"private_key"`
	MerchantID  string `json:"merchant_id"`
}

// SetType sets the resource type on the struct
func (b *BraintreeGateway) SetType() {
	b.Type = gatewayType
}
