package rules_test

import (
	"sync"
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// two deciding votes at the same moment carry out the change exactly once
func TestRace(t *testing.T) {
	realm := New(t)
	proposed := realm.As("wormtongue").Proposes(RemoveEomer, KingsCouncil, "")
	realm.As("eowyn").VotesFavour(RemoveEomer, proposed.Id())

	var wg sync.WaitGroup
	for _, name := range []string{"theoden", "eomer"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			realm.As(name).VotesFavour(RemoveEomer, proposed.Id())
		}()
	}
	wg.Wait()

	Expect(t, len(realm.Executed()) == 1 && len(realm.Audited()) == 1, "executions %d, audited %d", len(realm.Executed()), len(realm.Audited()))
}
