package host

import "github.com/digiconvent/qap/middleware/internal/concepts"

func (s *setup) RecordAuditTo(auditTrail concepts.AuditTrail) ReadyForMiddleware {
	s.auditTrail = auditTrail
	return s
}
