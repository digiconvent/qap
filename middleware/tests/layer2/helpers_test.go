package layer2_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer2"
	"github.com/gofiber/fiber/v3"
)

type fakeStore struct {
	submissions map[string]concepts.Submission
	touched     map[string]int64
}

func newFakeStore() *fakeStore {
	return &fakeStore{submissions: map[string]concepts.Submission{}, touched: map[string]int64{}}
}

func (f *fakeStore) Create(s concepts.Submission) {
	f.submissions[s.Id] = s
}

func (f *fakeStore) Get(id string) (concepts.Submission, bool) {
	s, ok := f.submissions[id]
	return s, ok
}

func (f *fakeStore) Vote(id string, user string, vote string) (concepts.Submission, error) {
	s, ok := f.submissions[id]
	if !ok {
		return s, concepts.ErrUnknownSubmission
	}
	if s.Resolution != concepts.Pending {
		return s, concepts.ErrAlreadyResolved
	}
	switch vote {
	case concepts.Favour:
		s.Favour = append(s.Favour, user)
	case concepts.Against:
		s.Against = append(s.Against, user)
	default:
		s.Withhold = append(s.Withhold, user)
	}
	f.submissions[id] = s
	return s, nil
}

func (f *fakeStore) Comment(id string, comment concepts.Comment) (concepts.Submission, error) {
	s, ok := f.submissions[id]
	if !ok {
		return s, concepts.ErrUnknownSubmission
	}
	s.Comments = append(s.Comments, comment)
	f.submissions[id] = s
	return s, nil
}

func (f *fakeStore) Delete(id string) error {
	s, ok := f.submissions[id]
	if !ok {
		return concepts.ErrUnknownSubmission
	}
	if s.Resolution != concepts.Pending {
		return concepts.ErrAlreadyResolved
	}
	delete(f.submissions, id)
	return nil
}

func (f *fakeStore) ListPending(route string, method string) []concepts.Submission {
	var pending []concepts.Submission
	for _, s := range f.submissions {
		if s.Route == route && s.Method == method && s.Resolution == concepts.Pending {
			pending = append(pending, s)
		}
	}
	return pending
}

func (f *fakeStore) Touch(id string, at int64) {
	f.touched[id] = at
}

type fakeDeadlines struct {
	started   []string
	cancelled []string
}

func (f *fakeDeadlines) Start(s concepts.Submission) {
	f.started = append(f.started, s.Id)
}

func (f *fakeDeadlines) Cancel(id string) {
	f.cancelled = append(f.cancelled, id)
}

var twoLibrarians = concepts.Policy{Id: "pol_librarians", Name: "two_librarians", Groups: []string{"librarians"}, FavourThreshold: 2}

var librarians = map[string]bool{"alice": true, "bob": true}

func stubEligible(user string, policy concepts.Policy) bool {
	return librarians[user]
}

func stubPolicyLookup(permission string) []concepts.Policy {
	if permission == "library.book.123.delete" {
		return []concepts.Policy{twoLibrarians}
	}
	return nil
}

func stubView(s concepts.Submission, user string) any {
	return map[string]any{"id": s.Id, "resolution": s.Resolution, "favour": len(s.Favour), "against": len(s.Against), "withhold": len(s.Withhold), "comments": len(s.Comments), "viewer": user}
}

func stubReadVote(header string) (string, bool) {
	votes := map[string]string{"1": concepts.Favour, "0": concepts.Withhold, "-1": concepts.Against}
	vote, ok := votes[header]
	return vote, ok
}

func stubEvaluate(s concepts.Submission, voterCount int, now int64) string {
	if len(s.Against) > 0 {
		return concepts.Rejected
	}
	if len(s.Favour) >= s.Policy.FavourThreshold {
		return concepts.Executed
	}
	return concepts.Pending
}

type harness struct {
	store        *fakeStore
	deadlines    *fakeDeadlines
	settled      []string
	endpoint     stubEndpoint
	dependencies layer2.Dependencies
	intents      layer2.Intents
}

type stubEndpoint struct {
	problems map[string][]error
}

func (e stubEndpoint) handler() concepts.Handler {
	return concepts.Handler{
		Check:   func(request concepts.Request) map[string][]error { return e.problems },
		Execute: func(request concepts.Request) concepts.Result { return concepts.Result{Status: 204} },
	}
}

func stubProblems(errs map[string][]error) map[string][]string {
	if len(errs) == 0 {
		return nil
	}
	problems := map[string][]string{}
	for field, list := range errs {
		for _, err := range list {
			problems[field] = append(problems[field], err.Error())
		}
	}
	return problems
}

func newHarness() *harness {
	h := &harness{store: newFakeStore(), deadlines: &fakeDeadlines{}}
	h.dependencies = layer2.Dependencies{
		PolicyLookup: stubPolicyLookup,
		Eligible:     stubEligible,
		VoterCount:   func(policy concepts.Policy) int { return 2 },
		Store:        h.store,
		Deadlines:    h.deadlines,
		Now:          func() time.Time { return time.Unix(1760000000, 0) },
		NewId:        func() string { return "qap_1" },
		CopyRequest: func(ctx fiber.Ctx) concepts.Request {
			return concepts.Request{Method: ctx.Method(), Url: ctx.OriginalURL(), Route: ctx.Route().Path, Params: map[string]string{"id": ctx.Params("id")}, Body: string(ctx.Body())}
		},
		Problems:        stubProblems,
		View:            stubView,
		ReadVote:        stubReadVote,
		WithholdAllowed: func(policy concepts.Policy) bool { return policy.WithholdPermitted },
		IsAuthor:        func(s concepts.Submission, user string) bool { return s.Author == user },
		Evaluate:        stubEvaluate,
		Execute: func(s concepts.Submission) concepts.Submission {
			h.settled = append(h.settled, "executed:"+s.Id)
			s.Resolution = concepts.Executed
			return s
		},
		Reject: func(s concepts.Submission) concepts.Submission {
			h.settled = append(h.settled, "rejected:"+s.Id)
			s.Resolution = concepts.Rejected
			return s
		},
		RecordWithdrawal: func(s concepts.Submission) {},
	}
	h.intents = layer2.New(h.dependencies)
	return h
}

func withPolicy(h *harness, policy concepts.Policy) *harness {
	h.dependencies.PolicyLookup = func(permission string) []concepts.Policy { return []concepts.Policy{policy} }
	h.intents = layer2.New(h.dependencies)
	return h
}

type racingStore struct {
	*fakeStore
}

func (racingStore) Vote(id string, user string, vote string) (concepts.Submission, error) {
	return concepts.Submission{}, concepts.ErrAlreadyResolved
}

func (racingStore) Delete(id string) error {
	return concepts.ErrAlreadyResolved
}

func racingIntents(h *harness) layer2.Intents {
	dependencies := h.dependencies
	dependencies.Store = racingStore{h.store}
	return layer2.New(dependencies)
}

func (h *harness) pending(id string, author string, favour ...string) concepts.Submission {
	s := concepts.Submission{Id: id, Endpoint: "/books/123", Method: "DELETE", Route: "/books/:id", Permission: "library.book.123.delete", Policy: twoLibrarians, Author: author, When: 1760000000, Resolution: concepts.Pending, Favour: favour}
	h.store.Create(s)
	return s
}

type response struct {
	status  int
	headers http.Header
	raw     string
	body    map[string]any
	list    []any
}

func (h *harness) send(t *testing.T, name string, user string, headers map[string]string, body string) response {
	t.Helper()
	app := fiber.New()
	handler := h.intents.Handlers()[name]
	app.All("/books/:id", func(ctx fiber.Ctx) error {
		return handler(ctx, layer2.Call{User: user, Permissions: []string{"library.book.123.delete"}, Endpoint: h.endpoint.handler()})
	})
	req := httptest.NewRequest(http.MethodDelete, "/books/123", strings.NewReader(body))
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	raw, _ := io.ReadAll(resp.Body)
	r := response{status: resp.StatusCode, headers: resp.Header, raw: string(raw)}
	json.Unmarshal(raw, &r.body)
	json.Unmarshal(raw, &r.list)
	return r
}

func (r response) error() string {
	message, _ := r.body["error"].(string)
	return message
}

type vanishingStore struct {
	*fakeStore
}

func (vanishingStore) Vote(id string, user string, vote string) (concepts.Submission, error) {
	return concepts.Submission{}, concepts.ErrUnknownSubmission
}

func (vanishingStore) Comment(id string, comment concepts.Comment) (concepts.Submission, error) {
	return concepts.Submission{}, concepts.ErrUnknownSubmission
}

func (vanishingStore) Delete(id string) error {
	return concepts.ErrUnknownSubmission
}

func vanishingIntents(h *harness) layer2.Intents {
	dependencies := h.dependencies
	dependencies.Store = vanishingStore{h.store}
	return layer2.New(dependencies)
}
