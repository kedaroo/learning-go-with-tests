package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		recipient := "Kedar"
		language := ""
		got := Hello(recipient, language)
		want := fmt.Sprintf("Hello, %s", recipient)
		assertCorrectMsg(t, got, want)
	})

	t.Run("saying \"Hello, world\" when an empty strin is supplied", func(t *testing.T) {
		recipient := ""
		language := ""
		got := Hello(recipient, language)
		want := "Hello, world"
		assertCorrectMsg(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		recipient := ""
		language := "Spanish"
		got := Hello(recipient, language)
		want := "Hola, world"
		assertCorrectMsg(t, got, want)
	})
}

func assertCorrectMsg(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}