package clock

import (
	"fmt"
)

// Clock API:
//
// type Clock                      // define the clock type
// New(hour, minute int) Clock     // a "constructor"
// (Clock) String() string         // a "stringer"
// (Clock) Add(minutes int) Clock
// (Clock) Subtract(minutes int) Clock
//
// To satisfy the README requirement about clocks being equal, values of
// your Clock type need to work with the == operator. This means that if your
// New function returns a pointer rather than a value, your clocks will
// probably not work with ==.
//
// While the time.Time type in the standard library (https://golang.org/pkg/time/#Time)
// doesn't necessarily need to be used as a basis for your Clock type, it might
// help to look at how constructors there (Date and Now) return values rather
// than pointers. Note also how most time.Time methods have value receivers
// rather than pointer receivers.
//
// For some useful guidelines on when to use a value receiver or a pointer
// receiver see: https://github.com/golang/go/wiki/CodeReviewComments#receiver-type
type Clock struct {
	minutes int
}

// New initializes a new clock with hours and minutes
func New(h, m int) Clock {
	// return a pointer not a value
	c := (h*60 + m) % 1440
	if c < 0 {
		c += 1440
	}
	return Clock{minutes: c}
}

func (c Clock) String() string {
	// Reads military time
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

// Add is a receiver function that adds minutes to the clock
func (c Clock) Add(m int) Clock {
	return New(0, c.minutes+m)
}

// Subtract is a receiver function that substracts minutes to the clock
func (c Clock) Subtract(m int) Clock {
	return New(0, c.minutes-m)
}
