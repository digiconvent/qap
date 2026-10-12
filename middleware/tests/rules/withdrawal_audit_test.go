package rules_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// a withdrawn submission leaves an audit summary with the resolution withdrawn
func TestWithdrawal_IsAudited(t *testing.T) {
	realm := New(t)
	proposed := realm.As("frodo").Proposes(CreateUser, PeerReview, `{"email":"bilbo@shire.test"}`)
	realm.As("frodo").Withdraws(CreateUser, proposed.Id())

	audited := realm.Audited()
	Expect(t, len(audited) == 1 && audited[0].Resolution == "withdrawn" && audited[0].Author == "frodo", "audited %+v", audited)
}
