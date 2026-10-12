package layer1

func (g *guard) AuthenticatedOrRequiresPermission(permissions ...string) Protection {
	mustName("AuthenticatedOrRequiresPermission", permissions)
	return protection{guard: g, access: loggedInOrIntent, permissions: permissions}
}
