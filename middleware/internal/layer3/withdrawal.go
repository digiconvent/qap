package layer3

import (
	"time"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

type WithdrawalDependencies struct {
	AuditTrail concepts.AuditTrail
	Now        func() time.Time
}

func NewWithdrawal(d WithdrawalDependencies) func(s concepts.Submission) {
	return func(s concepts.Submission) {
		s.Resolution = concepts.Withdrawn
		d.AuditTrail.Record(layer4.SummarizeAudit(s, d.Now().Unix()))
	}
}
