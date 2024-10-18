package job

import "time"

type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}
