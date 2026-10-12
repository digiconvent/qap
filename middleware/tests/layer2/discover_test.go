package layer2_test

import (
	"testing"

	"github.com/digiconvent/qap/middleware/internal/layer2"
)

// discover returns the endpoint, the permission and the policies the user may use, with their settings
func TestDiscover_ReturnsEligiblePolicies(t *testing.T) {
	r := newHarness().send(t, "discover", "alice", nil, "")

	policies, _ := r.body["policies"].([]any)
	if r.status != 200 || r.body["endpoint"] != "/books/123" || r.body["permission"] != "library.book.123.delete" || len(policies) != 1 {
		t.Fatalf("got %d %s", r.status, r.raw)
	}
	policy, _ := policies[0].(map[string]any)
	if policy["id"] != "pol_librarians" || policy["name"] != "two_librarians" || policy["favour_threshold"] != float64(2) || policy["withhold_permitted"] != false {
		t.Fatalf("unexpected policy %v", policy)
	}
	if extension, present := policy["extension_seconds"]; !present || extension != nil {
		t.Fatalf("expected extension_seconds as null, got %v", policy)
	}
}

// discover leaves out the policies the user has no eligibility for
func TestDiscover_LeavesOutIneligiblePolicies(t *testing.T) {
	r := newHarness().send(t, "discover", "mallory", nil, "")

	if policies, _ := r.body["policies"].([]any); r.status != 200 || policies == nil || len(policies) != 0 {
		t.Fatalf("got %d %s", r.status, r.raw)
	}
}

// discover names the version of QAP in the X-QAP-Version response header
func TestDiscover_NamesTheVersion(t *testing.T) {
	h := newHarness()
	h.dependencies.Version = "1.2.3"
	h.intents = layer2.New(h.dependencies)
	r := h.send(t, "discover", "alice", nil, "")

	if r.headers.Get("X-QAP-Version") != "1.2.3" {
		t.Fatalf("got version %q", r.headers.Get("X-QAP-Version"))
	}
}
