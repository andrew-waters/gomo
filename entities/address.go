package entities

// Address represents a Moltin customer address: https://docs.moltin.com/api/orders-and-customers/addresses
type Address struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Name         string `json:"name"`
	PhoneNumber  int32  `json:"phone_number,omitempty"`
	Instructions string `json:"instructions,omitempty"`
	CompanyName  string `json:"company_name"`
	Line1        string `json:"line_1"`
	Line2        string `json:"line_2"`
	City         string `json:"city"`
	County       string `json:"county"`
	Postcode     string `json:"postcode"`
	Country      string `json:"country"`
}

func (a *Address) SetType() {
	a.Type = addressType
}
