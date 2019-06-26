package gomo

import (
	"net/url"
	"testing"
)

func TestBody(t *testing.T) {
	var w wrapper
	Body("foobar")(&w)
	if s, ok := w.Body.(string); !ok || s != "foobar" {
		t.Fatal("failed to set data")
	}
}

func TestData(t *testing.T) {
	var w wrapper
	Data("foobar")(&w)
	if s, ok := w.Response.Data.(string); !ok || s != "foobar" {
		t.Fatal("failed to set data")
	}
}

func TestLinks(t *testing.T) {
	var w wrapper
	Links("foobar")(&w)
	if s, ok := w.Response.Links.(string); !ok || s != "foobar" {
		t.Fatal("failed to set links")
	}
}

func TestIncluded(t *testing.T) {
	var w wrapper
	Included("foobar")(&w)
	if s, ok := w.Response.Included.(string); !ok || s != "foobar" {
		t.Fatal("failed to set included")
	}
}

func TestMeta(t *testing.T) {
	var w wrapper
	Meta("foobar")(&w)
	if s, ok := w.Response.Meta.(string); !ok || s != "foobar" {
		t.Fatal("failed to set meta")
	}
}

func TestPaginate(t *testing.T) {
	var w wrapper
	w.Query = make(url.Values)
	Paginate(200, 100)(&w)
	if w.Query.Get("page[offset]") != "200" {
		t.Error("failed to set offset")
	}
	if w.Query.Get("page[limit]") != "100" {
		t.Error("failed to set limit")
	}
}

func TestFilter(t *testing.T) {
	for _, test := range []struct {
		name           string
		filters        []RequestResource
		expectedFilter string
	}{
		{
			"single",
			[]RequestResource{
				Filter("eq(status,live)"),
			},
			"eq(status,live)",
		},
		{
			"multiple",
			[]RequestResource{
				Filter("eq(status,live)"),
				Filter("like(name,Deck*)"),
			},
			"like(name,Deck*):eq(status,live)",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			var w wrapper
			w.Query = make(url.Values)
			for _, filter := range test.filters {
				filter(&w)
			}
			filter := w.Query.Get("filter")
			if filter != test.expectedFilter {
				t.Errorf("expected: %s, got %s", test.expectedFilter, filter)
			}
		})
	}
}
