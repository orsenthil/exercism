package splitsecondstopwatch

import (
	"errors"
	"fmt"
	"time"
)

type SplitSecondStopwatch struct {
	state        string
	currentTime  time.Duration
	previousLaps []time.Duration
}

func (sss *SplitSecondStopwatch) Start() error {
	if sss.state == "running" {
		return errors.New("cannot start an already running stopwatch")
	}
	sss.state = "running"
	return nil
}

func (sss *SplitSecondStopwatch) Stop() error {
	if sss.state != "running" {
		return errors.New("cannot stop a stopwatch that is not running")
	}
	sss.state = "stopped"
	return nil
}

func (sss *SplitSecondStopwatch) Reset() error {
	if sss.state != "stopped" {
		return errors.New("cannot reset a stopwatch that is not stopped")
	}
	sss.state = "ready"
	sss.currentTime = 0
	sss.previousLaps = []time.Duration{}
	return nil
}

func (sss *SplitSecondStopwatch) Lap() error {
	if sss.state != "running" {
		return errors.New("cannot lap a stopwatch that is not running")
	}
	sss.previousLaps = append(sss.previousLaps, sss.currentTime)
	sss.currentTime = 0
	return nil
}

func (sss *SplitSecondStopwatch) AdvanceTime(by string) {
	var h, m, s int
	fmt.Sscanf(by, "%d:%d:%d", &h, &m, &s)
	if sss.state == "running" {
		duration := time.Duration(h*3600+m*60+s) * time.Second
		sss.currentTime += duration
	}
}

func (sss *SplitSecondStopwatch) State() string {
	return sss.state
}

func (sss *SplitSecondStopwatch) CurrentLap() string {
	return formatDuration(sss.currentTime)
}

func (sss *SplitSecondStopwatch) Total() string {
	total := sss.currentTime
	for _, lap := range sss.previousLaps {
		total += lap
	}
	return formatDuration(total)
}

func (sss *SplitSecondStopwatch) PreviousLaps() []string {
	previousLaps := make([]string, len(sss.previousLaps))
	for i, lap := range sss.previousLaps {
		previousLaps[i] = formatDuration(lap)
	}
	return previousLaps
}

func NewSplitSecondStopwatch() *SplitSecondStopwatch {
	return &SplitSecondStopwatch{
		state:        "ready",
		currentTime:  0,
		previousLaps: []time.Duration{},
	}
}

func formatDuration(d time.Duration) string {
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}