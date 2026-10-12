package layer4

import "github.com/digiconvent/qap/middleware/internal/concepts"

func SummarizeAudit(s concepts.Submission, now int64) concepts.AuditSummary {
	if s.Policy.Anonymous {
		return concepts.AuditSummary{Anonymous: true, Author: s.Author, Decided: now, Resolution: s.Resolution}
	}
	summary := concepts.AuditSummary{
		SubmissionId: s.Id,
		Permission:   s.Permission,
		Policy:       concepts.PolicyReference{Id: s.Policy.Id, Name: s.Policy.Name},
		Author:       s.Author,
		Created:      s.When,
		Decided:      now,
		Url:          s.Endpoint,
		Body:         s.Body,
		Resolution:   s.Resolution,
		Result:       s.Result,
		Favour:       s.Favour,
		Against:      s.Against,
		Withhold:     s.Withhold,
	}
	if s.Policy.CommentsInAudit {
		summary.Comments = s.Comments
	}
	return summary
}
