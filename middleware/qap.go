package qap

import (
	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/host"
	"github.com/digiconvent/qap/middleware/internal/layer1"
)

type (
	Policy                = concepts.Policy
	PolicyReference       = concepts.PolicyReference
	Submission            = concepts.Submission
	Comment               = concepts.Comment
	Result                = concepts.Result
	Endpoint[Payload any] = concepts.TypedEndpoint[Payload]
	Request[Payload any]  = concepts.TypedRequest[Payload]
	Reader                = concepts.Reader
	Handler               = concepts.Handler
	Change                = concepts.Change
	AuditSummary          = concepts.AuditSummary
	AuditTrail            = concepts.AuditTrail
	Guard                 = layer1.Guard
	Protection            = layer1.Protection

	NeedsCurrentUser   = host.NeedsCurrentUser
	NeedsPolicyLookup  = host.NeedsPolicyLookup
	NeedsMembership    = host.NeedsMembership
	NeedsMemberCount   = host.NeedsMemberCount
	NeedsAuditTrail    = host.NeedsAuditTrail
	ReadyForMiddleware = host.ReadyForMiddleware
)

func New() NeedsCurrentUser {
	return host.New()
}

func Handle[Payload any](endpoint Endpoint[Payload]) Handler {
	return host.Handle(endpoint)
}
