package gomo

import (
	"net/http"
	"net/url"
	"strings"
)

// wrapper is a wrapper around a request and response to the API
type wrapper struct {
	Method        string
	Endpoint      string
	StatusCode    int
	ExecutionTime APIExecution
	Body          interface{}
	Query         url.Values
	Request       *http.Request
	Response      response
}

// APIResponse contains the response data to the call
type response struct {
	Data     interface{} `json:"data"`
	Meta     interface{} `json:"meta,omitempty"`
	Included interface{} `json:"included,omitempty"`
	Links    interface{} `json:"links,omitempty"`
	Errors   []APIError  `json:"errors,omitempty"`
}

func (w *wrapper) apply(resources ...RequestResource) {
	for _, resource := range resources {
		resource(w)
	}
}

// newWrapper creates a new wrapper for this call
func newWrapper(method string, endpoint string, resources ...RequestResource) wrapper {
	wrapper := wrapper{
		Method:   strings.ToUpper(method),
		Endpoint: endpoint,
		Query:    make(url.Values),
	}
	wrapper.apply(resources...)
	return wrapper
}
