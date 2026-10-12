package rules_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// a policy that changes while a submission waits only applies to new submissions
func TestPolicyChange(t *testing.T) {
	realm := New(t)
	proposed := realm.As("samwise").Proposes(UpdateFrodo, PeerReview, "")

	stricter := PeerReview
	stricter.FavourThreshold = 3
	realm.SetPolicies("iam.user.update", stricter)

	voted := realm.As("frodo").VotesFavour(UpdateFrodo, proposed.Id())
	Expect(t, voted.Resolution() == "executed", "the waiting submission keeps its two-vote threshold: %s", voted.Raw)

	next := realm.As("samwise").Proposes(UpdateFrodo, stricter, "")
	Expect(t, realm.As("frodo").VotesFavour(UpdateFrodo, next.Id()).Resolution() == "pending", "a new submission needs three votes")
}
