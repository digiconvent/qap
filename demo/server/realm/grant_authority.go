package realm

import (
	"slices"

	"github.com/digiconvent/qap/demo/server/story"
)

func (r *Realm) grantAuthority(to string) {
	if slices.Contains(r.memberships[to], story.KingsAuthority) {
		return
	}
	r.memberships[to] = append(r.memberships[to], story.KingsAuthority)
}
