package layer0_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// frodo lists what waits for his decision and looks at the proposal of samwise, while aragorn sees nothing to vote on
func TestReviewPendingWork(t *testing.T) {
	realm := New(t)
	proposed := realm.As("samwise").Proposes(CreateUser, PeerReview, "")

	listed := realm.As("frodo").Lists(CreateUser)
	Expect(t, listed.Status == 200 && len(listed.List) == 1, "frodo list: %d %s", listed.Status, listed.Raw)

	outsider := realm.As("aragorn").Lists(CreateUser)
	Expect(t, outsider.Status == 200 && outsider.Raw == "[]", "aragorn list: %d %s", outsider.Status, outsider.Raw)

	polled := realm.As("frodo").Polls(CreateUser, proposed.Id())
	Expect(t, polled.Status == 200 && polled.Id() == proposed.Id(), "poll: %d %s", polled.Status, polled.Raw)
}
