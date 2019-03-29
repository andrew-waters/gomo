package gomo

import "net/url"

const (
	// CCGrantType is the string value for a client credentials grant type
	CCGrantType = "client_credentials"
	// ICGrantType is the string value for an implicit credentials grant type
	ICGrantType = "implicit"
)

// Credentials provides an interface for different credential types
type credentials interface {
	grantType() string
	authFormValues() url.Values
}
