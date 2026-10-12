package rules_test

import (
	"testing"

	qap "github.com/digiconvent/qap/middleware"
	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// a route parameter id that is not a uuid is refused on every request
func TestEndpointInterface_IdMustBeUuid(t *testing.T) {
	realm := New(t)
	refused := realm.As("frodo").Discovers(UpdateWithoutUuid)
	Expect(t, refused.Status == 400 && refused.Message() == "the id must be a UUID", "got %d %s", refused.Status, refused.Raw)
}

// a body that is not the json payload of the endpoint is refused before the check
func TestEndpointInterface_PayloadMustDecode(t *testing.T) {
	realm := New(t)
	refused := realm.As("frodo").Proposes(CreateUser, PeerReview, `{"email":`)
	problems, _ := refused.Body["problems"].(map[string]any)
	Expect(t, refused.Status == 422 && problems["body"] != nil, "got %d %s", refused.Status, refused.Raw)
}

// the endpoint executes the decoded payload with the submission as context, the author as user and the route parameters
func TestEndpointInterface_ExecuteReceivesTheSubmission(t *testing.T) {
	realm := New(t)
	proposed := realm.As("frodo").Proposes(UpdateFrodo, PeerReview, `{"name":"Mr. Underhill"}`)
	realm.As("samwise").VotesFavour(UpdateFrodo, proposed.Id())

	executed := realm.Executed()
	Expect(t, len(executed) == 1, "executed %d", len(executed))
	request := executed[0]
	Expect(t, request.Payload["name"] == "Mr. Underhill" && request.User == "frodo" && request.Params["id"] == FrodoId, "request %+v", request)
	Expect(t, request.Submission != nil && request.Submission.Id == proposed.Id() && request.Submission.Permission == "iam.user.update", "submission %+v", request.Submission)
}

// an endpoint that reads its entity has the changed fields recorded in the audit trail, before and after
func TestEndpointInterface_ReadRecordsTheChanges(t *testing.T) {
	realm := New(t)
	proposed := realm.As("frodo").Proposes(UpdateFrodo, PeerReview, `{"name":"Mr. Underhill"}`)
	realm.As("samwise").VotesFavour(UpdateFrodo, proposed.Id())

	audited := realm.Audited()
	want := qap.Change{Before: "Frodo Baggins", After: "Mr. Underhill"}
	Expect(t, len(audited) == 1 && len(audited[0].Changes) == 1 && audited[0].Changes["name"] == want, "audited %+v", audited)
}

// an anonymous poll records no changes, only the author, the time and the outcome
func TestEndpointInterface_AnonymousPollRecordsNoChanges(t *testing.T) {
	realm := New(t)
	anonymous := PeerReview
	anonymous.Anonymous = true
	realm.SetPolicies("iam.user.update", anonymous)
	proposed := realm.As("frodo").Proposes(UpdateFrodo, anonymous, `{"name":"Mr. Underhill"}`)
	realm.As("samwise").VotesFavour(UpdateFrodo, proposed.Id())

	audited := realm.Audited()
	Expect(t, len(audited) == 1 && audited[0].Changes == nil && audited[0].Resolution == "executed", "audited %+v", audited)
}

// an endpoint without a reader records no changes
func TestEndpointInterface_WithoutReaderNoChanges(t *testing.T) {
	realm := New(t)
	proposed := realm.As("frodo").Proposes(RenameFrodo, PeerReview, `{"name":"Mr. Underhill"}`)
	realm.As("samwise").VotesFavour(RenameFrodo, proposed.Id())

	audited := realm.Audited()
	Expect(t, len(audited) == 1 && audited[0].Changes == nil, "audited %+v", audited)
}
