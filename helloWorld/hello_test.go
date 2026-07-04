package main

import (
	"testing"
	// "hello"
)

// func TestHello(t *testing.T) {
// 	got := hello("maneesh")
// 	want := "hello maneesh"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// func TestMore(t *testing.T) {
// 	t.Run("greeting pepole ", func(t *testing.T) {
// 		got := hello("manee")
// 		want := "hello manee"
// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// 	t.Run("greeting world ", func(t *testing.T) {
// 		got := hello("")
// 		want := "hello world"
// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// }

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Chris", "")
		want := "Hello Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := hello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("greeting in Hindi", func(t *testing.T) {
		got := hello("manee", "Hindi")
		want := "Namaste manee"
		assertCorrectMessage(t, got, want)
	})
	t.Run("greeting in Jaapanees", func(t *testing.T) {
		got := hello("manee", "Japanees")
		want := "Konnichiwa manee"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
