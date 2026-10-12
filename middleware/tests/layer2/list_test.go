package layer2_test

import "testing"

// list returns the pending submissions of the permission that the user may vote on
func TestList_ReturnsPendingSubmissions(t *testing.T) {
	h := newHarness()
	h.pending("qap_1", "alice", "alice")
	r := h.send(t, "list", "bob", nil, "")

	if r.status != 200 || len(r.list) != 1 {
		t.Fatalf("got %d %s", r.status, r.raw)
	}
}

// list returns an empty array to someone without eligibility
func TestList_IneligibleUserSeesNothing(t *testing.T) {
	h := newHarness()
	h.pending("qap_1", "alice", "alice")
	if r := h.send(t, "list", "mallory", nil, ""); r.status != 200 || r.raw != "[]" {
		t.Fatalf("got %d %s", r.status, r.raw)
	}
}
