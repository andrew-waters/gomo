package gomo

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// wrapper is a wrapper around a request and response to the API
type wrapper struct {
	method        string
	endpoint      string
	statusCode    int
	executionTime APIExecution
	body          interface{}
	query         url.Values
	request       *http.Request
	response      map[string]json.RawMessage
	resources     []resourceTarget
	errors        []APIError
}

type resourceTarget struct {
	section string
	target  interface{}
}

func (w *wrapper) addResource(section string, target interface{}) {
	w.resources = append(
		w.resources,
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
		method:   strings.ToUpper(method),
		endpoint: endpoint,
		query:    make(url.Values),
	}
	wrapper.apply(resources...)
	return wrapper
}
