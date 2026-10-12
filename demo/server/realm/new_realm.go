package realm

import (
	"github.com/digiconvent/qap/demo/server/story"
	qap "github.com/digiconvent/qap/middleware"
)

func New() *Realm {
	r := &Realm{executed: []execution{}, audited: []qap.AuditSummary{}, memberships: copyStartingMemberships()}
	r.Guard = qap.New().
		ReadUserFrom(currentUser).
		LoadPoliciesFrom(story.PolicyLookup).
		VerifyEligibilityWith(r.isMemberOf).
		ComputePercentagesWith(r.countMembersOf).
		RecordAuditTo(r).
		AsMiddleware()
	return r
}
