package stories_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// two favour votes are not enough until a third member takes part, even when that member votes against
func TestStory16_TheEntmootDecidesWhetherToMarch(t *testing.T) {
	realm := New(t)
	proposed := realm.As("samwise").Proposes(UpdatePolicy, FullFellowship, "")

	waiting := realm.As("frodo").VotesFavour(UpdatePolicy, proposed.Id())
	Expect(t, waiting.Resolution() == "pending", "two favour votes without quorum: %s", waiting.Raw)

	decided := realm.As("boromir").VotesAgainst(UpdatePolicy, proposed.Id())
	Expect(t, decided.Resolution() == "executed" && len(realm.Executed()) == 1, "quorum reached: %s", decided.Raw)
}
