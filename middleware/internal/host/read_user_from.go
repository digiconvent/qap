package host

import "github.com/gofiber/fiber/v3"

func (s *setup) ReadUserFrom(currentUser func(ctx fiber.Ctx) string) NeedsPolicyLookup {
	s.currentUser = currentUser
	return s
}
