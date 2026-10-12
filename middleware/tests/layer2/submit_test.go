package layer2_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer2"
)

// submit without a policy id is refused, and lists the policies the user may pick
func TestSubmit_MissingPolicyIdListsAcceptablePolicies(t *testing.T) {
	r := newHarness().send(t, "submit", "alice", nil, "")

	acceptable, _ := r.body["acceptable_policies"].([]any)
	if r.status != 406 || len(acceptable) != 1 {
		t.Fatalf("got %d %s", r.status, r.raw)
	}
}

// submit under a policy that does not include the user is refused
func TestSubmit_PolicyThatDoesNotIncludeUserIsRefused(t *testing.T) {
	for _, c := range []struct{ user, policy string }{{"mallory", "pol_librarians"}, {"alice", "pol_unknown"}} {
		r := newHarness().send(t, "submit", c.user, map[string]string{"X-QAP-Policy-ID": c.policy}, "")
		if r.status != 403 || r.error() != "Policy does not include you" {
			t.Errorf("%v: got %d %s", c, r.status, r.raw)
		}
	}
}

// submit stores the request with its author, the favour of its author and the chosen policy
func TestSubmit_StoresSubmission(t *testing.T) {
	h := newHarness()
	r := h.send(t, "submit", "alice", map[string]string{"X-QAP-Policy-ID": "pol_librarians"}, `{"reason":"damaged"}`)

	s, ok := h.store.Get("qap_1")
	if r.status != 202 || !ok {
		t.Fatalf("got %d %s", r.status, r.raw)
	}
	if s.Author != "alice" || s.Method != "DELETE" || s.Endpoint != "/books/123" || s.Body != `{"reason":"damaged"}` ||
		s.Permission != "library.book.123.delete" || s.Route != "/books/:id" || s.Policy.Id != "pol_librarians" || s.When != 1760000000 ||
		s.Resolution != concepts.Pending || !slices.Equal(s.Favour, []string{"alice"}) {
		t.Fatalf("unexpected submission %+v", s)
	}
	if r.body["id"] != "qap_1" || r.body["viewer"] != "alice" {
		t.Fatalf("expected the view of the new submission, got %s", r.raw)
	}
}

// a submission whose request the endpoint refuses is not stored, and the answer lists the problems
func TestSubmit_RefusedByTheEndpoint(t *testing.T) {
	h := newHarness()
	h.endpoint = stubEndpoint{problems: map[string][]error{"reason": {errors.New("is required")}}}
	r := h.send(t, "submit", "alice", map[string]string{"X-QAP-Policy-ID": "pol_librarians"}, "")

	problems, _ := r.body["problems"].(map[string]any)
	if _, stored := h.store.Get("qap_1"); r.status != 422 || stored || problems["reason"] == nil {
		t.Fatalf("got %d %s, stored %v", r.status, r.raw, stored)
	}
}

// the stored submission keeps the endpoint and the route parameters, so it can be carried out later
func TestSubmit_KeepsTheEndpointAndParameters(t *testing.T) {
	h := newHarness()
	h.send(t, "submit", "alice", map[string]string{"X-QAP-Policy-ID": "pol_librarians"}, "")

	s, _ := h.store.Get("qap_1")
	if s.Handler.Execute == nil || s.Params["id"] != "123" {
		t.Fatalf("got params %v, handler kept %v", s.Params, s.Handler.Execute != nil)
	}
}

// the member count is read once, when the submission is created, and every later evaluation uses the stored count
func TestSubmit_StoresTheMemberCount(t *testing.T) {
	h := newHarness()
	calls := 0
	var evaluated []int
	h.dependencies.VoterCount = func(policy concepts.Policy) int {
		calls++
		return 5
	}
	h.dependencies.Evaluate = func(s concepts.Submission, voterCount int, now int64) string {
		evaluated = append(evaluated, voterCount)
		return concepts.Pending
	}
	h.intents = layer2.New(h.dependencies)
	h.send(t, "submit", "alice", map[string]string{"X-QAP-Policy-ID": "pol_librarians"}, "")
	h.send(t, "vote", "bob", vote("1"), "")

	s, _ := h.store.Get("qap_1")
	if calls != 1 || s.VoterCount != 5 || !slices.Equal(evaluated, []int{5, 5}) {
		t.Fatalf("calls %d, stored %d, evaluated with %v", calls, s.VoterCount, evaluated)
	}
}

// a submission whose policy has a deadline starts a deadline timer
func TestSubmit_StartsDeadlineTimer(t *testing.T) {
	h := newHarness()
	deadline := 60
	twoLibrariansWithDeadline := twoLibrarians
	twoLibrariansWithDeadline.DeadlineSeconds = &deadline
	h = withPolicy(h, twoLibrariansWithDeadline)
	h.send(t, "submit", "alice", map[string]string{"X-QAP-Policy-ID": "pol_librarians"}, "")

	if !slices.Equal(h.deadlines.started, []string{"qap_1"}) {
		t.Fatalf("expected a deadline timer, got %v", h.deadlines.started)
	}
}

// a submission whose policy has no deadline starts no timer
func TestSubmit_WithoutDeadlineStartsNoTimer(t *testing.T) {
	h := newHarness()
	h.send(t, "submit", "alice", map[string]string{"X-QAP-Policy-ID": "pol_librarians"}, "")

	if len(h.deadlines.started) != 0 {
		t.Fatalf("expected no timer, got %v", h.deadlines.started)
	}
}

// a submission that one vote already decides is carried out straight away
func TestSubmit_DecidedStraightAwayIsExecuted(t *testing.T) {
	oneLibrarian := twoLibrarians
	oneLibrarian.FavourThreshold = 1
	h := withPolicy(newHarness(), oneLibrarian)
	r := h.send(t, "submit", "alice", map[string]string{"X-QAP-Policy-ID": "pol_librarians"}, "")

	if r.status != 202 || r.body["resolution"] != concepts.Executed || !slices.Equal(h.settled, []string{"executed:qap_1"}) {
		t.Fatalf("got %d %s, settled %v", r.status, r.raw, h.settled)
	}
}
