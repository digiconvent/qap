package layer1

func (g *guard) RequiresPermission(permissions ...string) Protection {
	mustName("RequiresPermission", permissions)
	return protection{guard: g, access: intentOnly, permissions: permissions}
}
