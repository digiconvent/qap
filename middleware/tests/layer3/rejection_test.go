package layer3_test

import (
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer3"
)

func reject(store *settlingStore, r *settlingRecorder) func(concepts.Submission) concepts.Submission {
	return layer3.NewRejection(layer3.RejectionDependencies{Store: store, AuditTrail: r, Now: seventy})
}

// rejection resolves the submission as rejected and writes the audit summary
func TestRejection(t *testing.T) {
	store, r := pendingSettlement(), &settlingRecorder{}
	got := reject(store, r)(store.s)

	if got.Resolution != concepts.Rejected || len(r.executed) != 0 || len(r.audited) != 1 || r.audited[0].Resolution != concepts.Rejected {
		t.Fatalf("got %+v, audited %+v", got, r.audited)
	}
}

// a submission that another request already resolved is not rejected again
func TestRejection_LosingTheRaceDoesNothing(t *testing.T) {
	store := &settlingStore{s: concepts.Submission{Id: "qap_1", Resolution: concepts.Executed}, resolveError: concepts.ErrAlreadyResolved}
	r := &settlingRecorder{}
	got := reject(store, r)(store.s)

	if got.Resolution != concepts.Executed || len(r.audited) != 0 {
		t.Fatalf("got %+v, audited %+v", got, r.audited)
	}
}
