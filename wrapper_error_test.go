package gomo

import "testing"

var errorTests = []struct {
	status   int
	title    string
	detail   string
	expected string
}{
	{200, "problem", "info", "{ Status: `200`, Title: `problem`, Detail: `info` }"},
	{200, "problem", "", "{ Status: `200`, Title: `problem`, Detail: `` }"},
	{200, "", "", "{ Status: `200`, Title: ``, Detail: `` }"},
	{300, "", "", "{ Status: `300`, Title: ``, Detail: `` }"},
}

func TestAPIErrorString(t *testing.T) {

	for _, test := range errorTests {

		e := APIError{
			Status: test.status,
			Title:  test.title,
			Detail: test.detail,
		}

		if e.String() != test.expected {
			t.Errorf("APIError to string returned `%s` (expected `%s`", e.String(), test.expected)
		}
	}
}
