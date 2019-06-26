package entities

// Order represents a Moltin order: https://docs.moltin.com/api/orders-and-customers/orders
type Order struct {
	ID              string               `json:"id,omitempty"`
	Type            string               `json:"type"`
	Status          string               `json:"status"`
	Payment         string               `json:"payment"`
	Shipping        string               `json:"shipping"`
	Customer        OrderCustomer        `json:"customer"`
	ShippingAddress OrderShippingAddress `json:"shipping_address"`
	BillingAddress  OrderBillingAddress  `json:"billing_address"`
	Links           Links                `json:"links"`
	Meta            OrderMeta            `json:"meta,omitempty"`
	Relationships   interface{}          `json:"relationships,omitempty"`
}

// OrderMeta represents the meta object for a Moltin order
type OrderMeta struct {
	DisplayPrice DisplayPriceWrapper `json:"display_price"`
	Timestamps   Timestamps          `json:"timestamps,omitempty"`
}

// OrderCustomer represents a Moltin customer object for a Moltin order (can be ID or Name and Email)
type OrderCustomer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// OrderBillingAddress represents a Moltin billing address for a Moltin order
type OrderBillingAddress struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CompanyName string `json:"company_name"`
	Line1       string `json:"line_1"`
	Line2       string `json:"line_2"`
	City        string `json:"city"`
	PostCode    string `json:"postcode"`
	County      string `json:"county"`
	Country     string `json:"country"`
}

// OrderShippingAddress represents a Moltin shipping address for a Moltin order
type OrderShippingAddress struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PhoneNumber  string `json:"phone_number"`
	CompanyName  string `json:"company_name"`
	Line1        string `json:"line_1"`
	Line2        string `json:"line_2"`
	City         string `json:"city"`
	PostCode     string `json:"postcode"`
	County       string `json:"county"`
	Country      string `json:"country"`
	Instructions string `json:"instructions"`
}

// SetType sets the resource type on the struct
func (o *Order) SetType() {
	o.Type = orderType
}
