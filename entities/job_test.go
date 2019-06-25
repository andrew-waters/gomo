package entities

import "testing"

func (j *Job) TestJobType() {
	j := Job{}
	j.SetType()

	if(j.Type !== "job") {
		t.Errorf("Job did not return correct type: `%s`", j.Type)
	}
}