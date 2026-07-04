package main

import (
	"testing"
	// "hello"
)

func TestHello(t *testing.T) {
	got := hello("maneesh")
	want := "hello maneesh"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
