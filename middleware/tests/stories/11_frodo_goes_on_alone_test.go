package stories_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// frodo takes back his proposal, so nothing waits for the fellowship anymore
func TestStory11_FrodoGoesOnAlone(t *testing.T) {
	realm := New(t)
	frodo := realm.As("frodo")
	proposed := frodo.Proposes(CreateUser, PeerReview, "")

	Expect(t, frodo.Withdraws(CreateUser, proposed.Id()).Status == 200, "frodo should withdraw his proposal")
	listed := realm.As("samwise").Lists(CreateUser)
	Expect(t, listed.Raw == "[]", "samwise list: %s", listed.Raw)
}
