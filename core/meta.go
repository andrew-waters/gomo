package core

// Meta is returned from many queries that support pagination
type Meta struct {
	Results MetaResults `json:"results"`
	Page    MetaPage    `json:"page"`
}

// MetaResults contains the total number of results
type MetaResults struct {
	Total int `json:"total"`
	All   int `json:"all,omitempty"`
}

// MetaPage contains pagination information
type MetaPage struct {
	Limit   int `json:"limit"`
	Offset  int `json:"offset"`
	Current int `json:"current"`
	Total   int `json:"total"`
}
