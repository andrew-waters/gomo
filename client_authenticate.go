package gomo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
		return errors.New("Unable to authenticate")
	}

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(b, &ar)
	if err != nil {
		return err
	}

	c.AccessToken = ar.AccessToken

	return nil
}
