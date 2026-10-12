package host

import (
	"time"

	"github.com/digiconvent/qap/middleware/internal/layer1"
)

func (s *setup) AsMiddleware() layer1.Guard {
	supplied := Host{
		CurrentUser:    s.currentUser,
		PolicyLookup:   s.policyLookup,
		IsMemberOf:     s.isMemberOf,
		CountMembersOf: s.countMembersOf,
		AuditTrail:     s.auditTrail,
	}
	clock := Clock{Now: time.Now, AfterFunc: time.AfterFunc, Ticks: time.NewTicker(time.Minute).C}
	return Assemble(supplied, clock)
}
