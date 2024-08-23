package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Abhishek")

	got := buffer.String()
	want := "Hello, Abhishek!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}