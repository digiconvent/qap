package stories_test

import (
	"slices"
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// one permission offers two policies, aragorn picks the secret ballot, and the response only ever shows vote counts
func TestStory07_TheShireMeanwhileTheMayorOfMichelDelving(t *testing.T) {
	realm := New(t)
	aragorn := realm.As("aragorn")

	Expect(t, slices.Equal(aragorn.Discovers(DeletePolicy).PolicyNames(), []string{"senior approval", "secret ballot"}), "aragorn should see both policies")

	proposed := aragorn.Proposes(DeletePolicy, SecretBallot, "")
	voted := realm.As("legolas").VotesFavour(DeletePolicy, proposed.Id())
	Expect(t, voted.Resolution() == "executed" && voted.Votes("favour") == 2, "vote: %s", voted.Raw)

	_, hasVoters := voted.Poll()["voters"]
	Expect(t, !hasVoters, "responses show counts, never who voted: %s", voted.Raw)
}
