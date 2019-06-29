package gomo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// RequestResource are functions that provide a request with the
// resources it requires. This includes Body() which is the body of
// the request, Data() which sets the target struct for the returned
// data, etc
type RequestResource func(*wrapper)

// Post makes a POST request to the API
func (c *Client) Post(endpoint string, resource ...RequestResource) error {
	wrapper := newWrapper("post", endpoint, resource...)
	return c.do(&wrapper)
}

// Get makes a GET request to the API
func (c *Client) Get(endpoint string, resource ...RequestResource) error {
	wrapper := newWrapper("get", endpoint, resource...)
	return c.do(&wrapper)
}

// Delete makes a DELETE request to the API
func (c *Client) Delete(endpoint string, resource ...RequestResource) error {
	wrapper := newWrapper("delete", endpoint, resource...)
	return c.do(&wrapper)
}

// Put makes a PUT request to the API
func (c *Client) Put(endpoint string, resource ...RequestResource) error {
	wrapper := newWrapper("put", endpoint, resource...)
	return c.do(&wrapper)
}

func (c Client) url(endpoint string) string {
	return fmt.Sprintf("%s/%s/%s", c.Endpoint, c.APIVersion, endpoint)
}

func (c *Client) do(wrapper *wrapper) error {
	var err error

	resp, err := c.makeRequest(wrapper)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}

	return c.parseResponse(resp, wrapper)

}

func (c Client) buildRequest(method string, endpoint string, query url.Values, body []byte) (*http.Request, error) {
	var err error

	req, err := http.NewRequest(method, c.url(endpoint), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer: %s", c.AccessToken))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", defaultUserAgent)

	req.URL.RawQuery = query.Encode()

	return req, nil
}

func (c *Client) makeRequest(wrapper *wrapper) (*http.Response, error) {
	var body []byte

	if wrapper.body != nil {
		rb := struct {
			Data interface{} `json:"data"`
		}{wrapper.body}
		rbj, err := json.Marshal(rb)
		if err != nil {
			return nil, err
		}

		body = rbj
	}

	req, err := c.buildRequest(wrapper.method, wrapper.endpoint, wrapper.query, body)
	wrapper.request = req
	if err != nil {
		return nil, err
	}

	wrapper.executionTime.Start()
	resp, err := c.httpClient.Do(wrapper.request)
	wrapper.executionTime.End()

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c Client) parseResponse(resp *http.Response, wrapper *wrapper) error {

	wrapper.statusCode = resp.StatusCode

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read request body: %v", err)
	}

	err = json.Unmarshal(b, &wrapper.response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %v", err)
	}
	if errsJSON, ok := wrapper.response["errors"]; ok {
		err = json.Unmarshal(errsJSON, &wrapper.errors)
		if err != nil {
			return fmt.Errorf("failed to unmarshal response errors: %v", err)
		}
	}

	for _, r := range wrapper.resources {
		sectionJSON, ok := wrapper.response[r.section]
		if !ok {
			continue
		}
		err := json.Unmarshal(sectionJSON, r.target)
		if err != nil {
			return fmt.Errorf("failed to unmarshal response %s: %v", r.section, err)
		}
	}

	if len(wrapper.errors) > 0 {
		e := wrapper.errors[0]
		return fmt.Errorf("%s: %s", e.Title, e.Detail)
	}

	return nil
}
