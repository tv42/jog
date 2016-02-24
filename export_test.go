package jog

import "time"

// SetClock is only intended for unit tests.
func (l *Logger) SetClock(clock func() time.Time) {
	l.clock = clock
}
