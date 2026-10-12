package realm

import "slices"

func (r *Realm) belongsToAny(userId string, groupIds []string) bool {
	for _, group := range r.memberships[userId] {
		if slices.Contains(groupIds, group) {
			return true
		}
	}
	return false
}
