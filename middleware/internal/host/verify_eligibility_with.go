package host

func (s *setup) VerifyEligibilityWith(isMemberOf func(userId string, groupId string) bool) NeedsMemberCount {
	s.isMemberOf = isMemberOf
	return s
}
