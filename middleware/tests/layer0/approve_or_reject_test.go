package layer0_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// frodo explains himself and casts the second favour vote, which carries out the change exactly once
func TestApprove(t *testing.T) {
	realm := New(t)
	proposed := realm.As("samwise").Proposes(CreateUser, PeerReview, `{"email":"frodo@shire.test"}`)
	frodo := realm.As("frodo")

	commented := frodo.Comments(CreateUser, proposed.Id(), "Frodo is ready to leave the Shire.")
	comments, _ := commented.Poll()["comments"].([]any)
	Expect(t, commented.Status == 200 && len(comments) == 1, "comment: %d %s", commented.Status, commented.Raw)

	voted := frodo.VotesFavour(CreateUser, proposed.Id())
	Expect(t, voted.Status == 200 && voted.Resolution() == "executed" && voted.Result()["status"] == float64(204), "vote: %d %s", voted.Status, voted.Raw)
	Expect(t, len(realm.Executed()) == 1 && realm.Executed()[0].Payload["email"] == "frodo@shire.test", "expected one execution with the stored body, got %+v", realm.Executed())
	Expect(t, len(realm.Audited()) == 1 && realm.Audited()[0].Resolution == "executed", "expected one audit summary, got %+v", realm.Audited())

	polled := realm.As("samwise").Polls(CreateUser, proposed.Id())
	Expect(t, polled.Resolution() == "executed", "samwise should see the outcome, got %s", polled.Raw)
}

// legolas uses his veto against the proposal of aragorn, which rejects it and carries out nothing
func TestReject(t *testing.T) {
	realm := New(t)
	proposed := realm.As("aragorn").Proposes(AddPermission, VetoPower, "")

	voted := realm.As("legolas").VotesAgainst(AddPermission, proposed.Id())
	Expect(t, voted.Status == 200 && voted.Resolution() == "rejected", "vote: %d %s", voted.Status, voted.Raw)
	Expect(t, len(realm.Executed()) == 0, "a rejected proposal may not be carried out")
	Expect(t, len(realm.Audited()) == 1 && realm.Audited()[0].Resolution == "rejected", "expected one audit summary, got %+v", realm.Audited())
}
