package mocking

import (
	"fmt"
	"io"
)

const finalWord = "GO!"
const tick = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := tick; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
