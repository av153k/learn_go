package main

import (
	"os"
	"time"

	"av153k.github.io/learn-go/mocking"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second, SleepFunc: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
