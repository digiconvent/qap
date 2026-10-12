package layer3_test

import (
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer3"
)

// a withdrawal writes an audit summary with the resolution withdrawn and the time of the withdrawal
func TestWithdrawal(t *testing.T) {
	r := &settlingRecorder{}
	record := layer3.NewWithdrawal(layer3.WithdrawalDependencies{AuditTrail: r, Now: seventy})
	record(concepts.Submission{Id: "qap_1", Author: "alice", Resolution: concepts.Pending})

	if len(r.audited) != 1 || r.audited[0].Resolution != concepts.Withdrawn || r.audited[0].Decided != 70 || r.audited[0].Author != "alice" {
		t.Fatalf("audited %+v", r.audited)
	}
}
