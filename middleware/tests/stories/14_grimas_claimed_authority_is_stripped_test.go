package stories_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// grima cannot use the veto policy himself, and legolas vetoes the grant of authority with a comment that lands in the audit
func TestStory14_GrimasClaimedAuthorityIsStripped(t *testing.T) {
	realm := New(t)
	refused := realm.As("wormtongue").Proposes(AddPermission, VetoPower, "")
	Expect(t, refused.Status == 403, "grima submit: %d %s", refused.Status, refused.Raw)

	proposed := realm.As("aragorn").Proposes(AddPermission, VetoPower, `{"permission":"iam.group.remove-user"}`)
	legolas := realm.As("legolas")
	legolas.Comments(AddPermission, proposed.Id(), "Wormtongue speaks for Saruman.")
	vetoed := legolas.VotesAgainst(AddPermission, proposed.Id())

	Expect(t, vetoed.Resolution() == "rejected" && len(realm.Executed()) == 0, "veto: %s", vetoed.Raw)
	Expect(t, len(realm.Audited()) == 1 && len(realm.Audited()[0].Comments) == 1 && realm.Audited()[0].Comments[0].Text == "Wormtongue speaks for Saruman.", "audited %+v", realm.Audited())
}
