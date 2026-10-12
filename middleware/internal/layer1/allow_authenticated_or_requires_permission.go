package layer1

func (g *guard) AllowAuthenticatedOrRequiresPermission() Protection {
	return protection{guard: g, access: loggedIn}
}
