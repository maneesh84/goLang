package main

import (
	"testing"
	// "hello"
)

func TestHello(t *testing.T) {
	got := hello()
	want := "hello world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
