package gomo

import (
	"net/url"
	"testing"
)

func TestBody(t *testing.T) {
	var w wrapper
	Body("foobar")(&w)
	if s, ok := w.body.(string); !ok || s != "foobar" {
		t.Fatal("failed to set data")
	}
}

func TestResouceSetters(t *testing.T) {
	for _, test := range []struct {
		section         string
		requestResource RequestResource
	}{
		{
			"data",
			Data("foobar"),
		},
		{
			"links",
			Links("foobar"),
		},
		{
			"included",
			Included("foobar"),
		},
		{
			"meta",
			Meta("foobar"),
		},
	} {
		t.Run(test.section, func(t *testing.T) {
			var w wrapper
			test.requestResource(&w)
			if len(w.resources) == 0 {
				t.Fatal("failed to add resource")
			}
			rt := w.resources[0]
			if rt.section != test.section {
				t.Errorf("wrong section: %s", rt.section)
			}
			if s, ok := rt.target.(string); !ok || s != "foobar" {
				t.Fatal("failed to set target")
			}
		})
	}
}

func TestPaginate(t *testing.T) {
	var w wrapper
	w.query = make(url.Values)
	Paginate(200, 100)(&w)
	if w.query.Get("page[offset]") != "200" {
		t.Error("failed to set offset")
	}
	if w.query.Get("page[limit]") != "100" {
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
			w.query = make(url.Values)
			for _, filter := range test.filters {
				filter(&w)
			}
			filter := w.query.Get("filter")
			if filter != test.expectedFilter {
				t.Errorf("expected: %s, got %s", test.expectedFilter, filter)
			}
		})
	}
}

func TestInclude(t *testing.T) {
	for _, test := range []struct {
		name            string
		includes        []RequestResource
		expectedInclude string
	}{
		{
			"single",
			[]RequestResource{
				Include("products"),
			},
			"products",
		},
		{
			"multiple",
			[]RequestResource{
				Include("products"),
				Include("categories"),
			},
			"categories,products",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			var w wrapper
			w.query = make(url.Values)
			for _, include := range test.includes {
				include(&w)
			}
			include := w.query.Get("include")
			if include != test.expectedInclude {
				t.Errorf("expected: %s, got %s", test.expectedInclude, include)
			}
		})
	}
}

func TestSort(t *testing.T) {
	var w wrapper
	w.query = make(url.Values)
	Sort("foo")(&w)
	sort := w.query.Get("sort")
	if sort != "foo" {
		t.Error("failed to sort")
	}
}
