package core

import "encoding/json"

// Customer represents a Moltin Customer Entity: https://docs.moltin.com/api/orders-and-customers/customers
type Customer struct {
	ID       string            `json:"id,omitempty"`
	Type     string            `json:"type"`
	Name     string            `json:"name"`
	Email    string            `json:"email"`
	Password *CustomerPassword `json:"password,omitempty"`
}

// CustomerPassword is a customer's password when making an API request, in
// which case the Password field should be set, or an indicator that the
// password was set when reading an API response, in which case Password
// will be empty and IsSet is indicative.
type CustomerPassword struct {
	IsSet    bool
	Password string
}

// MarshalJSON marshals a CustomerPassword into JSON
func (p CustomerPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Password)
}

// UnmarshalJSON unmarshals a CustomerPassword from JSON
func (p *CustomerPassword) UnmarshalJSON(b []byte) error {
	var isSet bool
	err := json.Unmarshal(b, &isSet)
	if err != nil {
		return err
	}
	p.IsSet = isSet
	p.Password = ""
	return nil
}

// SetType sets the resource type on the struct
func (c *Customer) SetType() {
	c.Type = customerType
}
