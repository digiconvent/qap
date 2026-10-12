package stories_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// samwise vouches for frodo under fast track, and his one vote carries out the creation straight away
func TestStory01_BilboLeavesBagEndToFrodo(t *testing.T) {
	realm := New(t)
	proposed := realm.As("samwise").Proposes(CreateUser, FastTrack, `{"email":"frodo@shire.test"}`)

	Expect(t, proposed.Status == 202 && proposed.Resolution() == "executed" && proposed.Result()["status"] == float64(204), "submit: %d %s", proposed.Status, proposed.Raw)
	Expect(t, len(realm.Executed()) == 1 && realm.Executed()[0].Payload["email"] == "frodo@shire.test", "expected one execution, got %+v", realm.Executed())
}
