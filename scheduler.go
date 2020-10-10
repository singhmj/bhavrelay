package main

import (
	"time"
)

// IntervalPeriod :
const IntervalPeriod time.Duration = 24 * time.Hour

type jobTicker struct {
	t            *time.Timer
	hourToTick   int
	minuteToTick int
	secondToTick int
}

func getNextTickDuration(hour, minute, second int) time.Duration {
	now := time.Now()
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, second, 0, time.Local)
	if nextTick.Before(now) {
		nextTick = nextTick.Add(IntervalPeriod)
	}
	return nextTick.Sub(time.Now())
}

func NewJobTicker(hour, minute, second int) jobTicker {
	job := jobTicker{
		t:            time.NewTimer(getNextTickDuration(hour, minute, second)),
		hourToTick:   hour,
		minuteToTick: minute,
		secondToTick: second,
	}
	return job
}

func (job jobTicker) updateJobTicker() {
	job.t.Reset(getNextTickDuration(job.hourToTick, job.minuteToTick, job.secondToTick))
}

// startScheduler :
func startScheduler() error {
	config := GetConfig()
	jobTicker := NewJobTicker(config.Dumper.Scheduler.Hour, config.Dumper.Scheduler.Minute, config.Dumper.Scheduler.Second)

	for {
		<-jobTicker.t.C
		err := restoreCacheWithBhavData(time.Now())
		if err != nil {
			return err
		}
		jobTicker.updateJobTicker()
	}
}
