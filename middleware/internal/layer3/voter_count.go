package layer3

import "github.com/digiconvent/qap/middleware/internal/concepts"

func VoterCount(countMembersOf func(groupIds []string) int) func(policy concepts.Policy) int {
	return func(policy concepts.Policy) int {
		return countMembersOf(policy.Groups)
	}
}
