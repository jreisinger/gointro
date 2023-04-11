// Package stopwatch measure elapsed time.
package stopwatch

import (
	"time"
)

type StopWatch struct {
	start time.Time
}

func Start() *StopWatch {
	s := StopWatch{start: time.Now()}
	return &s
}

func (s *StopWatch) Stop() time.Duration {
	elapsed := time.Since(s.start)
	return elapsed
}
