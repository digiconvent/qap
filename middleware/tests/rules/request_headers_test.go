package rules_test

import (
	"slices"
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// the executor receives the headers of the proposed request, without the headers of the protocol itself
func TestRequestHeaders_ReachTheExecutor(t *testing.T) {
	realm := New(t)
	realm.As("frodo").Intent(CreateUser, "submit", map[string]string{"X-QAP-Policy-ID": FastTrack.Id, "Content-Type": "application/json"}, `{"email":"bilbo@shire.test"}`)

	executed := realm.Executed()
	Expect(t, len(executed) == 1, "executed %d", len(executed))
	headers := executed[0].Headers
	Expect(t, slices.Equal(headers["Content-Type"], []string{"application/json"}), "headers %v", headers)
	for name := range headers {
		Expect(t, name != "X-Qap-Intent" && name != "X-Qap-Policy-Id", "the protocol header %s reached the executor", name)
	}
}
