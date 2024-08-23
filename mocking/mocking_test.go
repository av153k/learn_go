package mocking

import (
	"bytes"
	"testing"
)

func TestMocking(t *testing.T) {
	buffer := &bytes.Buffer{}
	mockSleeper := &MockSleeper{}
	Countdown(buffer, mockSleeper)

	got := buffer.String()
	want := "3\n2\n1\nGO!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if mockSleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", mockSleeper.Calls)
	}

}
