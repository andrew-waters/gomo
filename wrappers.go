package moltin

import (
	"net/http"
	"strings"
	"time"
)

// APIWrapper is a wrapper around a request and response to the API
type APIWrapper struct {
	Method        string
	Endpoint      string
	StatusCode    int
	ExecutionTime APIExecution
	Body          interface{}
	Request       *http.Request
	Response      APIResponse
}

// NewAPIWrapper creates a new wrapper for this call
func NewAPIWrapper(method string, endpoint string, resources ...interface{}) APIWrapper {

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

	return APIWrapper{
		Method:   strings.ToUpper(method),
		Endpoint: endpoint,
		Body:     targetResource,
		Response: APIResponse{
			Data:     targetResource,
			Included: targetIncludes,
		},
	}
}

// APIResponse contains the response data to the call
type APIResponse struct {
	Data     interface{} `json:"data"`
	Meta     interface{} `json:"meta"`
	Included interface{} `json:"included"`
	Errors   []APIError  `json:"errors"`
}

// APIExecution records the execution time of the call
type APIExecution struct {
	StartTime time.Time
	EndTime   time.Time
}

// Start the timer
func (e *APIExecution) Start() {
	e.StartTime = time.Now()
}

// End the timer
func (e *APIExecution) End() {
	e.EndTime = time.Now()
}

// Elapsed returns the duration of the timer
func (e APIExecution) Elapsed() time.Duration {
	return e.EndTime.Sub(e.StartTime)
}

// APIError is the format of an error returned byt the API
type APIError struct {
	Status int    `json:"status"`
	Detail string `json:"detail"`
	Title  string `json:"title"`
}

// APIRequestBodyWrapper is a wrapper for the reuqest body which elimates the need for `data` to be in the entity structs
type APIRequestBodyWrapper struct {
	Data interface{} `json:"data"`
}
