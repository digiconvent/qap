package harness

import (
	"slices"
	"sync"
	"time"
)

type scheduled struct {
	at       time.Time
	callback func()
}

type clock struct {
	mutex   sync.Mutex
	now     time.Time
	pending []scheduled
	ticks   chan time.Time
}

func (c *clock) Now() time.Time {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.now
}

func (c *clock) AfterFunc(delay time.Duration, callback func()) *time.Timer {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.pending = append(c.pending, scheduled{at: c.now.Add(delay), callback: callback})
	return time.AfterFunc(24*time.Hour, func() {})
}

func (c *clock) advance(duration time.Duration) {
	c.mutex.Lock()
	c.now = c.now.Add(duration)
	var due []scheduled
	c.pending = slices.DeleteFunc(c.pending, func(s scheduled) bool {
		if !s.at.After(c.now) {
			due = append(due, s)
			return true
		}
		return false
	})
	c.mutex.Unlock()
	for _, s := range due {
		s.callback()
	}
}
