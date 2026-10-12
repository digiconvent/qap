package layer0_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// samwise withdraws his own proposal before anyone votes, which frodo cannot do for him
func TestTakeBackAProposal(t *testing.T) {
	realm := New(t)
	samwise := realm.As("samwise")
	proposed := samwise.Proposes(CreateUser, PeerReview, "")

	notTheAuthor := realm.As("frodo").Withdraws(CreateUser, proposed.Id())
	Expect(t, notTheAuthor.Status == 403, "frodo withdraw: %d %s", notTheAuthor.Status, notTheAuthor.Raw)

	withdrawn := samwise.Withdraws(CreateUser, proposed.Id())
	Expect(t, withdrawn.Status == 200 && withdrawn.Raw == "", "withdraw: %d %q", withdrawn.Status, withdrawn.Raw)

	polled := samwise.Polls(CreateUser, proposed.Id())
	Expect(t, polled.Status == 404, "poll after withdraw: %d %s", polled.Status, polled.Raw)
}
