package gomo

import "github.com/moltin/gomo/core"

// NextPage advances the supplied Meta to the next page by returning a new
// offset and limit
func NextPage(meta core.Meta) (int, int) {
	offset := meta.Page.Offset + meta.Page.Limit
	return offset, meta.Page.Limit
}

// MorePages returns true is the meta suggests there are more results
// available
func MorePages(meta core.Meta) bool {
	return meta.Page.Current < meta.Page.Total
}

// Iterate calls the function with a RequestResource that should be
// passed to a Get() request. The function is called for each page
// of size limit. If the function returns an error at any point
// the iteration stops and the error is returned.
func Iterate(limit int, f func(RequestResource, *core.Meta) error) error {
	meta := core.Meta{
		Page: core.MetaPage{
			Offset:  0,
			Limit:   limit,
			Current: 0,
			Total:   1,
		},
	}
	offset := 0
	for MorePages(meta) {
		err := f(func(w *wrapper) {
			Paginate(offset, limit)(w)
			Meta(&meta)(w)
		}, &meta)
		if err != nil {
			return err
		}
		offset, limit = NextPage(meta)
	}
	return nil
}
