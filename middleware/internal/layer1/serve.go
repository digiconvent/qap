package layer1

import (
	"strings"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer4"
	"github.com/gofiber/fiber/v3"
)

func (g *guard) serveLoggedIn(ctx fiber.Ctx, endpoint concepts.Handler) error {
	if g.dependencies.CurrentUser(ctx) == "" {
		return concepts.Unauthorized.Send(ctx)
	}
	return g.serve(ctx, endpoint)
}

func (g *guard) serve(ctx fiber.Ctx, endpoint concepts.Handler) error {
	request := layer4.CopyRequest(ctx)
	request.User = strings.Clone(g.dependencies.CurrentUser(ctx))
	if endpoint.Check != nil {
		if problems := layer4.Problems(endpoint.Check(request.Body)); problems != nil {
			return concepts.CannotBeCarriedOut(problems).Send(ctx)
		}
	}
	return layer4.WriteResult(ctx, endpoint.Execute(request))
}
