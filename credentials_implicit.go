package gomo

import "net/url"

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
