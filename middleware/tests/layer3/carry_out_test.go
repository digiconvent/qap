package layer3_test

import (
	"errors"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer3"
)

type library struct {
	refuses  bool
	executed []concepts.Request
	titles   map[string]string
}

func (l *library) handler(reads bool) concepts.Handler {
	handler := concepts.Handler{
		Check: func(request concepts.Request) map[string][]error {
			if l.refuses {
				return map[string][]error{"body": {errors.New("no longer valid")}}
			}
			return nil
		},
		Execute: func(request concepts.Request) concepts.Result {
			l.executed = append(l.executed, request)
			l.titles[request.Params["id"]] = "The Red Book"
			return concepts.Result{Status: 204}
		},
	}
	if reads {
		handler.Read = func(id string) any { return map[string]any{"title": l.titles[id]} }
	}
	return handler
}

func newLibrary() *library {
	return &library{titles: map[string]string{"7": "There and Back Again"}}
}

// carrying out checks the stored request again and executes it as the author, with the submission as context
func TestCarryOut(t *testing.T) {
	l := newLibrary()
	s := concepts.Submission{Id: "qap_1", Method: "POST", Endpoint: "/books/7", Route: "/books/:id", Params: map[string]string{"id": "7"}, Body: "{}", Author: "alice", Handler: l.handler(false)}
	result, changes := layer3.CarryOut(s)

	if result.Status != 204 || changes != nil || len(l.executed) != 1 {
		t.Fatalf("got %+v %v, executed %d", result, changes, len(l.executed))
	}
	request := l.executed[0]
	if request.User != "alice" || request.Params["id"] != "7" || request.Url != "/books/7" || request.Submission == nil || request.Submission.Id != "qap_1" {
		t.Fatalf("got %+v", request)
	}
}

// an endpoint that reads its entity gives the changed fields, before and after
func TestCarryOut_ReadGivesTheChanges(t *testing.T) {
	l := newLibrary()
	_, changes := layer3.CarryOut(concepts.Submission{Params: map[string]string{"id": "7"}, Handler: l.handler(true)})

	if len(changes) != 1 || changes["title"] != (concepts.Change{Before: "There and Back Again", After: "The Red Book"}) {
		t.Fatalf("got %+v", changes)
	}
}

// without an id in the route, nothing is read
func TestCarryOut_WithoutIdNothingIsRead(t *testing.T) {
	l := newLibrary()
	result, changes := layer3.CarryOut(concepts.Submission{Handler: l.handler(true)})

	if result.Status != 204 || changes != nil {
		t.Fatalf("got %+v %v", result, changes)
	}
}

// a request that no longer passes its check is not executed, and the result lists the problems
func TestCarryOut_CheckFails(t *testing.T) {
	l := newLibrary()
	l.refuses = true
	result, changes := layer3.CarryOut(concepts.Submission{Params: map[string]string{"id": "7"}, Handler: l.handler(true)})

	body, _ := result.Body.(map[string]any)
	if result.Status != 422 || len(l.executed) != 0 || body["problems"] == nil || changes != nil {
		t.Fatalf("got %+v, executed %d", result, len(l.executed))
	}
}

// a handler without a check is executed as it is
func TestCarryOut_WithoutCheck(t *testing.T) {
	l := newLibrary()
	l.refuses = true
	handler := l.handler(false)
	handler.Check = nil
	result, _ := layer3.CarryOut(concepts.Submission{Params: map[string]string{"id": "7"}, Handler: handler})

	if result.Status != 204 || len(l.executed) != 1 {
		t.Fatalf("got %+v, executed %d", result, len(l.executed))
	}
}
