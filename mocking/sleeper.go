package mocking

import "time"

type Sleeper interface {
	Sleep()
}

type MockSleeper struct {
	Calls int
}

func (m *MockSleeper) Sleep() {
	m.Calls++
}

type ConfigurableSleeper struct {
	Duration time.Duration
	SleepFunc    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleepFunc(c.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
