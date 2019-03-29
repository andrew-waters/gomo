package gomo

import (
	"net/http"
	"strings"
)

// wrapper is a wrapper around a request and response to the API
type wrapper struct {
	Method        string
	Endpoint      string
	StatusCode    int
	ExecutionTime APIExecution
	Body          interface{}
	Request       *http.Request
	Response      response
}

// APIResponse contains the response data to the call
type response struct {
	Data     interface{} `json:"data"`
	Meta     interface{} `json:"meta"`
	Included interface{} `json:"included"`
	Errors   []APIError  `json:"errors"`
}

// newAPIWrapper creates a new wrapper for this call
func newWrapper(method string, endpoint string, resources ...interface{}) wrapper {

	var targetResource interface{}
	var targetIncludes interface{}

	for k, resource := range resources {
		switch k {
		case 0:
			targetResource = resource
		case 1:
			targetIncludes = resource
		}
	}

	// set the resource type if the entity has the SetType method
	if targetResource, ok := targetResource.(interface{ SetType() }); ok {
		targetResource.SetType()
	}

	return wrapper{
		Method:   strings.ToUpper(method),
		Endpoint: endpoint,
		Body:     targetResource,
		Response: response{
			Data:     targetResource,
			Included: targetIncludes,
		},
	}
}
