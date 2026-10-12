package layer3_test

import (
	"errors"
	"testing"
	"time"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer3"
)

type settlingStore struct {
	s            concepts.Submission
	resolveError error
	resultError  error
}

func (f *settlingStore) Resolve(id string, revision int, resolution string, at int64) (concepts.Submission, error) {
	if f.resolveError != nil {
		return concepts.Submission{}, f.resolveError
	}
	f.s.Resolution = resolution
	return f.s, nil
}

func (f *settlingStore) SetResult(id string, result concepts.Result) (concepts.Submission, error) {
	if f.resultError != nil {
		return concepts.Submission{}, f.resultError
	}
	f.s.Result = &result
	return f.s, nil
}

func (f *settlingStore) Get(id string) (concepts.Submission, bool) {
	return f.s, true
}

type settlingRecorder struct {
	executed []concepts.Submission
	audited  []concepts.AuditSummary
}

func (r *settlingRecorder) executor(s concepts.Submission) (concepts.Result, map[string]concepts.Change) {
	r.executed = append(r.executed, s)
	return concepts.Result{Status: 204}, map[string]concepts.Change{"title": {Before: "old", After: "new"}}
}

func (r *settlingRecorder) Record(summary concepts.AuditSummary) {
	r.audited = append(r.audited, summary)
}

func seventy() time.Time {
	return time.Unix(70, 0)
}

func pendingSettlement() *settlingStore {
	return &settlingStore{s: concepts.Submission{Id: "qap_1", Resolution: concepts.Pending}}
}

func execute(store *settlingStore, r *settlingRecorder) func(concepts.Submission) concepts.Submission {
	return layer3.NewExecution(layer3.ExecutionDependencies{Store: store, Executor: r.executor, AuditTrail: r, Now: seventy})
}

// execution resolves the submission, runs the executor, stores the result and writes the audit summary
func TestExecution(t *testing.T) {
	store, r := pendingSettlement(), &settlingRecorder{}
	got := execute(store, r)(store.s)

	if got.Resolution != concepts.Executed || got.Result == nil || got.Result.Status != 204 || len(r.executed) != 1 {
		t.Fatalf("got %+v, executed %d", got, len(r.executed))
	}
	if len(r.audited) != 1 || r.audited[0].Resolution != concepts.Executed || r.audited[0].Decided != 70 || r.audited[0].Changes["title"].After != "new" {
		t.Fatalf("audited %+v", r.audited)
	}
}

// an anonymous poll leaves the changes out of its audit summary
func TestExecution_AnonymousLeavesOutTheChanges(t *testing.T) {
	store, r := pendingSettlement(), &settlingRecorder{}
	store.s.Policy.Anonymous = true
	execute(store, r)(store.s)

	if len(r.audited) != 1 || r.audited[0].Changes != nil {
		t.Fatalf("audited %+v", r.audited)
	}
}

// a submission that another request already resolved is not carried out again
func TestExecution_LosingTheRaceDoesNothing(t *testing.T) {
	store := &settlingStore{s: concepts.Submission{Id: "qap_1", Resolution: concepts.Rejected}, resolveError: concepts.ErrAlreadyResolved}
	r := &settlingRecorder{}
	got := execute(store, r)(store.s)

	if got.Resolution != concepts.Rejected || len(r.executed) != 0 || len(r.audited) != 0 {
		t.Fatalf("got %+v, executed %d, audited %d", got, len(r.executed), len(r.audited))
	}
}

// a submission removed while it was carried out still reports its result
func TestExecution_ResultThatCannotBeStored(t *testing.T) {
	store, r := pendingSettlement(), &settlingRecorder{}
	store.resultError = errors.New("gone")
	got := execute(store, r)(store.s)

	if got.Resolution != concepts.Executed || got.Result == nil || len(r.executed) != 1 || len(r.audited) != 1 {
		t.Fatalf("got %+v, executed %d, audited %d", got, len(r.executed), len(r.audited))
	}
}

// an executor that panics leaves the submission executed with a 500 result, and it is still audited
func TestExecution_ExecutorThatPanics(t *testing.T) {
	store, r := pendingSettlement(), &settlingRecorder{}
	panicking := func(s concepts.Submission) (concepts.Result, map[string]concepts.Change) { panic("the host failed") }
	got := layer3.NewExecution(layer3.ExecutionDependencies{Store: store, Executor: panicking, AuditTrail: r, Now: seventy})(store.s)

	if got.Resolution != concepts.Executed || got.Result == nil || got.Result.Status != 500 || len(r.audited) != 1 {
		t.Fatalf("got %+v, audited %d", got, len(r.audited))
	}
}
