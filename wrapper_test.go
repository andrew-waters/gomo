package gomo

import "testing"

func TestBody(t *testing.T) {
	var w wrapper
	Body("foobar")(&w)
	if s, ok := w.Body.(string); !ok || s != "foobar" {
		t.Fatal("failed to set data")
	}
}

func TestData(t *testing.T) {
	var w wrapper
	Data("foobar")(&w)
	if s, ok := w.Response.Data.(string); !ok || s != "foobar" {
		t.Fatal("failed to set data")
	}
}

func TestLinks(t *testing.T) {
	var w wrapper
	Links("foobar")(&w)
	if s, ok := w.Response.Links.(string); !ok || s != "foobar" {
		t.Fatal("failed to set links")
	}
}

func TestIncluded(t *testing.T) {
	var w wrapper
	Included("foobar")(&w)
	if s, ok := w.Response.Included.(string); !ok || s != "foobar" {
		t.Fatal("failed to set included")
	}
}

func TestMeta(t *testing.T) {
	var w wrapper
	Meta("foobar")(&w)
	if s, ok := w.Response.Meta.(string); !ok || s != "foobar" {
		t.Fatal("failed to set meta")
	}
}

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
