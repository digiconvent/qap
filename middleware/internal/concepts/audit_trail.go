package concepts

type AuditTrail interface {
	Record(summary AuditSummary)
}
