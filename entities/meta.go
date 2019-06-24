package entities

// Meta is returned from many queries that support pagination
type Meta struct {
	Results struct {
		Total int `json:"total"`
		All   int `json:"all"`
	} `json:"results"`
	Page struct {
		Limit   int `json:"limit"`
		Offset  int `json:"offset"`
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"page"`
}
