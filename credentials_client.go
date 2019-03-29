package gomo

import "net/url"

// NewClientCredentials creates and returns a clientCredentials struct from the passed in values
func NewClientCredentials(clientID string, clientSecret string) clientCredentials {
	return clientCredentials{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

// clientCredentials contains credentials for server side processing
type clientCredentials struct {
	clientID     string
	clientSecret string
}

func (cc clientCredentials) grantType() string {
	return clientCredentialsGrantType
}

func (cc clientCredentials) authFormValues() url.Values {
	return url.Values{
		"grant_type":    {cc.grantType()},
		"client_id":     {cc.clientID},
		"client_secret": {cc.clientSecret},
	}
}
