package layer3_test

import (
	"errors"
	"fmt"
	"slices"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"

	"github.com/digiconvent/qap/middleware/internal/layer3"
)

func newStore() *layer3.Store {
	return layer3.NewStore()
}

func pending(id string, route string) concepts.Submission {
	return concepts.Submission{Id: id, Route: route, Method: "POST", Resolution: concepts.Pending, Favour: []string{"alice"}}
}

// a created submission can be found again, and an unknown id cannot
func TestStore_CreateAndGet(t *testing.T) {
	s := newStore()
	s.Create(pending("qap_1", "p"))

	if got, ok := s.Get("qap_1"); !ok || got.Id != "qap_1" {
		t.Fatalf("got %+v %v", got, ok)
	}
	if _, ok := s.Get("qap_2"); ok {
		t.Fatal("expected an unknown id not to be found")
	}
}

// a vote is recorded on a pending submission, replacing an earlier vote of the same person
func TestStore_Vote(t *testing.T) {
	s := newStore()
	s.Create(pending("qap_1", "p"))

	s.Vote("qap_1", "bob", concepts.Favour)
	got, err := s.Vote("qap_1", "bob", concepts.Against)
	if err != nil || !slices.Equal(got.Favour, []string{"alice"}) || !slices.Equal(got.Against, []string{"bob"}) {
		t.Fatalf("got %+v %v", got, err)
	}
	if _, err := s.Vote("qap_2", "bob", concepts.Favour); !errors.Is(err, concepts.ErrUnknownSubmission) {
		t.Fatalf("unknown: got %v", err)
	}
	s.Resolve("qap_1", 2, concepts.Executed, 1)
	if _, err := s.Vote("qap_1", "carol", concepts.Favour); !errors.Is(err, concepts.ErrAlreadyResolved) {
		t.Fatalf("resolved: got %v", err)
	}
}

// a comment is added to any existing submission
func TestStore_Comment(t *testing.T) {
	s := newStore()
	s.Create(pending("qap_1", "p"))

	comment := concepts.Comment{Author: "bob", At: 5, Text: "hi"}
	if got, err := s.Comment("qap_1", comment); err != nil || len(got.Comments) != 1 || got.Comments[0] != comment {
		t.Fatalf("got %+v %v", got, err)
	}
	if _, err := s.Comment("qap_2", comment); !errors.Is(err, concepts.ErrUnknownSubmission) {
		t.Fatalf("unknown: got %v", err)
	}
}

// a submission resolves once, records when, and stays as it was decided
func TestStore_ResolveOnce(t *testing.T) {
	s := newStore()
	s.Create(pending("qap_1", "p"))

	got, err := s.Resolve("qap_1", 0, concepts.Rejected, 70)
	if err != nil || got.Resolution != concepts.Rejected || got.LastPolled != 70 {
		t.Fatalf("got %+v %v", got, err)
	}
	if _, err := s.Resolve("qap_1", 0, concepts.Executed, 80); !errors.Is(err, concepts.ErrAlreadyResolved) {
		t.Fatalf("second resolve: got %v", err)
	}
	if _, err := s.Resolve("qap_2", 0, concepts.Executed, 80); !errors.Is(err, concepts.ErrUnknownSubmission) {
		t.Fatalf("unknown: got %v", err)
	}
}

// when many deciding requests race, exactly one resolves the submission
func TestStore_ConcurrentResolveHasOneWinner(t *testing.T) {
	s := newStore()
	s.Create(pending("qap_1", "p"))

	var winners atomic.Int32
	var wg sync.WaitGroup
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := s.Resolve("qap_1", 0, concepts.Executed, 1); err == nil {
				winners.Add(1)
			}
		}()
	}
	wg.Wait()
	if winners.Load() != 1 {
		t.Fatalf("expected one winner, got %d", winners.Load())
	}
}

// the result is stored on an existing submission
func TestStore_SetResult(t *testing.T) {
	s := newStore()
	s.Create(pending("qap_1", "p"))

	result := concepts.Result{Status: 204}
	if got, err := s.SetResult("qap_1", result); err != nil || got.Result == nil || got.Result.Status != 204 {
		t.Fatalf("got %+v %v", got, err)
	}
	if _, err := s.SetResult("qap_2", result); !errors.Is(err, concepts.ErrUnknownSubmission) {
		t.Fatalf("unknown: got %v", err)
	}
}

// only a pending submission can be deleted
func TestStore_Delete(t *testing.T) {
	s := newStore()
	s.Create(pending("qap_1", "p"))
	s.Create(pending("qap_2", "p"))
	s.Resolve("qap_2", 0, concepts.Executed, 1)

	if err := s.Delete("qap_1"); err != nil {
		t.Fatalf("pending: got %v", err)
	}
	if _, ok := s.Get("qap_1"); ok {
		t.Fatal("expected the submission to be gone")
	}
	if err := s.Delete("qap_1"); !errors.Is(err, concepts.ErrUnknownSubmission) {
		t.Fatalf("unknown: got %v", err)
	}
	if err := s.Delete("qap_2"); !errors.Is(err, concepts.ErrAlreadyResolved) {
		t.Fatalf("resolved: got %v", err)
	}
}

// list pending returns the pending submissions of one route template and method, in creation order
func TestStore_ListPending(t *testing.T) {
	s := newStore()
	s.Create(pending("qap_1", "/api/library/:id/add-book"))
	s.Create(pending("qap_2", "/api/library/:id/remove-book"))
	s.Create(pending("qap_3", "/api/library/:id/add-book"))
	deleting := pending("qap_4", "/api/library/:id/add-book")
	deleting.Method = "DELETE"
	s.Create(deleting)
	s.Create(pending("qap_5", "/api/library/:id/add-book"))
	s.Resolve("qap_5", 0, concepts.Executed, 1)

	var ids []string
	for _, pending := range s.ListPending("/api/library/:id/add-book", "POST") {
		ids = append(ids, pending.Id)
	}
	if !slices.Equal(ids, []string{"qap_1", "qap_3"}) {
		t.Fatalf("got %v", ids)
	}
}

// touch records when a submission was last polled, and ignores unknown ids
func TestStore_Touch(t *testing.T) {
	s := newStore()
	s.Create(pending("qap_1", "p"))
	s.Touch("qap_1", 42)
	s.Touch("qap_2", 42)

	if got, _ := s.Get("qap_1"); got.LastPolled != 42 {
		t.Fatalf("got %d", got.LastPolled)
	}
}

// remove expired removes the submissions that expiry allows, and keeps the rest
func TestStore_RemoveExpired(t *testing.T) {
	s := newStore()
	s.Create(pending("qap_1", "p"))
	s.Create(pending("qap_2", "p"))
	s.Resolve("qap_2", 0, concepts.Executed, 0)
	s.RemoveExpired(60)

	if _, ok := s.Get("qap_1"); !ok {
		t.Fatal("expected the pending submission to stay")
	}
	if _, ok := s.Get("qap_2"); ok {
		t.Fatal("expected the expired submission to be removed")
	}
}

// values passed in or returned never share their lists with the stored submission
func TestStore_ValuesAreCopies(t *testing.T) {
	s := newStore()
	input := pending("qap_1", "p")
	input.Comments = []concepts.Comment{{Author: "bob"}}
	input.Result = &concepts.Result{Status: 1}
	s.Create(input)
	input.Favour[0] = "mallory"
	input.Comments[0].Author = "mallory"
	input.Result.Status = 2

	got, _ := s.Get("qap_1")
	got.Favour[0] = "mallory"
	again, _ := s.Get("qap_1")
	if again.Favour[0] != "alice" || again.Comments[0].Author != "bob" || again.Result.Status != 1 {
		t.Fatalf("stored submission changed through a copy: %+v", again)
	}
}

// many concurrent votes on one submission are all recorded
func TestStore_ConcurrentVotesAreAllRecorded(t *testing.T) {
	s := newStore()
	s.Create(pending("qap_1", "p"))

	var wg sync.WaitGroup
	for i := range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Vote("qap_1", fmt.Sprintf("u%d", i), concepts.Favour)
		}()
	}
	wg.Wait()
	if got, _ := s.Get("qap_1"); len(got.Favour) != 21 {
		t.Fatalf("expected 21 votes, got %d", len(got.Favour))
	}
}

// a submission that received a vote after it was evaluated is not resolved on that outdated evaluation
func TestStore_ResolveOnOutdatedEvaluation(t *testing.T) {
	s := newStore()
	s.Create(pending("qap_1", "p"))
	evaluated, _ := s.Get("qap_1")
	s.Vote("qap_1", "bob", concepts.Against)

	if _, err := s.Resolve("qap_1", evaluated.Revision, concepts.Executed, 1); !errors.Is(err, concepts.ErrOutdated) {
		t.Fatalf("got %v", err)
	}
	if got, _ := s.Get("qap_1"); got.Resolution != concepts.Pending || got.Revision != 1 {
		t.Fatalf("got %+v", got)
	}
}

// the store keeps its own copy of the policy, so the host changing its policy afterwards changes nothing
func TestStore_KeepsItsOwnCopyOfThePolicy(t *testing.T) {
	s := newStore()
	against := 1
	submission := pending("qap_1", "p")
	submission.Policy = concepts.Policy{Groups: []string{"council"}, AgainstThreshold: &against}
	s.Create(submission)
	submission.Policy.Groups[0] = "everyone"
	against = 99

	got, _ := s.Get("qap_1")
	if got.Policy.Groups[0] != "council" || *got.Policy.AgainstThreshold != 1 {
		t.Fatalf("got %+v", got.Policy)
	}
}

// the store keeps its own copy of the request headers, so changing them afterwards changes nothing
func TestStore_KeepsItsOwnCopyOfTheHeaders(t *testing.T) {
	s := newStore()
	submission := pending("qap_1", "p")
	submission.Headers = map[string][]string{"Content-Type": {"application/json"}}
	s.Create(submission)
	submission.Headers["Content-Type"][0] = "text/plain"
	submission.Headers["Accept"] = []string{"*/*"}

	got, _ := s.Get("qap_1")
	if len(got.Headers) != 1 || got.Headers["Content-Type"][0] != "application/json" {
		t.Fatalf("got %v", got.Headers)
	}
}
