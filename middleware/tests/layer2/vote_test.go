package layer2_test

import (
	"slices"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer2"
)

func vote(value string) map[string]string {
	return map[string]string{"X-QAP-ID": "qap_1", "X-QAP-Vote": value}
}

// vote refuses requests in the order of its errors: no id, unknown, resolved, not eligible, invalid vote, withhold not permitted
func TestVote_RefusesInvalidRequests(t *testing.T) {
	cases := []struct {
		name       string
		resolution string
		user       string
		headers    map[string]string
		status     int
	}{
		{"no id", concepts.Pending, "bob", map[string]string{"X-QAP-Vote": "1"}, 406},
		{"unknown", concepts.Pending, "bob", map[string]string{"X-QAP-ID": "qap_404", "X-QAP-Vote": "1"}, 404},
		{"resolved", concepts.Executed, "bob", vote("1"), 409},
		{"not eligible", concepts.Pending, "mallory", vote("1"), 403},
		{"invalid vote", concepts.Pending, "bob", vote("x"), 400},
		{"withhold not permitted", concepts.Pending, "bob", vote("0"), 400},
	}
	for _, c := range cases {
		h := newHarness()
		s := h.pending("qap_1", "alice", "alice")
		s.Resolution = c.resolution
		h.store.Create(s)
		if r := h.send(t, "vote", c.user, c.headers, ""); r.status != c.status {
			t.Errorf("%s: got %d %s", c.name, r.status, r.raw)
		}
	}
}

// a vote is recorded, and a submission that stays undecided stays pending
func TestVote_RecordsVote(t *testing.T) {
	permitting := twoLibrarians
	permitting.WithholdPermitted = true
	h := withPolicy(newHarness(), permitting)
	s := h.pending("qap_1", "alice", "alice")
	s.Policy = permitting
	h.store.Create(s)
	r := h.send(t, "vote", "bob", vote("0"), "")

	if r.status != 200 || r.body["withhold"] != float64(1) || r.body["resolution"] != concepts.Pending || len(h.settled) != 0 {
		t.Fatalf("got %d %s, settled %v", r.status, r.raw, h.settled)
	}
}

// the deciding favour vote carries the submission out and cancels its timer
func TestVote_DecidingFavourExecutes(t *testing.T) {
	h := newHarness()
	h.pending("qap_1", "alice", "alice")
	r := h.send(t, "vote", "bob", vote("1"), "")

	if r.status != 200 || r.body["resolution"] != concepts.Executed || !slices.Equal(h.settled, []string{"executed:qap_1"}) || !slices.Equal(h.deadlines.cancelled, []string{"qap_1"}) {
		t.Fatalf("got %d %s, settled %v, cancelled %v", r.status, r.raw, h.settled, h.deadlines.cancelled)
	}
}

// the deciding against vote rejects the submission and cancels its timer
func TestVote_DecidingAgainstRejects(t *testing.T) {
	h := newHarness()
	h.pending("qap_1", "alice", "alice")
	r := h.send(t, "vote", "bob", vote("-1"), "")

	if r.status != 200 || r.body["resolution"] != concepts.Rejected || !slices.Equal(h.settled, []string{"rejected:qap_1"}) || !slices.Equal(h.deadlines.cancelled, []string{"qap_1"}) {
		t.Fatalf("got %d %s, settled %v, cancelled %v", r.status, r.raw, h.settled, h.deadlines.cancelled)
	}
}

// a vote that loses the race against another deciding vote is refused as already resolved
func TestVote_LosingRaceIsRefused(t *testing.T) {
	h := newHarness()
	h.pending("qap_1", "alice", "alice")
	h.intents = racingIntents(h)
	if r := h.send(t, "vote", "bob", vote("1"), ""); r.status != 409 {
		t.Fatalf("got %d %s", r.status, r.raw)
	}
}

// a decision that lost to a newer vote leaves the submission pending, and its deadline keeps running
func TestVote_OutdatedDecisionKeepsTheDeadline(t *testing.T) {
	h := newHarness()
	h.dependencies.Execute = func(s concepts.Submission) concepts.Submission { return s }
	h.intents = layer2.New(h.dependencies)
	h.pending("qap_1", "alice", "alice")
	r := h.send(t, "vote", "bob", map[string]string{"X-QAP-ID": "qap_1", "X-QAP-Vote": "1"}, "")

	if r.status != 200 || r.body["resolution"] != concepts.Pending || len(h.deadlines.cancelled) != 0 {
		t.Fatalf("got %d %v, cancelled %v", r.status, r.body, h.deadlines.cancelled)
	}
}
