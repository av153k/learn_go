package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := hello("Abhishek", "")
		want := "Hello, Abhishek!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := hello("Abhishek", "Spanish")
		want := "Hola, Abhishek!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := hello("Abhishek", "French")
		want := "Bonjour, Abhishek!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Japanese", func(t *testing.T) {
		got := hello("Abhishek", "Japanese")
		want := "Konnichiwa, Abhishek!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Hindi", func(t *testing.T) {
		got := hello("Abhishek", "Hindi")
		want := "Namaste, Abhishek!"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
