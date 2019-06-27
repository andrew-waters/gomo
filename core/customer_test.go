package core

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCustomerType(t *testing.T) {
	c := Customer{}
	c.SetType()
	if c.Type != "customer" {
		t.Errorf("Customer did not return correct type: `%s`", c.Type)
	}
}

func TestCustomerMarshalJSON(t *testing.T) {
	for _, test := range []struct {
		name     string
		customer Customer
		expected string
	}{
		{
			"no password",
			Customer{
				Type:  "customer",
				Name:  "Ron Swanson",
				Email: "ron@swanson.com",
			},
			`{"type":"customer","name":"Ron Swanson","email":"ron@swanson.com"}`,
		},
		{
			"with password",
			Customer{
				Type:  "customer",
				Name:  "Ron Swanson",
				Email: "ron@swanson.com",
				Password: &CustomerPassword{
					Password: "password123",
				},
			},
			`{"type":"customer","name":"Ron Swanson","email":"ron@swanson.com","password":"password123"}`,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			data, err := json.Marshal(test.customer)
			if err != nil {
				t.Fatal(err)
			}
			if string(data) != test.expected {
				t.Errorf(
					"\nexpected: %s\ngot:       %s\n",
					test.expected,
					string(data),
				)
			}
		})
	}
}

func TestCustomerUnmarshalJSON(t *testing.T) {
	for _, test := range []struct {
		name     string
		json     string
		expected Customer
	}{
		{
			"no password",
			`{"type":"customer","name":"Ron Swanson","email":"ron@swanson.com","password":false}`,
			Customer{
				Type:  "customer",
				Name:  "Ron Swanson",
				Email: "ron@swanson.com",
				Password: &CustomerPassword{
					IsSet:    false,
					Password: "",
				},
			},
		},
		{
			"with password",
			`{"type":"customer","name":"Ron Swanson","email":"ron@swanson.com","password":true}`,
			Customer{
				Type:  "customer",
				Name:  "Ron Swanson",
				Email: "ron@swanson.com",
				Password: &CustomerPassword{
					IsSet:    true,
					Password: "",
				},
			},
		},
		{
			"missing password",
			`{"type":"customer","name":"Ron Swanson","email":"ron@swanson.com"}`,
			Customer{
				Type:  "customer",
				Name:  "Ron Swanson",
				Email: "ron@swanson.com",
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			var customer Customer
			err := json.Unmarshal([]byte(test.json), &customer)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(customer, test.expected) {
				t.Errorf(
					"\nexpected: %#v\ngot:       %#v\n",
					test.expected,
					customer,
				)
			}
		})
	}
}
