package layer2_test

import "testing"

// a vote, comment or withdrawal on a submission that was withdrawn at the same moment answers unknown submission
func TestIntents_SubmissionWithdrawnAtTheSameMomentIsUnknown(t *testing.T) {
	requests := []struct {
		name    string
		user    string
		headers map[string]string
	}{
		{"vote", "bob", vote("1")},
		{"comment", "bob", map[string]string{"X-QAP-ID": "qap_1", "X-QAP-Comment": "hi"}},
		{"delete-submission", "alice", map[string]string{"X-QAP-ID": "qap_1"}},
	}
	for _, request := range requests {
		h := newHarness()
		h.pending("qap_1", "alice", "alice")
		h.intents = vanishingIntents(h)
		if r := h.send(t, request.name, request.user, request.headers, "hi"); r.status != 404 {
			t.Errorf("%s: got %d %s", request.name, r.status, r.raw)
		}
	}
}
