package entities

type Job struct {
	ID string `json:"id",omitempty`
	Type string `json:"type"`
	JobType string `json:"job_type"`
	Link struct {
		Href string `json:"href"`
	} `json:"link",omitempty`
	Status string `json:"",omitempty`
	Timestamps Timestamps `json:"",omitempty`
}

func (j *Job) SetType() {
	j.Type = jobType
}