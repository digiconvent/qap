package layer3

import (
	"time"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

type RejectionStore interface {
	Resolve(id string, revision int, resolution string, at int64) (concepts.Submission, error)
	Get(id string) (concepts.Submission, bool)
}

type RejectionDependencies struct {
	Store      RejectionStore
	AuditTrail concepts.AuditTrail
	Now        func() time.Time
}

func NewRejection(d RejectionDependencies) func(s concepts.Submission) concepts.Submission {
	return func(s concepts.Submission) concepts.Submission {
		now := d.Now().Unix()
		resolved, err := d.Store.Resolve(s.Id, s.Revision, concepts.Rejected, now)
		if err != nil {
			current, _ := d.Store.Get(s.Id)
			return current
		}
		d.AuditTrail.Record(layer4.SummarizeAudit(resolved, now))
		return resolved
	}
}
