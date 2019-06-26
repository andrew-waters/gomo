package core

// Promotion represents a Moltin promotion: https://docs.moltin.com/api/carts-and-checkout/promotions
type Promotion struct {
	ID            string          `json:"id,omitempty"`
	Type          string          `json:"type"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	Enabled       bool            `json:"enabled"`
	PromotionType string          `json:"promotion_type"`
	Schema        PromotionSchema `json:"schema"`
	Start         string          `json:"start"`
	End           string          `json:"end"`
	Meta          struct {
		Timestamps Timestamps `json:"timestamps,omitempty"`
	} `json:"meta,omitempty"`
}

type PromotionSchema struct {
	Currencies []PromotionCurrency `json:"currencies"`
}

type PromotionCurrency struct {
	Currency   string  `json:"currency"`
	Amount     int32   `json:"amount,omitempty"`
	Percentage float32 `json:"percentage,omitempty"`
}

// SetType sets the resource type on the struct
func (p *Promotion) SetType() {
	p.Type = promotionType
}
