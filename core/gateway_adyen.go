package core

// AdyenGateway is a Moltin AdyenGateway - https://docs.moltin.com/payments/gateways/configure-adyen
type AdyenGateway struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Enabled         bool   `json:"enabled"`
	Test            bool   `json:"test"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	MerchantAccount string `json:"merchant_account"`
}

// SetType sets the resource type on the struct
func (a *AdyenGateway) SetType() {
	a.Type = gatewayType
}
