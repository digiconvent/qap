package layer2

import (
	"errors"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/gofiber/fiber/v3"
)

func (i Intents) Withdraw(ctx fiber.Ctx, call Call) error {
	s, failure := i.find(ctx)
	if failure != nil {
		return failure.Send(ctx)
	}
	if !i.dependencies.IsAuthor(s, call.User) {
		return concepts.NotTheAuthor.Send(ctx)
	}
	if s.Resolution != concepts.Pending {
		return concepts.AlreadyResolved.Send(ctx)
	}

	err := i.dependencies.Store.Delete(s.Id)
	if errors.Is(err, concepts.ErrUnknownSubmission) {
		return concepts.UnknownSubmission.Send(ctx)
	}
	if err != nil {
		return concepts.AlreadyResolved.Send(ctx)
	}
	i.dependencies.Deadlines.Cancel(s.Id)
	i.dependencies.RecordWithdrawal(s)
	return ctx.Status(fiber.StatusOK).Send(nil)
}
