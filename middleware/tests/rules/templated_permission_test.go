package rules_test

import (
	"slices"
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// a permission with an id gives one library its own policy, which another library does not have
func TestTemplatedPermission(t *testing.T) {
	realm := New(t)
	samwise := realm.As("samwise")

	Expect(t, slices.Equal(samwise.Discovers(ShireAddsBook).PolicyNames(), []string{"shire library"}), "the shire library offers its policy")
	Expect(t, len(samwise.Discovers(MordorAddsBook).PolicyNames()) == 0, "mordor offers nothing")

	samwise.Proposes(ShireAddsBook, ShireLibrary, `{"title":"There and Back Again"}`)
	Expect(t, len(realm.Audited()) == 1 && realm.Audited()[0].Permission == "library."+ShireId+".add-book", "audited %+v", realm.Audited())
}
