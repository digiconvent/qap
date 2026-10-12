package host

func (s *setup) ComputePercentagesWith(countMembersOf func(groupIds []string) int) NeedsAuditTrail {
	s.countMembersOf = countMembersOf
	return s
}
