package rules_test

import (
	"testing"
	"time"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// a decided submission disappears once nobody has polled it for a minute, while a pending one stays
func TestCleanup(t *testing.T) {
	realm := New(t)
	samwise := realm.As("samwise")
	decided := samwise.Proposes(CreateUser, FastTrack, "")
	pending := samwise.Proposes(UpdateFrodo, PeerReview, "")

	realm.Waits(59 * time.Second)
	realm.CleansUp()
	Expect(t, samwise.Polls(CreateUser, decided.Id()).Status == 200, "still there before a minute has passed")

	realm.Waits(61 * time.Second)
	realm.CleansUp()
	Expect(t, samwise.Polls(CreateUser, decided.Id()).Status == 404, "gone after a minute without polls")
	Expect(t, samwise.Polls(UpdateFrodo, pending.Id()).Status == 200, "a pending submission is never removed")
}
