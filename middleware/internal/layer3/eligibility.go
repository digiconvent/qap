package layer3

import "github.com/digiconvent/qap/middleware/internal/concepts"

func Eligibility(isMemberOf func(userId string, groupId string) bool) func(user string, policy concepts.Policy) bool {
	return func(user string, policy concepts.Policy) bool {
		for _, group := range policy.Groups {
			if isMemberOf(user, group) {
				return true
			}
		}
		return false
	}
}
