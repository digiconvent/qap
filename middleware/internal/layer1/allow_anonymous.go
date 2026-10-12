package layer1

func (g *guard) AllowAnonymous() Protection {
	return protection{guard: g, access: anyone}
}
