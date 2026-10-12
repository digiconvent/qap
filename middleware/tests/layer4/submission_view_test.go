package layer4_test

import (
	"encoding/json"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

func shape(t *testing.T, view any) map[string]any {
	t.Helper()
	raw, _ := json.Marshal(view)
	result := map[string]any{}
	json.Unmarshal(raw, &result)
	return result
}

// the view has the submission and its poll, with vote counts and comments
func TestSubmissionView_Shape(t *testing.T) {
	s := concepts.Submission{
		Id: "qap_1", Endpoint: "/api/library/book/123/delete", Method: "DELETE", Body: "",
		Permission: "library.book.delete", Policy: concepts.Policy{Id: "pol_librarians", Name: "two_librarians"},
		Author: "alice", When: 1760000000, Resolution: concepts.Pending,
		Favour: []string{"alice", "bob"}, Against: []string{"carol"}, Comments: []concepts.Comment{},
	}
	view := shape(t, layer4.SubmissionView(s, "bob"))

	part, _ := view["submission"].(map[string]any)
	p, _ := view["poll"].(map[string]any)
	policy, _ := part["policy"].(map[string]any)
	comments, _ := p["comments"].([]any)
	if part["id"] != "qap_1" || part["endpoint"] != "/api/library/book/123/delete" || part["method"] != "DELETE" || part["permission"] != "library.book.delete" ||
		part["author"] != "alice" || part["when"] != float64(1760000000) || part["resolution"] != "pending" || part["result"] != nil || policy["id"] != "pol_librarians" {
		t.Fatalf("unexpected submission %v", part)
	}
	if p["favour"] != float64(2) || p["against"] != float64(1) || p["withhold"] != float64(0) || comments == nil || p["resolution"] != "pending" {
		t.Fatalf("unexpected poll %v", p)
	}
}

// only the author sees the data of a sensitive read
func TestSubmissionView_SensitiveReadDataIsForTheAuthorOnly(t *testing.T) {
	s := concepts.Submission{Method: "GET", Author: "alice", Resolution: concepts.Executed, Result: &concepts.Result{Status: 200, Body: "history"}}

	author, _ := shape(t, layer4.SubmissionView(s, "alice"))["submission"].(map[string]any)
	other, _ := shape(t, layer4.SubmissionView(s, "bob"))["submission"].(map[string]any)
	authorResult, _ := author["result"].(map[string]any)
	otherResult, _ := other["result"].(map[string]any)
	if authorResult["body"] != "history" || otherResult["body"] != nil || otherResult["status"] != float64(200) {
		t.Fatalf("author %v, other %v", authorResult, otherResult)
	}
	if s.Result.Body != "history" {
		t.Fatal("the stored result must not change")
	}
}
