package gomo

import (
	"time"
)

// APIExecution records the execution time of the call
type APIExecution struct {
	StartTime time.Time
	EndTime   time.Time
	clock     func() time.Time
}

// Start the timer
func (e *APIExecution) Start() {
	e.StartTime = e.now()
}

// End the timer
func (e *APIExecution) End() {
	e.EndTime = e.now()
}

// Elapsed returns the duration of the timer
func (e APIExecution) Elapsed() time.Duration {
	return e.EndTime.Sub(e.StartTime)
}

func (e APIExecution) now() time.Time {
	if e.clock == nil {
		return time.Now()
	}
	return e.clock()
}
