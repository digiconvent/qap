package rules_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// samwise asks to read the history of frodo, and once approved only samwise sees the data
func TestSensitiveRead(t *testing.T) {
	realm := New(t)
	proposed := realm.As("samwise").Proposes(FrodosHistory, PeerReview, "")
	realm.As("frodo").VotesFavour(FrodosHistory, proposed.Id())

	author := realm.As("samwise").Polls(FrodosHistory, proposed.Id())
	data, _ := author.Result()["body"].(map[string]any)
	Expect(t, author.Resolution() == "executed" && data["user"] == "frodo", "author: %s", author.Raw)

	voter := realm.As("frodo").Polls(FrodosHistory, proposed.Id())
	Expect(t, voter.Result()["status"] == float64(200) && voter.Result()["body"] == nil, "voter: %s", voter.Raw)
}
