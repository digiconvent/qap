package realm

import (
	"maps"
	"slices"

	"github.com/digiconvent/qap/demo/server/story"
)

func copyStartingMemberships() map[string][]string {
	copied := maps.Clone(story.StartingMemberships)
	for user, groups := range copied {
		copied[user] = slices.Clone(groups)
	}
	return copied
}
