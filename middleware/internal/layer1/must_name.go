package layer1

func mustName(guard string, permissions []string) {
	if len(permissions) == 0 {
		panic("qap: " + guard + " needs at least one permission")
	}
}
