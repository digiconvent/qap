package rules_test

import (
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// listing on a route shows the pending submissions of every id of that route template, as far as their policy includes you
func TestList_ShowsEveryIdOfTheRoute(t *testing.T) {
	realm := New(t)
	realm.As("frodo").Proposes(UpdateFrodo, PeerReview, `{"name":"Mr. Underhill"}`)
	realm.As("frodo").Proposes(UpdateSamwise, PeerReview, `{"name":"Samwise the Brave"}`)

	for _, route := range []Route{UpdateFrodo, UpdateSamwise} {
		listed := realm.As("boromir").Lists(route)
		Expect(t, listed.Status == 200 && len(listed.List) == 2, "boromir on %s: got %d %s", route.Path, listed.Status, listed.Raw)
	}
	outsider := realm.As("gandalf").Lists(UpdateFrodo)
	Expect(t, outsider.Status == 200 && len(outsider.List) == 0, "gandalf: got %d %s", outsider.Status, outsider.Raw)
}

// another route under the same permission keeps its own list
func TestList_LeavesOutOtherRoutesWithTheSamePermission(t *testing.T) {
	realm := New(t)
	realm.As("frodo").Proposes(UpdateFrodo, PeerReview, `{"name":"Mr. Underhill"}`)

	renames := realm.As("boromir").Lists(RenameFrodo)
	Expect(t, renames.Status == 200 && len(renames.List) == 0, "got %d %s", renames.Status, renames.Raw)
}
