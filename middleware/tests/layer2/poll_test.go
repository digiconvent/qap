package layer2_test

import "testing"

// poll without a submission id is refused
func TestPoll_MissingIdIsRefused(t *testing.T) {
	if r := newHarness().send(t, "poll", "alice", nil, ""); r.status != 406 {
		t.Fatalf("got %d %s", r.status, r.raw)
	}
}

// poll of an unknown submission is refused
func TestPoll_UnknownSubmissionIsRefused(t *testing.T) {
	if r := newHarness().send(t, "poll", "alice", map[string]string{"X-QAP-ID": "qap_404"}, ""); r.status != 404 {
		t.Fatalf("got %d %s", r.status, r.raw)
	}
}

// poll by someone without eligibility is refused
func TestPoll_IneligibleUserIsRefused(t *testing.T) {
	h := newHarness()
	h.pending("qap_1", "alice", "alice")
	if r := h.send(t, "poll", "mallory", map[string]string{"X-QAP-ID": "qap_1"}, ""); r.status != 403 {
		t.Fatalf("got %d %s", r.status, r.raw)
	}
}

// poll returns the view of the submission for the user who asks, and restarts its expiry clock
func TestPoll_ReturnsViewAndTouches(t *testing.T) {
	h := newHarness()
	h.pending("qap_1", "alice", "alice")
	r := h.send(t, "poll", "bob", map[string]string{"X-QAP-ID": "qap_1"}, "")

	if r.status != 200 || r.body["id"] != "qap_1" || r.body["viewer"] != "bob" || h.store.touched["qap_1"] != 1760000000 {
		t.Fatalf("got %d %s, touched %v", r.status, r.raw, h.store.touched)
	}
}
