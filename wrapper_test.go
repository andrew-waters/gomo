package gomo

import "testing"

func TestNewWrapper(t *testing.T) {
	called := false
	w := newWrapper(
		"method",
		"endpoint",
		func(wp *wrapper) {
			called = true
		},
	)
	if w.Method != "METHOD" {
		t.Error("failed to set method")
	}
	if w.Endpoint != "endpoint" {
		t.Error("failed to set endpoint")
	}
	if called == false {
		t.Error("failed to call option function")
	}
}
