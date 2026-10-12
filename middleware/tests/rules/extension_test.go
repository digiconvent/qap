package rules_test

import (
	"testing"
	"time"

	qap "github.com/digiconvent/qap/middleware"
	. "github.com/digiconvent/qap/middleware/tests/harness"
)

func patientCourt() qap.Policy {
	deadline, extension := 60, 30
	return qap.Policy{Id: "the patient court", Name: "the patient court", Groups: []string{CourtOfEdoras}, FavourThreshold: 3, DeadlineSeconds: &deadline, ExtensionSeconds: &extension}
}

// when nobody votes against, a submission that misses its threshold by the deadline still passes at the end of the extension
func TestExtension_SilencePasses(t *testing.T) {
	realm := New(t)
	realm.SetPolicies("iam.group.remove-user", patientCourt())
	proposed := realm.As("theoden").Proposes(RemoveEomer, patientCourt(), "")

	realm.Waits(60 * time.Second)
	Expect(t, realm.As("theoden").Polls(RemoveEomer, proposed.Id()).Resolution() == "pending", "still pending in the extension")

	realm.Waits(30 * time.Second)
	polled := realm.As("theoden").Polls(RemoveEomer, proposed.Id())
	Expect(t, polled.Resolution() == "executed" && len(realm.Executed()) == 1, "after the extension: %s", polled.Raw)
}

// a single vote against during the extension rejects the submission
func TestExtension_AVoteAgainstRejects(t *testing.T) {
	realm := New(t)
	realm.SetPolicies("iam.group.remove-user", patientCourt())
	proposed := realm.As("theoden").Proposes(RemoveEomer, patientCourt(), "")

	realm.Waits(70 * time.Second)
	voted := realm.As("eomer").VotesAgainst(RemoveEomer, proposed.Id())
	realm.Waits(20 * time.Second)

	Expect(t, voted.Resolution() == "rejected" && len(realm.Executed()) == 0, "vote against: %s", voted.Raw)
}
