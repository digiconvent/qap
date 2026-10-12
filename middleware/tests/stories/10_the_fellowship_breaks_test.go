package stories_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// senior approval asks for three votes but the white council has only two members, so the votes of both members are enough
func TestStory10_TheFellowshipBreaks(t *testing.T) {
	realm := New(t)
	proposed := realm.As("aragorn").Proposes(DeleteFellowship, SeniorApproval, "")
	Expect(t, proposed.Resolution() == "pending", "one vote of two: %s", proposed.Raw)

	voted := realm.As("legolas").VotesFavour(DeleteFellowship, proposed.Id())
	Expect(t, voted.Resolution() == "executed" && len(realm.Executed()) == 1, "every member agrees: %s", voted.Raw)
}
