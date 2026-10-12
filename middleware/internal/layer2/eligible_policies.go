package layer2

import "github.com/digiconvent/qap/middleware/internal/concepts"

type permittedPolicy struct {
	permission string
	policy     concepts.Policy
}

func (i Intents) eligiblePolicies(call Call) []permittedPolicy {
	var eligible []permittedPolicy
	for _, permission := range call.Permissions {
		for _, policy := range i.dependencies.PolicyLookup(permission) {
			if i.dependencies.Eligible(call.User, policy) {
				eligible = append(eligible, permittedPolicy{permission: permission, policy: policy})
			}
		}
	}
	return eligible
}
