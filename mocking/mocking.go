package mocking

import (
	"fmt"
	"io"
)

const (
	sleep = "SLEEP"
	write = "WRITE"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const finalWord = "GO!"
const tick = 3

func Countdown(out io.Writer, sleeper Sleeper) {

	for i := tick; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
