package layer2

import "github.com/gofiber/fiber/v3"

func (i Intents) List(ctx fiber.Ctx, call Call) error {
	views := []any{}
	for _, s := range i.dependencies.Store.ListPending(ctx.Route().Path, ctx.Method()) {
		if i.dependencies.Eligible(call.User, s.Policy) {
			views = append(views, i.dependencies.View(s, call.User))
		}
	}
	return ctx.JSON(views)
}
