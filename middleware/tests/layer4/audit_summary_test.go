package layer4_test

import (
	"reflect"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

func decided(commentsInAudit bool) concepts.Submission {
	return concepts.Submission{
		Id: "qap_1", Endpoint: "/api/library/book/123/delete", Body: "{}", Permission: "library.book.delete",
		Policy: concepts.Policy{Id: "pol_librarians", Name: "two_librarians", CommentsInAudit: commentsInAudit},
		Author: "alice", When: 1000, Resolution: concepts.Executed, Result: &concepts.Result{Status: 204},
		Favour: []string{"alice", "bob"}, Against: []string{}, Withhold: []string{},
		Comments: []concepts.Comment{{Author: "bob", Text: "damaged"}},
	}
}

// the audit summary records what was decided, when, by whom and with which outcome
func TestSummarize(t *testing.T) {
	summary := layer4.SummarizeAudit(decided(true), 1300)

	if summary.SubmissionId != "qap_1" || summary.Permission != "library.book.delete" || summary.Policy.Id != "pol_librarians" ||
		summary.Author != "alice" || summary.Anonymous || summary.Created != 1000 || summary.Decided != 1300 ||
		summary.Url != "/api/library/book/123/delete" || summary.Body != "{}" ||
		summary.Resolution != concepts.Executed || summary.Result.Status != 204 || len(summary.Favour) != 2 || len(summary.Comments) != 1 {
		t.Fatalf("got %+v", summary)
	}
}

// comments only appear in the audit summary when the policy includes them
func TestSummarize_LeavesOutComments(t *testing.T) {
	if summary := layer4.SummarizeAudit(decided(false), 1300); summary.Comments != nil {
		t.Fatalf("expected no comments, got %+v", summary.Comments)
	}
}

// an anonymous policy audits only the author, the time of the decision and the outcome
func TestSummarize_Anonymous(t *testing.T) {
	s := decided(true)
	s.Policy.Anonymous = true
	summary := layer4.SummarizeAudit(s, 1300)

	want := concepts.AuditSummary{Anonymous: true, Author: "alice", Decided: 1300, Resolution: concepts.Executed}
	if !reflect.DeepEqual(summary, want) {
		t.Fatalf("got %+v, want %+v", summary, want)
	}
}
