package gomo

import (
	"testing"
	"time"

	"github.com/bouk/monkey"
)

func TestWrapperExecutionTimes(t *testing.T) {
	start := time.Date(1985, time.October, 26, 01, 21, 0, 0, time.UTC)
	patch := monkey.Patch(time.Now, func() time.Time {
		return start
	})

	e := APIExecution{}
	e.Start()
	patch.Unpatch()

	end := time.Date(1985, time.October, 26, 01, 22, 0, 0, time.UTC)
	patch = monkey.Patch(time.Now, func() time.Time {
		return end
	})
	e.End()
	patch.Unpatch()

	if e.StartTime != start {
		t.Errorf("Start time did return correct start time")
	}

	if e.EndTime != end {
		t.Errorf("End time did return correct end time")
	}

	duration := end.Sub(start)
	if e.Elapsed() != duration {
		t.Errorf("Elapsed did return correct elapsed time")
	}
}
