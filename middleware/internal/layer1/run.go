package layer1

import (
	"slices"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
	"github.com/gofiber/fiber/v3"
)

func (p protection) Run(endpoint concepts.Handler) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		if slices.Contains(ctx.Route().Params, "id") && !layer4.IsUuid(ctx.Params("id")) {
			return concepts.InvalidId.Send(ctx)
		}
		switch p.access {
		case anyone:
			return p.guard.serve(ctx, endpoint)
		case loggedIn:
			return p.guard.serveLoggedIn(ctx, endpoint)
		}
		header := ctx.Get(intentHeader)
		handler, outcome := layer4.Intent(header, p.guard.dependencies.Intents)
		switch {
		case outcome == layer4.MissingIntent && p.access == intentOnly:
			return concepts.MissingIntent.Send(ctx)
		case outcome == layer4.MissingIntent:
			return p.guard.serveLoggedIn(ctx, endpoint)
		case outcome == layer4.UnknownIntent:
			return concepts.UnknownIntent(header).Send(ctx)
		}
		return p.guard.handle(ctx, header, handler, p.permissions, endpoint)
	}
}
