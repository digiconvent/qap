package layer2

import (
	"errors"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/gofiber/fiber/v3"
)

func (i Intents) Vote(ctx fiber.Ctx, call Call) error {
	s, failure := i.find(ctx)
	if failure != nil {
		return failure.Send(ctx)
	}
	if s.Resolution != concepts.Pending {
		return concepts.AlreadyResolved.Send(ctx)
	}
	if !i.dependencies.Eligible(call.User, s.Policy) {
		return concepts.NotEligible.Send(ctx)
	}
	vote, ok := i.dependencies.ReadVote(ctx.Get(voteHeader))
	if !ok {
		return concepts.InvalidVote.Send(ctx)
	}
	if vote == concepts.Withhold && !i.dependencies.WithholdAllowed(s.Policy) {
		return concepts.WithholdNotPermitted.Send(ctx)
	}

	s, err := i.dependencies.Store.Vote(s.Id, call.User, vote)
	if errors.Is(err, concepts.ErrUnknownSubmission) {
		return concepts.UnknownSubmission.Send(ctx)
	}
	if err != nil {
		return concepts.AlreadyResolved.Send(ctx)
	}
	return ctx.JSON(i.dependencies.View(i.settle(s), call.User))
}
