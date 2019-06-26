package gomo

import "strconv"

// Body sets the body for a Post() or Put() request.
func Body(target interface{}) RequestResource {
	return func(w *wrapper) {
		w.Body = target

		// set the resource type if the entity has the SetType method
		if resource, ok := w.Body.(interface{ SetType() }); ok {
			resource.SetType()
		}
	}
}

// Data sets the target for a responses data resource
func Data(target interface{}) RequestResource {
	return func(w *wrapper) {
		w.Response.Data = target
	}
}

// Included sets the target for a responses included resource
func Included(target interface{}) RequestResource {
	return func(w *wrapper) {
		w.Response.Included = target
	}
}

// Meta sets the target for a responses meta resource
func Meta(target interface{}) RequestResource {
	return func(w *wrapper) {
		w.Response.Meta = target
	}
}

// Links sets the target for a responses links resource
func Links(target interface{}) RequestResource {
	return func(w *wrapper) {
		w.Response.Links = target
	}
}

// Paginate sets the page to select bases on the offset and limit
func Paginate(offset, limit int) RequestResource {
	return func(w *wrapper) {
		w.Query.Add("page[limit]", strconv.Itoa(limit))
		w.Query.Add("page[offset]", strconv.Itoa(offset))
	}
}
