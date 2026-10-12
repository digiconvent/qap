package realm

import "github.com/gofiber/fiber/v3"

func currentUser(ctx fiber.Ctx) string {
	return ctx.Get("X-User")
}
