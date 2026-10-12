package host

import (
	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer1"
	"github.com/gofiber/fiber/v3"
)

type NeedsCurrentUser interface {
	ReadUserFrom(currentUser func(ctx fiber.Ctx) string) NeedsPolicyLookup
}

type NeedsPolicyLookup interface {
	LoadPoliciesFrom(policyLookup func(permission string) []concepts.Policy) NeedsMembership
}

type NeedsMembership interface {
	VerifyEligibilityWith(isMemberOf func(userId string, groupId string) bool) NeedsMemberCount
}

type NeedsMemberCount interface {
	ComputePercentagesWith(countMembersOf func(groupIds []string) int) NeedsAuditTrail
}

type NeedsAuditTrail interface {
	RecordAuditTo(auditTrail concepts.AuditTrail) ReadyForMiddleware
}

type ReadyForMiddleware interface {
	AsMiddleware() layer1.Guard
}
