package entities

import "testing"

func TestJobType(t *testing.T) {
	j := Job{}
	j.SetType()

	if(j.Type !== "job") {
		t.Errorf("Job did not return correct type: `%s`", j.Type)
	}
}