package mocking

type Sleeper interface {
	Sleep()
}

type MockSleeper struct {
	Calls int
}

func (m *MockSleeper) Sleep() {
	m.Calls++
}
