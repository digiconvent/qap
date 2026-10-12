package layer3_test

import (
	"slices"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer3"
)

// the voter count is the member count of all groups of the policy
func TestVoterCount(t *testing.T) {
	var asked []string
	count := layer3.VoterCount(func(groupIds []string) int {
		asked = groupIds
		return 7
	})

	if got := count(concepts.Policy{Groups: []string{"a", "b"}}); got != 7 || !slices.Equal(asked, []string{"a", "b"}) {
		t.Fatalf("got %d for groups %v", got, asked)
	}
}
