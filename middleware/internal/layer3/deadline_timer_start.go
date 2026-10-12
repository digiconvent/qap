package layer3

import (
	"time"

	"github.com/digiconvent/qap/middleware/internal/concepts"
)

func (t *DeadlineTimer) Start(s concepts.Submission) {
	now := t.dependencies.Now().Unix()
	end := s.When + int64(*s.Policy.DeadlineSeconds)
	if s.Policy.ExtensionSeconds != nil && now >= end {
		end += int64(*s.Policy.ExtensionSeconds)
	}
	delay := time.Duration(max(end-now, 0)) * time.Second
	id := s.Id
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.timers[id] = t.dependencies.AfterFunc(delay, func() { t.fire(id) })
}
