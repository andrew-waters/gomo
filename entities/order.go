package entities

type Order struct {
	ID              string               `json:"id,omitempty"`
	Type            string               `json:"type"`
	Status          string               `json:"status"`
	Payment         string               `json:"payment"`
	Shipping        string               `json:"shipping"`
	Customer        OrderCustomer        `json:"customer"`
	ShippingAddress OrderShippingAddress `json:"shipping_address"`
	BillingAddress  OrderAddress         `json:"billing_address"`
	Meta            struct {
		DisplayPrice DisplayPriceWrapper `json:"display_price"`
		Timestamps   Timestamps          `json:"timestamps,omitempty"`
	} `json:"meta"`
	Relationships struct {
		Items struct {
			Data []Relationship `json:"data"`
		} `json:"items"`
	} `json:"relationships"`
}

type OrderCustomer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type OrderShippingAddress struct {
	OrderAddress
	Instructions string `json:"instructions"`
}

type OrderAddress struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	CompanyName string `json:"company_name"`
	Line1       string `json:"line_1"`
	Line2       string `json:"line_2"`
	City        string `json:"city"`
	Postcode    string `json:"postcode"`
	County      string `json:"county"`
	Country     string `json:"country"`
}

const (
	orderType = "order"
)

// SetType sets the resource type on the struct
func (c *Order) SetType() {
	c.Type = orderType
}
