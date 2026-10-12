package layer1

import (
	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer2"
	"github.com/gofiber/fiber/v3"
)

type Guard interface {
	AllowAnonymous() Protection
	AllowAuthenticatedOrRequiresPermission() Protection
	AuthenticatedOrRequiresPermission(permissions ...string) Protection
	RequiresPermission(permissions ...string) Protection
}

type Protection interface {
	Run(endpoint concepts.Handler) fiber.Handler
}

type Dependencies struct {
	CurrentUser func(ctx fiber.Ctx) string
	Intents     map[string]layer2.Handler
}

type guard struct {
	dependencies Dependencies
}

func New(dependencies Dependencies) Guard {
	return &guard{dependencies: dependencies}
}
