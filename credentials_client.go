package gomo

import "net/url"

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
