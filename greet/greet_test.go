package greet

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jack", "English")
		want := "Hello, Jack"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Bob", "Spanish")
		want := "Hola, Bob"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Bob", "French")
		want := "Bonjour, Bob"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in English", func(t *testing.T) {
		got := Hello("Bob", "English")
		want := "Hello, Bob"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Arabic", func(t *testing.T) {
		got := Hello("Adam", "Arabic")
		want := "Assalamu Alaikum, Adam"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
