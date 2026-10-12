package layer4

import "github.com/digiconvent/qap/middleware/internal/concepts"

func WithholdAllowed(policy concepts.Policy) bool {
	return policy.WithholdPermitted
}
