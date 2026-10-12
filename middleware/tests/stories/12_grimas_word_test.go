package stories_test

import (
	"testing"
	"time"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// under the plan, the regent's word runs out after 60 seconds without enough votes and the proposal of grima is rejected, unlike the seed title
func TestStory12_GrimasWordRunsOut(t *testing.T) {
	realm := New(t)
	proposed := realm.As("wormtongue").Proposes(RemoveEomer, RegentsWord, `{"user":"eomer"}`)

	realm.Waits(59 * time.Second)
	Expect(t, realm.As("wormtongue").Polls(RemoveEomer, proposed.Id()).Resolution() == "pending", "still pending before the deadline")

	realm.Waits(time.Second)
	polled := realm.As("wormtongue").Polls(RemoveEomer, proposed.Id())
	Expect(t, polled.Resolution() == "rejected" && len(realm.Executed()) == 0, "poll after deadline: %s", polled.Raw)
	Expect(t, len(realm.Audited()) == 1 && realm.Audited()[0].Resolution == "rejected", "audited %+v", realm.Audited())
}

// the regent's word is carried out when the court agrees before the deadline, and the deadline then changes nothing
func TestStory12_TheCourtAgreesInTime(t *testing.T) {
	realm := New(t)
	proposed := realm.As("wormtongue").Proposes(RemoveEomer, RegentsWord, "")
	realm.As("eowyn").VotesFavour(RemoveEomer, proposed.Id())
	voted := realm.As("theoden").VotesFavour(RemoveEomer, proposed.Id())
	realm.Waits(time.Minute)

	Expect(t, voted.Resolution() == "executed" && len(realm.Executed()) == 1 && len(realm.Audited()) == 1, "vote: %s, executions %d, audited %d", voted.Raw, len(realm.Executed()), len(realm.Audited()))
}
