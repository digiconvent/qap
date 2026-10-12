package layer4_test

import (
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

// only a uuid passes as id
func TestIsUuid(t *testing.T) {
	for id, want := range map[string]bool{
		"00000000-0000-4000-8000-000000000001":  true,
		"ABCDEF01-2345-4789-ABCD-EF0123456789":  true,
		"123":                                   false,
		"00000000-0000-4000-8000-00000000000":   false,
		"00000000-0000-4000-8000-0000000000001": false,
		"00000000-0000-4000-8000-00000000000g":  false,
	} {
		if layer4.IsUuid(id) != want {
			t.Errorf("%q: got %v", id, !want)
		}
	}
}

type title struct {
	Title string `json:"title"`
}

// the changes list every field that differs before and after, also fields that appear or disappear
func TestChanges(t *testing.T) {
	before := map[string]any{"title": "There and Back Again", "author": "Bilbo", "pages": 300}
	after := map[string]any{"title": "The Red Book", "author": "Bilbo", "editor": "Frodo"}
	changes := layer4.Changes(before, after)

	want := map[string]concepts.Change{
		"title":  {Before: "There and Back Again", After: "The Red Book"},
		"pages":  {Before: float64(300), After: nil},
		"editor": {Before: nil, After: "Frodo"},
	}
	if len(changes) != len(want) {
		t.Fatalf("got %+v", changes)
	}
	for field, change := range want {
		if changes[field] != change {
			t.Errorf("%s: got %+v, want %+v", field, changes[field], change)
		}
	}
}

// nothing changed gives no changes, an entity that disappears gives all its fields, and a value that is not an object is compared as a whole
func TestChanges_Edges(t *testing.T) {
	if changes := layer4.Changes(title{"Same"}, title{"Same"}); changes != nil {
		t.Fatalf("unchanged: got %+v", changes)
	}
	if changes := layer4.Changes(title{"Gone"}, nil); changes["title"] != (concepts.Change{Before: "Gone", After: nil}) {
		t.Fatalf("deleted: got %+v", changes)
	}
	if changes := layer4.Changes("old", "new"); changes[""] != (concepts.Change{Before: "old", After: "new"}) {
		t.Fatalf("not an object: got %+v", changes)
	}
	if changes := layer4.Changes(make(chan int), nil); changes != nil {
		t.Fatalf("not encodable: got %+v", changes)
	}
}
