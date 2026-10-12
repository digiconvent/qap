package layer3_test

import (
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer3"
)

// a user is eligible when they are a member of at least one group of the policy, and the check stops at the first yes
func TestEligibility(t *testing.T) {
	var asked []string
	isMemberOf := func(userId string, groupId string) bool {
		asked = append(asked, groupId)
		return userId == "alice" && groupId == "editors"
	}
	eligible := layer3.Eligibility(isMemberOf)
	policy := concepts.Policy{Groups: []string{"readers", "editors", "admins"}}

	if !eligible("alice", policy) || len(asked) != 2 {
		t.Fatalf("alice: expected eligible after asking about 2 groups, asked %v", asked)
	}
	if eligible("bob", policy) {
		t.Fatal("bob: expected not eligible")
	}
}
