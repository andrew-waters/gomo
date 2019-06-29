package gomo

import (
	"encoding/json"
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
	Response      map[string]json.RawMessage
	Resources     []resourceTarget
	Errors        []APIError
}

type resourceTarget struct {
	Section string
	Target  interface{}
}

func (w *wrapper) addResource(section string, target interface{}) {
	w.Resources = append(
		w.Resources,
		resourceTarget{section, target},
	)
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
