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
	Links    interface{} `json:"links"`
	Errors   []APIError  `json:"errors"`
}

// Data sets the target for a responses data resource
func Data(target interface{}) func(*wrapper) {
	return func(w *wrapper) {
		w.Body = target
		w.Response.Data = target

		// set the resource type if the entity has the SetType method
		if resource, ok := target.(interface{ SetType() }); ok {
			resource.SetType()
		}
	}
}

// Included sets the target for a responses included resource
func Included(target interface{}) func(*wrapper) {
	return func(w *wrapper) {
		w.Response.Included = target
	}
}

// Meta sets the target for a responses meta resource
func Meta(target interface{}) func(*wrapper) {
	return func(w *wrapper) {
		w.Response.Meta = target
	}
}

// Links sets the target for a responses links resource
func Links(target interface{}) func(*wrapper) {
	return func(w *wrapper) {
		w.Response.Links = target
	}
}

// newWrapper creates a new wrapper for this call
func newWrapper(method string, endpoint string, resources ...func(*wrapper)) wrapper {
	wrapper := wrapper{
		Method:   strings.ToUpper(method),
		Endpoint: endpoint,
	}
	for _, resource := range resources {
		resource(&wrapper)
	}
	return wrapper
}
