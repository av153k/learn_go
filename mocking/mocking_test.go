package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestMocking(t *testing.T) {
	t.Run("Test simple read write with mock sleeper", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()
		want := "3\n2\n1\nGO!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("Check if it sleeps before each print", func(t *testing.T) {
		spyCountdown := &SpyCountdownOperations{}

		Countdown(spyCountdown, spyCountdown)

		want := []string{
			write, sleep, write, sleep, write, sleep, write,
		}

		if !reflect.DeepEqual(want, spyCountdown.Calls) {
			t.Errorf("wanted calls %v got %v", want, spyCountdown.Calls)
		}
	})

}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}

}
