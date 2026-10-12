package realm

import "slices"

func (r *Realm) isMemberOf(userId string, groupId string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return slices.Contains(r.memberships[userId], groupId)
}
