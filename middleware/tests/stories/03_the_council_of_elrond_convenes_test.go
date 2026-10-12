package stories_test

import (
	"slices"
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// samwise may create users under two policies, while aragorn of the white council may use neither
func TestStory03_TheCouncilOfElrondConvenes(t *testing.T) {
	realm := New(t)
	samwise := realm.As("samwise")
	aragorn := realm.As("aragorn")

	Expect(t, slices.Equal(samwise.Discovers(CreateUser).PolicyNames(), []string{"peer review", "fast track"}), "samwise should see peer review and fast track")
	Expect(t, len(aragorn.Discovers(CreateUser).PolicyNames()) == 0, "aragorn should see no policies")

	refused := aragorn.Proposes(CreateUser, PeerReview, "")
	Expect(t, refused.Status == 403 && refused.Message() == "Policy does not include you", "aragorn submit: %d %s", refused.Status, refused.Raw)

	withoutPolicy := samwise.ProposesWithoutPolicy(CreateUser)
	acceptable, _ := withoutPolicy.Body["acceptable_policies"].([]any)
	Expect(t, withoutPolicy.Status == 406 && len(acceptable) == 2, "samwise submit without policy: %d %s", withoutPolicy.Status, withoutPolicy.Raw)
}
