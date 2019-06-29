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
		w.addResource("data", target)
	}
}

// Included sets the target for a responses included resource
func Included(target interface{}) RequestResource {
	return func(w *wrapper) {
		w.addResource("included", target)
	}
}

// Meta sets the target for a responses meta resource
func Meta(target interface{}) RequestResource {
	return func(w *wrapper) {
		w.addResource("meta", target)
	}
}

// Links sets the target for a responses links resource
func Links(target interface{}) RequestResource {
	return func(w *wrapper) {
		w.addResource("links", target)
	}
}

// Paginate sets the page to select bases on the offset and limit
func Paginate(offset, limit int) RequestResource {
	return func(w *wrapper) {
		w.Query.Add("page[limit]", strconv.Itoa(limit))
		w.Query.Add("page[offset]", strconv.Itoa(offset))
	}
}

// Filter adds a filter to a query, prepending to any existing
// filters
func Filter(filter string) RequestResource {
	return func(w *wrapper) {
		existingFilter := w.Query.Get("filter")
		if existingFilter != "" {
			filter += ":" + existingFilter
		}
		w.Query.Set("filter", filter)
	}
}
