package layer2

import (
	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/gofiber/fiber/v3"
)

type Call struct {
	User        string
	Permissions []string
	Endpoint    concepts.Handler
}

type Handler func(ctx fiber.Ctx, call Call) error
