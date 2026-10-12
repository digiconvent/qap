package layer3

import (
	"time"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

type ExecutionStore interface {
	Resolve(id string, revision int, resolution string, at int64) (concepts.Submission, error)
	SetResult(id string, result concepts.Result) (concepts.Submission, error)
	Get(id string) (concepts.Submission, bool)
}

type ExecutionDependencies struct {
	Store      ExecutionStore
	Executor   func(s concepts.Submission) (concepts.Result, map[string]concepts.Change)
	AuditTrail concepts.AuditTrail
	Now        func() time.Time
}

func NewExecution(d ExecutionDependencies) func(s concepts.Submission) concepts.Submission {
	return func(s concepts.Submission) concepts.Submission {
		now := d.Now().Unix()
		resolved, err := d.Store.Resolve(s.Id, s.Revision, concepts.Executed, now)
		if err != nil {
			current, _ := d.Store.Get(s.Id)
			return current
		}
		result, changes := runExecutor(d.Executor, resolved)
		updated, err := d.Store.SetResult(resolved.Id, result)
		if err != nil {
			resolved.Result = &result
			updated = resolved
		}
		summary := layer4.SummarizeAudit(updated, now)
		if !updated.Policy.Anonymous {
			summary.Changes = changes
		}
		d.AuditTrail.Record(summary)
		return updated
	}
}
