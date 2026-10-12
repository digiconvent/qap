package stories_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// boromir objects, but peer review has no against threshold, so frodo's favour still carries the proposal
func TestStory05_BoromirObjects(t *testing.T) {
	realm := New(t)
	proposed := realm.As("samwise").Proposes(CreateUser, PeerReview, "")
	boromir := realm.As("boromir")

	objected := boromir.VotesAgainst(CreateUser, proposed.Id())
	Expect(t, objected.Resolution() == "pending" && objected.Votes("against") == 1, "boromir against: %s", objected.Raw)

	withheld := boromir.Withholds(CreateUser, proposed.Id())
	Expect(t, withheld.Status == 400 && withheld.Message() == "withhold is not permitted for this poll", "boromir withhold: %d %s", withheld.Status, withheld.Raw)

	carried := realm.As("frodo").VotesFavour(CreateUser, proposed.Id())
	Expect(t, carried.Resolution() == "executed" && carried.Votes("against") == 1 && carried.Votes("favour") == 2, "frodo favour: %s", carried.Raw)
}
