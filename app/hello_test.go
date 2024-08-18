package app

import "testing"

func TestHelloWorld(t *testing.T) {
	got := hello("World")
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
