package moltin

import "net/url"

const (
	// CCGrantType is the string value for a client credentials grant type
	CCGrantType = "client_credentials"
	// ICGrantType is the string value for an implicit credentials grant type
	ICGrantType = "implicit"
)

// Credentials provides an interface for different credential levels
type Credentials interface {
	grantType() string
	authFormValues() url.Values
}

// ClientCredentials contains credentials for server side processing
type ClientCredentials struct {
	clientID     string
	clientSecret string
}

func (cc ClientCredentials) grantType() string {
	return CCGrantType
}

func (cc ClientCredentials) authFormValues() url.Values {
	return url.Values{
		"grant_type":    {cc.grantType()},
		"client_id":     {cc.clientID},
		"client_secret": {cc.clientSecret},
	}
}

// NewClientCredentials creates and returns a clientCredentials struct from the passed in values
func NewClientCredentials(clientID string, clientSecret string) ClientCredentials {
	return ClientCredentials{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

// ImplicitCredentials gives you read access (and some write access) and is for untrusted environments
type ImplicitCredentials struct {
	clientID string
}

//
func (ic ImplicitCredentials) grantType() string {
	return ICGrantType
}

func (ic ImplicitCredentials) authFormValues() url.Values {
	return url.Values{
		"grant_type": {ic.grantType()},
		"client_id":  {ic.clientID},
	}
}

// NewImplicitCredentials creates and returns an implicitCredentials struct from the passed in values
func NewImplicitCredentials(clientID string) ImplicitCredentials {
	return ImplicitCredentials{
		clientID: clientID,
	}
}
