package stories_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// the fellowship is created and gains a member, each change under peer review on its own route
func TestStory06_TheFellowshipOfTheRingIsFormed(t *testing.T) {
	realm := New(t)
	samwise := realm.As("samwise")
	frodo := realm.As("frodo")

	group := samwise.Proposes(CreateGroup, PeerReview, `{"name":"The Fellowship"}`)
	Expect(t, frodo.VotesFavour(CreateGroup, group.Id()).Resolution() == "executed", "group creation should be carried out")

	member := samwise.Proposes(AddUserToFellowship, PeerReview, `{"user":"boromir"}`)
	Expect(t, frodo.VotesFavour(AddUserToFellowship, member.Id()).Resolution() == "executed", "adding boromir should be carried out")

	Expect(t, len(realm.Executed()) == 2 && realm.Executed()[0].Url == CreateGroup.Path && realm.Executed()[1].Url == AddUserToFellowship.Path, "executions: %+v", realm.Executed())
}
