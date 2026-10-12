package host

import "github.com/digiconvent/qap/middleware/internal/concepts"

func (s *setup) LoadPoliciesFrom(policyLookup func(permission string) []concepts.Policy) NeedsMembership {
	s.policyLookup = policyLookup
	return s
}
