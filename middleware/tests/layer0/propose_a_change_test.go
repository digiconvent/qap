package layer0_test

import (
	"slices"
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// samwise discovers the policies for creating a user, proposes it under peer review and polls it while it waits
func TestProposeAChange(t *testing.T) {
	realm := New(t)
	samwise := realm.As("samwise")

	discovered := samwise.Discovers(CreateUser)
	Expect(t, discovered.Status == 200 && slices.Equal(discovered.PolicyNames(), []string{"peer review", "fast track"}), "discover: %d %s", discovered.Status, discovered.Raw)

	proposed := samwise.Proposes(CreateUser, PeerReview, `{"email":"frodo@shire.test"}`)
	Expect(t, proposed.Status == 202 && proposed.Resolution() == "pending" && proposed.Votes("favour") == 1, "submit: %d %s", proposed.Status, proposed.Raw)

	polled := samwise.Polls(CreateUser, proposed.Id())
	Expect(t, polled.Status == 200 && polled.Resolution() == "pending" && polled.Submission()["author"] == "samwise", "poll: %d %s", polled.Status, polled.Raw)
	Expect(t, len(realm.Executed()) == 0, "nothing may be carried out before the second vote")
}
