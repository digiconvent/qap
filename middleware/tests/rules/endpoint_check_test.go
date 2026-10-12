package rules_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// a submission whose request the endpoint refuses is not stored, and the answer lists the problems
func TestEndpointCheck_RefusesTheSubmission(t *testing.T) {
	realm := New(t)
	realm.RefusesEveryRequest()
	proposed := realm.As("frodo").Proposes(CreateUser, PeerReview, `{"email":""}`)

	problems, _ := proposed.Body["problems"].(map[string]any)
	Expect(t, proposed.Status == 422 && proposed.Message() == "the request cannot be carried out" && problems["body"] != nil, "got %d %s", proposed.Status, proposed.Raw)
	listed := realm.As("samwise").Lists(CreateUser)
	Expect(t, len(listed.List) == 0, "nothing should be stored: %s", listed.Raw)
}

// a request that no longer passes its check when it is approved is not executed, and its result says why
func TestEndpointCheck_IsRepeatedBeforeExecution(t *testing.T) {
	realm := New(t)
	proposed := realm.As("frodo").Proposes(CreateUser, PeerReview, `{"email":"bilbo@shire.test"}`)
	realm.RefusesEveryRequest()
	voted := realm.As("samwise").VotesFavour(CreateUser, proposed.Id())

	Expect(t, voted.Resolution() == "executed" && voted.Result()["status"] == float64(422) && len(realm.Executed()) == 0, "got %s, executions %d", voted.Raw, len(realm.Executed()))
}

// a plain request without an intent is checked before the endpoint executes it
func TestEndpointCheck_GuardsPlainRequests(t *testing.T) {
	realm := New(t)
	Expect(t, realm.As("frodo").Request(SetFrodosPassword, nil, "").Raw == "reached", "a passing check reaches the endpoint")

	realm.RefusesEveryRequest()
	refused := realm.As("frodo").Request(SetFrodosPassword, nil, "")
	Expect(t, refused.Status == 422, "a failing check is refused: %d %s", refused.Status, refused.Raw)
}
