package layer2_test

import (
	"slices"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer2"
)

// withdraw refuses a request without an id, for an unknown submission, from someone else, or for a decided submission
func TestWithdraw_RefusesInvalidRequests(t *testing.T) {
	cases := []struct {
		user       string
		headers    map[string]string
		resolution string
		status     int
	}{
		{"alice", nil, concepts.Pending, 406},
		{"alice", map[string]string{"X-QAP-ID": "qap_404"}, concepts.Pending, 404},
		{"bob", map[string]string{"X-QAP-ID": "qap_1"}, concepts.Pending, 403},
		{"alice", map[string]string{"X-QAP-ID": "qap_1"}, concepts.Rejected, 409},
	}
	for _, c := range cases {
		h := newHarness()
		s := h.pending("qap_1", "alice", "alice")
		s.Resolution = c.resolution
		h.store.Create(s)
		if r := h.send(t, "delete-submission", c.user, c.headers, ""); r.status != c.status {
			t.Errorf("%v: got %d %s", c, r.status, r.raw)
		}
	}
}

// the author withdraws a pending submission: it is removed, its timer is cancelled and the body is empty
func TestWithdraw_RemovesSubmission(t *testing.T) {
	h := newHarness()
	h.pending("qap_1", "alice", "alice")
	r := h.send(t, "delete-submission", "alice", map[string]string{"X-QAP-ID": "qap_1"}, "")

	if _, ok := h.store.Get("qap_1"); r.status != 200 || r.raw != "" || ok || !slices.Equal(h.deadlines.cancelled, []string{"qap_1"}) {
		t.Fatalf("got %d %q, still stored %v, cancelled %v", r.status, r.raw, ok, h.deadlines.cancelled)
	}
}

// a withdrawal is recorded in the audit trail
func TestWithdraw_IsAudited(t *testing.T) {
	h := newHarness()
	var withdrawn []string
	h.dependencies.RecordWithdrawal = func(s concepts.Submission) { withdrawn = append(withdrawn, s.Id) }
	h.intents = layer2.New(h.dependencies)
	h.pending("qap_1", "alice", "alice")
	h.send(t, "delete-submission", "alice", map[string]string{"X-QAP-ID": "qap_1"}, "")

	if !slices.Equal(withdrawn, []string{"qap_1"}) {
		t.Fatalf("recorded %v", withdrawn)
	}
}

// a withdrawal that loses the race against a deciding vote is refused as already resolved
func TestWithdraw_LosingRaceIsRefused(t *testing.T) {
	h := newHarness()
	h.pending("qap_1", "alice", "alice")
	h.intents = racingIntents(h)
	if r := h.send(t, "delete-submission", "alice", map[string]string{"X-QAP-ID": "qap_1"}, ""); r.status != 409 {
		t.Fatalf("got %d %s", r.status, r.raw)
	}
}
