package main

import (
	"time"
)

type Task struct { 
	taskName string
	startTime time.Time
	workTime int
	breakTime int
	cycleCount int
	isWorkTime bool
}

type Time struct {
	minutes int
	seconds int
}

func (t Task) getElapsedTimeInSeconds() int {
	return int(time.Since(t.startTime).Seconds())
}
