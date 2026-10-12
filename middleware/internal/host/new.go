package host

import (
	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/gofiber/fiber/v3"
)

type setup struct {
	currentUser    func(ctx fiber.Ctx) string
	policyLookup   func(permission string) []concepts.Policy
	isMemberOf     func(userId string, groupId string) bool
	countMembersOf func(groupIds []string) int
	auditTrail     concepts.AuditTrail
}

func New() NeedsCurrentUser {
	return &setup{}
}
