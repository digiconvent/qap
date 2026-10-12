package rules_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// intents on one submission are refused without an id, for an unknown id, and from members outside the policy
func TestRefusals_FindingASubmission(t *testing.T) {
	realm := New(t)
	proposed := realm.As("samwise").Proposes(CreateUser, PeerReview, "")
	frodo := realm.As("frodo")
	aragorn := realm.As("aragorn")

	withoutId := frodo.Intent(CreateUser, "poll", nil, "")
	Expect(t, withoutId.Status == 406 && withoutId.Message() == "QAP id header required", "without id: %d %s", withoutId.Status, withoutId.Raw)

	for name, got := range map[string]Answer{
		"poll":     frodo.Polls(CreateUser, "qap_unknown"),
		"vote":     frodo.VotesFavour(CreateUser, "qap_unknown"),
		"comment":  frodo.Comments(CreateUser, "qap_unknown", "hi"),
		"withdraw": frodo.Withdraws(CreateUser, "qap_unknown"),
	} {
		Expect(t, got.Status == 404 && got.Message() == "Submission state does not exist", "%s unknown: %d %s", name, got.Status, got.Raw)
	}

	for name, got := range map[string]Answer{
		"poll":    aragorn.Polls(CreateUser, proposed.Id()),
		"vote":    aragorn.VotesFavour(CreateUser, proposed.Id()),
		"comment": aragorn.Comments(CreateUser, proposed.Id(), "hi"),
	} {
		Expect(t, got.Status == 403 && got.Message() == "You cannot vote on this", "%s by aragorn: %d %s", name, got.Status, got.Raw)
	}
}

// a vote must be a whole number
func TestRefusals_InvalidVote(t *testing.T) {
	realm := New(t)
	proposed := realm.As("samwise").Proposes(CreateUser, PeerReview, "")
	got := realm.As("frodo").Intent(CreateUser, "vote", map[string]string{"X-QAP-ID": proposed.Id(), "X-QAP-Vote": "yes"}, "")
	Expect(t, got.Status == 400 && got.Message() == "X-QAP-Vote must be a whole number", "got %d %s", got.Status, got.Raw)
}

// a decided submission accepts no more votes and cannot be withdrawn
func TestRefusals_DecidedSubmission(t *testing.T) {
	realm := New(t)
	proposed := realm.As("samwise").Proposes(CreateUser, FastTrack, "")

	voted := realm.As("frodo").VotesAgainst(CreateUser, proposed.Id())
	Expect(t, voted.Status == 409 && voted.Message() == "Submission is already resolved", "vote: %d %s", voted.Status, voted.Raw)

	withdrawn := realm.As("samwise").Withdraws(CreateUser, proposed.Id())
	Expect(t, withdrawn.Status == 409, "withdraw: %d %s", withdrawn.Status, withdrawn.Raw)
	Expect(t, len(realm.Executed()) == 1, "the executed submission stays carried out once")
}
