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

// Data sets a target for a responses data resource
func Data(target interface{}) RequestResource {
	return func(w *wrapper) {
		w.addResource("data", target)
	}
}

// Included sets a target for a responses included resource
func Included(target interface{}) RequestResource {
	return func(w *wrapper) {
		w.addResource("included", target)
	}
}

// Meta sets the a for a responses meta resource
func Meta(target interface{}) RequestResource {
	return func(w *wrapper) {
		w.addResource("meta", target)
	}
}

// Links sets a target for a responses links resource
func Links(target interface{}) RequestResource {
	return func(w *wrapper) {
		w.addResource("links", target)
	}
}

// Errors sets a target for the responses errors
func Errors(target *[]APIError) RequestResource {
	return func(w *wrapper) {
		w.addResource("errors", target)
	}
}

// Paginate sets the page to select bases on the offset and limit.
// See https://docs.moltin.com/api/basics/pagination
func Paginate(offset, limit int) RequestResource {
	return func(w *wrapper) {
		w.Query.Add("page[limit]", strconv.Itoa(limit))
		w.Query.Add("page[offset]", strconv.Itoa(offset))
	}
}

// Filter adds a filter to a query, prepending to any existing
// filters. See https://docs.moltin.com/api/basics/filtering
func Filter(filter string) RequestResource {
	return queryParameter("filter", filter, ":")
}

// Sort sorts the results.
// See https://docs.moltin.com/api/basics/sorting
func Sort(by string) RequestResource {
	return func(w *wrapper) {
		w.Query.Set("sort", by)
	}
}

// Include adds a resource to be included in the request.
// See https://docs.moltin.com/api/basics/includes
func Include(include string) RequestResource {
	return queryParameter("include", include, ",")
}

// ExecutionTime returns a pointer to the ExecutionTime for the request
func ExecutionTime(e **APIExecution) RequestResource {
	return func(w *wrapper) {
		*e = &w.ExecutionTime
	}
}

func queryParameter(parameter, value, separator string) RequestResource {
	return func(w *wrapper) {
		existing := w.Query.Get(parameter)
		if existing != "" {
			value += separator + existing
		}
		w.Query.Set(parameter, value)
	}
}
