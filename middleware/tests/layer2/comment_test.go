package layer2_test

import (
	"strings"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
)

// a comment of exactly 500 characters is accepted, counted as characters and not as bytes
func TestComment_AcceptsTheLongestComment(t *testing.T) {
	h := newHarness()
	h.pending("qap_1", "alice", "alice")
	if r := h.send(t, "comment", "bob", map[string]string{"X-QAP-ID": "qap_1", "X-QAP-Comment": strings.Repeat("%C3%A9", 500)}, ""); r.status != 200 {
		t.Fatalf("got %d %s", r.status, r.raw)
	}
}

// comment refuses a request without an id, for an unknown submission, from someone without eligibility, without a readable comment, and with a comment over 500 characters
func TestComment_RefusesInvalidRequests(t *testing.T) {
	cases := []struct {
		user    string
		headers map[string]string
		status  int
	}{
		{"bob", map[string]string{"X-QAP-Comment": "hi"}, 406},
		{"bob", map[string]string{"X-QAP-ID": "qap_404", "X-QAP-Comment": "hi"}, 404},
		{"mallory", map[string]string{"X-QAP-ID": "qap_1", "X-QAP-Comment": "hi"}, 403},
		{"bob", map[string]string{"X-QAP-ID": "qap_1"}, 406},
		{"bob", map[string]string{"X-QAP-ID": "qap_1", "X-QAP-Comment": "50%"}, 400},
		{"bob", map[string]string{"X-QAP-ID": "qap_1", "X-QAP-Comment": strings.Repeat("%C3%A9", 501)}, 400},
	}
	for _, c := range cases {
		h := newHarness()
		h.pending("qap_1", "alice", "alice")
		if r := h.send(t, "comment", c.user, c.headers, ""); r.status != c.status {
			t.Errorf("%v: got %d %s", c, r.status, r.raw)
		}
	}
}

// a comment is read from the percent-encoded X-QAP-Comment header and stored with its author and time
func TestComment_StoresComment(t *testing.T) {
	h := newHarness()
	h.pending("qap_1", "alice", "alice")
	r := h.send(t, "comment", "bob", map[string]string{"X-QAP-ID": "qap_1", "X-QAP-Comment": "This%20copy%20is%20damaged%20beyond%20repair%20%E2%80%94%20sadly."}, "a body is ignored")

	s, _ := h.store.Get("qap_1")
	want := concepts.Comment{Author: "bob", At: 1760000000, Text: "This copy is damaged beyond repair — sadly."}
	if r.status != 200 || r.body["comments"] != float64(1) || len(s.Comments) != 1 || s.Comments[0] != want {
		t.Fatalf("got %d %s, comments %v", r.status, r.raw, s.Comments)
	}
}
