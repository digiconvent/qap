package stories_test

import (
	"slices"
	"testing"
	"time"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// every decision is written down: what was asked, under which policy, when, by whom and with which outcome
func TestStory18_TheRedBookOfWestmarch(t *testing.T) {
	realm := New(t)
	proposed := realm.As("samwise").Proposes(CreateUser, PeerReview, `{"email":"bilbo@shire.test"}`)
	realm.As("frodo").Comments(CreateUser, proposed.Id(), "Bilbo deserves an account.")
	realm.Waits(30 * time.Second)
	realm.As("frodo").VotesFavour(CreateUser, proposed.Id())

	Expect(t, len(realm.Audited()) == 1, "audited %+v", realm.Audited())
	entry := realm.Audited()[0]
	Expect(t, entry.SubmissionId == proposed.Id() && entry.Permission == "iam.user.create" && entry.Policy.Id == "peer review", "identity: %+v", entry)
	Expect(t, entry.Url == CreateUser.Path && entry.Body == `{"email":"bilbo@shire.test"}` && entry.Decided-entry.Created == 30, "request and time: %+v", entry)
	Expect(t, entry.Resolution == "executed" && entry.Result != nil && entry.Result.Status == 204, "outcome: %+v", entry)
	Expect(t, slices.Equal(entry.Favour, []string{"samwise", "frodo"}) && entry.Comments == nil, "votes and comments: %+v", entry)
}
