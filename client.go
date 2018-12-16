package gomo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

const (
	// DefaultAPIVersion means less config from the user
	DefaultAPIVersion = "v2"
	// DefaulEndpoint means less config from the user
	DefaulEndpoint = "https://api.moltin.com"
	// DefaultUserAgent is sent as a header in the API call
	DefaultUserAgent = "go-client"
)

// Client is the main client struct
type Client struct {
	credentials Credentials
	APIVersion  string
	Endpoint    string
	AccessToken string
	Debug       bool
	Logs        []interface{}
	httpClient  *http.Client
}

// NewClient creates a new client for you to make requests with
func NewClient(c Credentials) (Client, error) {
	var err error

	client := Client{
		credentials: c,
		APIVersion:  DefaultAPIVersion,
		Endpoint:    DefaulEndpoint,
		Debug:       false,
		httpClient:  &http.Client{},
	}
	err = client.authenticate()

	return client, err
}

// GrantType returns the string value of the current crednetials grant type
func (c Client) GrantType() string {
	return c.credentials.grantType()
}

func (c *Client) EnableDebug() {
	c.Debug = true
}

func (c *Client) Log(msgs ...interface{}) {
	for _, msg := range msgs {
		c.Logs = append(c.Logs, msg)
		if c.Debug {
			spew.Dump(msg)
		}
	}
}

func (c Client) url(endpoint string) string {
	return fmt.Sprintf("%s/%s/%s", c.Endpoint, c.APIVersion, endpoint)
}

func (c Client) authURL() string {
	return fmt.Sprintf("%s/oauth/access_token", c.Endpoint)
}

func (c Client) buildRequest(method string, endpoint string, body []byte) (*http.Request, error) {
	var err error

	req, err := http.NewRequest(method, c.url(endpoint), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.AccessToken)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", DefaultUserAgent)

	return req, nil
}

func (c *Client) do(wrapper *APIWrapper) error {
	var err error

	var body []byte
	if wrapper.Body != nil {

		rb := APIResponse{
			Data: wrapper.Body,
		}
		rbj, err := json.Marshal(rb)
		if err != nil {
			return err
		}

		body = rbj
	}

	wrapper.Request, err = c.buildRequest(wrapper.Method, wrapper.Endpoint, body)
	if err != nil {
		return err
	}

	wrapper.ExecutionTime.Start()
	r, err := c.httpClient.Do(wrapper.Request)
	wrapper.ExecutionTime.End()

	if err != nil {
		return err
	}

	wrapper.StatusCode = r.StatusCode

	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &wrapper.Response)
	c.Log(wrapper)

	if len(wrapper.Response.Errors) > 0 {
		return errors.New(wrapper.Response.Errors[0].Detail)
	}

	return nil
}

// Post makes a POST request to the API
func (c *Client) Post(endpoint string, resource interface{}) (APIWrapper, error) {
	var err error

	wrapper := NewAPIWrapper("post", endpoint, resource)
	err = c.do(&wrapper)

	return wrapper, err
}

// Get makes a GET request to the API
func (c *Client) Get(endpoint string, resource ...interface{}) (APIWrapper, error) {
	var err error

	wrapper := NewAPIWrapper("get", endpoint, resource...)
	err = c.do(&wrapper)

	return wrapper, err
}

// Delete makes a DELETE request to the API
func (c *Client) Delete(endpoint string) (APIWrapper, error) {
	var err error

	wrapper := NewAPIWrapper("delete", endpoint)
	err = c.do(&wrapper)

	return wrapper, err
}

// Put makes a PUT request to the API
func (c *Client) Put(endpoint string, resource interface{}) (APIWrapper, error) {
	var err error

	wrapper := NewAPIWrapper("put", endpoint, resource)
	err = c.do(&wrapper)

	return wrapper, err
}

// authenticate makes a call to get the access token
func (c *Client) authenticate() error {
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
