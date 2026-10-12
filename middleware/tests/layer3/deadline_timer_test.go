package layer3_test

import (
	"slices"
	"testing"
	"time"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer3"
)

type clockwork struct {
	delays    []time.Duration
	callbacks []func()
	timers    []*time.Timer
}

func (c *clockwork) afterFunc(delay time.Duration, callback func()) *time.Timer {
	c.delays = append(c.delays, delay)
	c.callbacks = append(c.callbacks, callback)
	timer := time.AfterFunc(time.Hour, func() {})
	c.timers = append(c.timers, timer)
	return timer
}

type timerHarness struct {
	clock       *clockwork
	submissions map[string]concepts.Submission
	outcome     string
	counted     []int
	settled     []string
	timers      *layer3.DeadlineTimer
}

func newTimerHarness() *timerHarness {
	h := &timerHarness{clock: &clockwork{}, submissions: map[string]concepts.Submission{}, outcome: concepts.Pending}
	h.timers = layer3.NewDeadlineTimer(layer3.DeadlineTimerDependencies{
		AfterFunc: h.clock.afterFunc,
		Get:       func(id string) (concepts.Submission, bool) { s, ok := h.submissions[id]; return s, ok },
		Evaluate: func(s concepts.Submission, voterCount int, now int64) string {
			h.counted = append(h.counted, voterCount)
			return h.outcome
		},
		Execute: func(s concepts.Submission) concepts.Submission {
			h.settled = append(h.settled, "executed")
			return s
		},
		Reject: func(s concepts.Submission) concepts.Submission {
			h.settled = append(h.settled, "rejected")
			return s
		},
		Now: func() time.Time { return time.Unix(1000, 0) },
	})
	return h
}

func (h *timerHarness) pending(id string, when int64, deadline int) concepts.Submission {
	s := concepts.Submission{Id: id, When: when, Resolution: concepts.Pending, Policy: concepts.Policy{DeadlineSeconds: &deadline}}
	h.submissions[id] = s
	return s
}

// a timer starts with the time left until the deadline, never negative
func TestTimers_StartsWithTimeLeft(t *testing.T) {
	h := newTimerHarness()
	h.timers.Start(h.pending("qap_1", 990, 60))
	h.timers.Start(h.pending("qap_2", 900, 60))

	if len(h.clock.delays) != 2 || h.clock.delays[0] != 50*time.Second || h.clock.delays[1] != 0 {
		t.Fatalf("got delays %v", h.clock.delays)
	}
}

// when the timer fires on a pending submission that still misses its criteria, it is rejected
func TestTimers_FiringRejectsPendingSubmission(t *testing.T) {
	h := newTimerHarness()
	h.timers.Start(h.pending("qap_1", 990, 60))
	h.outcome = concepts.Rejected
	h.clock.callbacks[0]()

	if len(h.settled) != 1 || h.settled[0] != "rejected" {
		t.Fatalf("settled %v", h.settled)
	}
}

// a submission still pending in its extension gets a new timer for the end of the extension, and is not rejected yet
func TestTimers_FiringIntoTheExtensionWaitsForItsEnd(t *testing.T) {
	h := newTimerHarness()
	s := h.pending("qap_1", 940, 60)
	thirty := 30
	s.Policy.ExtensionSeconds = &thirty
	h.submissions["qap_1"] = s
	h.timers.Start(s)
	h.clock.callbacks[0]()

	if len(h.settled) != 0 || len(h.clock.delays) != 2 || h.clock.delays[1] != 30*time.Second {
		t.Fatalf("settled %v, delays %v", h.settled, h.clock.delays)
	}
}

// when the timer fires, the submission is evaluated with the member count stored on it
func TestTimers_FiringUsesTheStoredMemberCount(t *testing.T) {
	h := newTimerHarness()
	s := h.pending("qap_1", 990, 60)
	s.VoterCount = 7
	h.submissions["qap_1"] = s
	h.timers.Start(s)
	h.clock.callbacks[0]()

	if !slices.Equal(h.counted, []int{7}) {
		t.Fatalf("evaluated with %v", h.counted)
	}
}

// when the timer fires on a pending submission that meets its criteria, it is carried out
func TestTimers_FiringExecutesWhenCriteriaAreMet(t *testing.T) {
	h := newTimerHarness()
	h.timers.Start(h.pending("qap_1", 990, 60))
	h.outcome = concepts.Executed
	h.clock.callbacks[0]()

	if len(h.settled) != 1 || h.settled[0] != "executed" {
		t.Fatalf("settled %v", h.settled)
	}
}

// a timer that fires on a submission that is gone or already decided does nothing
func TestTimers_FiringOnDecidedOrGoneDoesNothing(t *testing.T) {
	h := newTimerHarness()
	decided := h.pending("qap_1", 990, 60)
	h.timers.Start(decided)
	h.timers.Start(h.pending("qap_2", 990, 60))
	decided.Resolution = concepts.Executed
	h.submissions["qap_1"] = decided
	delete(h.submissions, "qap_2")
	h.clock.callbacks[0]()
	h.clock.callbacks[1]()

	if len(h.settled) != 0 {
		t.Fatalf("settled %v", h.settled)
	}
}

// cancelling stops the timer, and cancelling an unknown id is harmless
func TestTimers_Cancel(t *testing.T) {
	h := newTimerHarness()
	h.timers.Start(h.pending("qap_1", 990, 60))
	h.timers.Cancel("qap_1")
	h.timers.Cancel("qap_404")

	if h.clock.timers[0].Stop() {
		t.Fatal("expected the timer to be stopped already")
	}
}
