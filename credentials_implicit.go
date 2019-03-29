package gomo

import "net/url"

// NewImplicitCredentials creates and returns an implicitCredentials struct from the passed in client ID
func NewImplicitCredentials(clientID string) implicitCredentials {
	return implicitCredentials{
		clientID: clientID,
	}
}

// ImplicitCredentials gives you read access (and some write access) and is for untrusted environments
type implicitCredentials struct {
	clientID string
}

func (ic implicitCredentials) grantType() string {
	return implicitGrantType
}

func (ic implicitCredentials) authFormValues() url.Values {
	return url.Values{
		"grant_type": {ic.grantType()},
		"client_id":  {ic.clientID},
	}
}
