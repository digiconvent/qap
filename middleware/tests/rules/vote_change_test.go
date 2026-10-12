package rules_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// boromir changes his mind, and his new vote replaces the old one
func TestVoteChange(t *testing.T) {
	realm := New(t)
	proposed := realm.As("samwise").Proposes(CreateUser, PeerReview, "")
	boromir := realm.As("boromir")

	boromir.VotesAgainst(CreateUser, proposed.Id())
	changed := boromir.VotesFavour(CreateUser, proposed.Id())
	Expect(t, changed.Resolution() == "executed" && changed.Votes("against") == 0 && changed.Votes("favour") == 2, "vote change: %s", changed.Raw)
}
