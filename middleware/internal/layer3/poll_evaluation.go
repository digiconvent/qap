package layer3

import (
	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
)

func reaches(votes int, threshold int, voterCount int) bool {
	needed, satisfiable := layer4.RequiredCount(threshold, voterCount)
	return satisfiable && votes >= needed
}

func EvaluatePoll(s concepts.Submission, voterCount int, now int64) string {
	counts := layer4.CountVotes(s)
	policy := s.Policy
	vetoes := policy.AgainstThreshold != nil && *policy.AgainstThreshold != 0
	if vetoes && reaches(counts.Against, *policy.AgainstThreshold, voterCount) {
		return concepts.Rejected
	}
	quorumMet := policy.Quorum == nil || reaches(counts.Participation, *policy.Quorum, voterCount)
	if quorumMet && counts.Favour >= layer4.RequiredFavourCount(policy.FavourThreshold, voterCount) {
		return concepts.Executed
	}
	if !layer4.DeadlinePassed(s.When, policy.DeadlineSeconds, now) {
		return concepts.Pending
	}
	if policy.ExtensionSeconds == nil || counts.Against > 0 {
		return concepts.Rejected
	}
	if layer4.ExtensionPassed(s.When, policy.DeadlineSeconds, policy.ExtensionSeconds, now) {
		return concepts.Executed
	}
	return concepts.Pending
}
