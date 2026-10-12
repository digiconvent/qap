package layer4

import (
	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/gofiber/fiber/v3"
)

func WriteResult(ctx fiber.Ctx, result concepts.Result) error {
	for name, value := range result.Headers {
		ctx.Set(name, value)
	}
	ctx.Status(result.Status)
	switch body := result.Body.(type) {
	case nil:
		return nil
	case string:
		return ctx.SendString(body)
	case []byte:
		return ctx.Send(body)
	}
	return ctx.JSON(result.Body)
}
