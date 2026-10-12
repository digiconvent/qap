package rules_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// half the court of four must agree and three must take part, and a withheld vote counts for neither
func TestPercentageQuorumAndWithhold(t *testing.T) {
	realm := New(t)
	proposed := realm.As("theoden").Proposes(MusterTheArmy, HalfTheCourt, "")

	withheld := realm.As("eomer").Withholds(MusterTheArmy, proposed.Id())
	Expect(t, withheld.Status == 200 && withheld.Votes("withhold") == 1 && withheld.Resolution() == "pending", "withhold: %s", withheld.Raw)

	half := realm.As("eowyn").VotesFavour(MusterTheArmy, proposed.Id())
	Expect(t, half.Resolution() == "pending", "two in favour is half the court, but only two took part: %s", half.Raw)

	decided := realm.As("wormtongue").VotesAgainst(MusterTheArmy, proposed.Id())
	Expect(t, decided.Resolution() == "executed" && len(realm.Executed()) == 1, "a third participant meets the quorum: %s", decided.Raw)
}
