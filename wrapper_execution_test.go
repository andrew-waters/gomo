package gomo

import (
	"testing"
	"time"
)

func TestWrapperExecutionTimes(t *testing.T) {
	e := APIExecution{}
	c := make(chan time.Time, 1)
	e.clock = func() time.Time {
		return <-c
	}

	start := time.Date(1985, time.October, 26, 01, 21, 0, 0, time.UTC)
	end := time.Date(1985, time.October, 26, 01, 22, 0, 0, time.UTC)

	c <- start
	e.Start()
	c <- end
	e.End()

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
