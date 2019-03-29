package gomo

import (
	"time"
)

// APIExecution records the execution time of the call
type APIExecution struct {
	StartTime time.Time
	EndTime   time.Time
}

// Start the timer
func (e *APIExecution) Start() {
	e.StartTime = time.Now()
}

// End the timer
func (e *APIExecution) End() {
	e.EndTime = time.Now()
}

// Elapsed returns the duration of the timer
func (e APIExecution) Elapsed() time.Duration {
	return e.EndTime.Sub(e.StartTime)
}
