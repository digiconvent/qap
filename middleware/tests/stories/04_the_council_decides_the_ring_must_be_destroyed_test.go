package stories_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// samwise proposes and frodo agrees, which is the second vote peer review needs
func TestStory04_TheCouncilDecidesTheRingMustBeDestroyed(t *testing.T) {
	realm := New(t)
	proposed := realm.As("samwise").Proposes(CreateUser, PeerReview, "")
	voted := realm.As("frodo").VotesFavour(CreateUser, proposed.Id())

	Expect(t, voted.Resolution() == "executed" && len(realm.Executed()) == 1, "vote: %s, executions %d", voted.Raw, len(realm.Executed()))
}
