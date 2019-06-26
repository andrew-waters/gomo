package gomo_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/moltin/gomo"
	"github.com/moltin/gomo/core"
)

func TestNextPage(t *testing.T) {
	for _, test := range []struct {
		name           string
		meta           core.Meta
		expectedOffset int
		expectedLimit  int
	}{
		{
			"more than a page left",
			core.Meta{
				Page: core.MetaPage{
					Offset: 0,
					Limit:  100,
					Total:  3,
				},
			},
			100,
			100,
		},
		{
			"exactly one page left",
			core.Meta{
				Page: core.MetaPage{
					Offset: 100,
					Limit:  100,
					Total:  3,
				},
			},
			200,
			100,
		},
		{
			"less than one page left",
			core.Meta{
				Page: core.MetaPage{
					Offset: 250,
					Limit:  100,
					Total:  3,
				},
			},
			350,
			100,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			offset, limit := gomo.NextPage(test.meta)
			if offset != test.expectedOffset {
				t.Errorf("bad offset: expected %d, got %d", test.expectedOffset, offset)
			}
			if limit != test.expectedLimit {
				t.Errorf("bad limit: expected %d, got %d", test.expectedLimit, limit)
			}
		})
	}
}

func TestMorePages(t *testing.T) {
	for _, test := range []struct {
		name            string
		meta            core.Meta
		expectMorePages bool
	}{
		{
			"exactly one page left",
			core.Meta{
				Page: core.MetaPage{
					Current: 1,
					Total:   2,
				},
			},
			true,
		},
		{
			"no pages left",
			core.Meta{
				Page: core.MetaPage{
					Current: 2,
					Total:   2,
				},
			},
			false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			more := gomo.MorePages(test.meta)
			if more != test.expectMorePages {
				t.Errorf("expected %v, got %v", test.expectMorePages, more)
			}
		})
	}
}

func TestIterate(t *testing.T) {
	count := 0
	handler := func(w http.ResponseWriter, r *http.Request) {
		count++
		offset, _ := strconv.Atoi(r.URL.Query().Get("page[offset]"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("page[limit]"))
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"meta":{"page":{"offset":%d,"limit":%d,"current":%d,"total":10}}}`, offset, limit, offset/limit+1)))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	client := gomo.NewClient(
		gomo.ClientCredentials("id", "secret"),
		gomo.Endpoint(server.URL),
	)
	err := gomo.Iterate(10, func(paginate gomo.RequestResource, meta *core.Meta) error {
		_, err := client.Get("", paginate)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 10 {
		t.Errorf("wrong number of pages: %d", count)
	}
}
