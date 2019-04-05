package gomo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var errUnableToAuthenticate = errors.New("Unable to authenticate")

// AuthResponse contains the response from the auth call
type authResponse struct {
	Expires     int    `json:"expires"`
	ExpiresIn   int    `json:"expires_in"`
	Identifier  string `json:"identifier"`
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

func (c Client) authURL() string {
	return fmt.Sprintf("%s/oauth/access_token", c.Endpoint)
}

// Authenticate makes a call to get the access token for the client's credentials
func (c *Client) Authenticate() error {
	var err error

	r, err := http.PostForm(c.authURL(), c.credentials.authFormValues())
	if err != nil {
		return err
	}

	if r.StatusCode != 200 {
		return errUnableToAuthenticate
	}

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	ar := authResponse{}
	err = json.Unmarshal(b, &ar)
	if err != nil {
		return err
	}

	c.AccessToken = ar.AccessToken

	return nil
}
