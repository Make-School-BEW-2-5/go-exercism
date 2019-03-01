package clock

import "fmt"

type Clock int

func New(hour int, minute int) Clock {
	return Clock(0).Add(hour*60 + minute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

func (c Clock) Add(min int) Clock {
	day := 24 * 60
	return Clock((int(c) + day + min%day) % day)
}

func (c Clock) Subtract(min int) Clock {
	day := 24 * 60
	return Clock((int(c) + day - min%day) % day)
}
