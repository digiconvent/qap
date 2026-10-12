package stories_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// the king's council needs three in favour, and with theoden against grima's proposal stalls
func TestStory13_GrimasAttemptToBanishEomerStalls(t *testing.T) {
	realm := New(t)
	proposed := realm.As("wormtongue").Proposes(RemoveEomer, KingsCouncil, `{"user":"eomer"}`)
	realm.As("eowyn").VotesFavour(RemoveEomer, proposed.Id())
	voted := realm.As("theoden").VotesAgainst(RemoveEomer, proposed.Id())

	Expect(t, voted.Resolution() == "pending" && voted.Votes("favour") == 2 && voted.Votes("against") == 1 && len(realm.Executed()) == 0, "vote: %s", voted.Raw)
}
