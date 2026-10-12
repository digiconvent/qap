package story

import qap "github.com/digiconvent/qap/middleware"

func PolicyLookup(permission string) []qap.Policy {
	return policies[permission]
}
