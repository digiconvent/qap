package layer3

import (
	"sync"
	"time"

	"github.com/digiconvent/qap/middleware/internal/concepts"
)

type DeadlineTimerDependencies struct {
	AfterFunc func(delay time.Duration, callback func()) *time.Timer
	Get       func(id string) (concepts.Submission, bool)
	Evaluate  func(s concepts.Submission, voterCount int, now int64) string
	Execute   func(s concepts.Submission) concepts.Submission
	Reject    func(s concepts.Submission) concepts.Submission
	Now       func() time.Time
}

type DeadlineTimer struct {
	mutex        sync.Mutex
	dependencies DeadlineTimerDependencies
	timers       map[string]*time.Timer
}

func NewDeadlineTimer(dependencies DeadlineTimerDependencies) *DeadlineTimer {
	return &DeadlineTimer{dependencies: dependencies, timers: map[string]*time.Timer{}}
}
