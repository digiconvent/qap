package host_test

import (
	"errors"
	"testing"

	qap "github.com/digiconvent/qap/middleware"
	"github.com/digiconvent/qap/middleware/internal/concepts"
)

type shelf struct {
	checked  []book
	executed []qap.Request[book]
}

func (s *shelf) Check(payload book) map[string][]error {
	s.checked = append(s.checked, payload)
	if payload.Title == "" {
		return map[string][]error{"title": {errors.New("is required")}}
	}
	return nil
}

func (s *shelf) Execute(request qap.Request[book]) qap.Result {
	s.executed = append(s.executed, request)
	return qap.Result{Status: 204}
}

type readingShelf struct {
	*shelf
}

func (readingShelf) Read(id string) any {
	return book{Title: "read " + id}
}

// handle decodes the body into the payload of the endpoint before its check, and refuses a body that does not decode
func TestHandle_ChecksTheDecodedPayload(t *testing.T) {
	s := &shelf{}
	handler := qap.Handle[book](s)

	if problems := handler.Check(`{"title":"The Red Book"}`); problems != nil || len(s.checked) != 1 || s.checked[0].Title != "The Red Book" {
		t.Fatalf("got %v, checked %+v", problems, s.checked)
	}
	if problems := handler.Check(`{"title":`); problems["body"] == nil || len(s.checked) != 1 {
		t.Fatalf("a body that does not decode: got %v, checked %d", problems, len(s.checked))
	}
	if problems := handler.Check(""); problems["title"] == nil {
		t.Fatalf("an empty body is the empty payload: got %v", problems)
	}
}

// handle executes the decoded payload together with the request and its submission
func TestHandle_ExecutesTheTypedRequest(t *testing.T) {
	s := &shelf{}
	submission := &concepts.Submission{Id: "qap_1"}
	handler := qap.Handle[book](s)
	handler.Execute(concepts.Request{Method: "POST", Url: "/books/7", Route: "/books/:id", Params: map[string]string{"id": "7"}, Body: `{"title":"The Red Book"}`, User: "alice", Submission: submission})

	request := s.executed[0]
	if request.Payload.Title != "The Red Book" || request.User != "alice" || request.Params["id"] != "7" || request.Submission != submission || request.Route != "/books/:id" {
		t.Fatalf("got %+v", request)
	}
}

// handle only reads when the endpoint is a reader
func TestHandle_ReadOnlyForReaders(t *testing.T) {
	if qap.Handle[book](&shelf{}).Read != nil {
		t.Fatal("an endpoint without Read should not read")
	}
	read := qap.Handle[book](readingShelf{&shelf{}}).Read
	if read == nil || read("7") != (book{Title: "read 7"}) {
		t.Fatal("a reader should read")
	}
}
