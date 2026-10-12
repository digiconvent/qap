package layer2

import (
	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/gofiber/fiber/v3"
)

func (i Intents) Poll(ctx fiber.Ctx, call Call) error {
	s, failure := i.find(ctx)
	if failure != nil {
		return failure.Send(ctx)
	}
	if !i.dependencies.Eligible(call.User, s.Policy) {
		return concepts.NotEligible.Send(ctx)
	}
	i.dependencies.Store.Touch(s.Id, i.dependencies.Now().Unix())
	return ctx.JSON(i.dependencies.View(s, call.User))
}
