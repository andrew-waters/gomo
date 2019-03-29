package gomo

import "net/url"

const (
	clientCredentialsGrantType = "client_credentials"
	implicitGrantType          = "implicit"
)

// Credentials provides an interface for different credential types
type credentials interface {
	grantType() string
	authFormValues() url.Values
}
