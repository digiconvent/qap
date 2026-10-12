package realm

func (r *Realm) countMembersOf(groupIds []string) int {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	count := 0
	for user := range r.memberships {
		if r.belongsToAny(user, groupIds) {
			count++
		}
	}
	return count
}
