package stories_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// boromir reaches for an action of the white council and is refused at every turn
func TestStory09_BoromirTriesToTakeTheRingByForce(t *testing.T) {
	realm := New(t)
	boromir := realm.As("boromir")

	Expect(t, len(boromir.Discovers(DeleteBoromir).PolicyNames()) == 0, "boromir should see no policies")

	refused := boromir.Proposes(DeleteBoromir, SeniorApproval, "")
	Expect(t, refused.Status == 403, "boromir submit: %d %s", refused.Status, refused.Raw)

	proposed := realm.As("aragorn").Proposes(DeleteBoromir, SeniorApproval, "")
	voted := boromir.VotesFavour(DeleteBoromir, proposed.Id())
	Expect(t, voted.Status == 403 && voted.Message() == "You cannot vote on this", "boromir vote: %d %s", voted.Status, voted.Raw)
	polled := boromir.Polls(DeleteBoromir, proposed.Id())
	Expect(t, polled.Status == 403, "boromir poll: %d %s", polled.Status, polled.Raw)
}
