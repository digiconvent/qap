package layer4

import (
	"strings"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/gofiber/fiber/v3"
)

const protocolHeaderPrefix = "x-qap-"

func CopyRequest(ctx fiber.Ctx) concepts.Request {
	return concepts.Request{
		Method:  strings.Clone(ctx.Method()),
		Url:     strings.Clone(ctx.OriginalURL()),
		Route:   strings.Clone(ctx.Route().Path),
		Params:  copyParams(ctx),
		Headers: copyHeaders(ctx.GetReqHeaders()),
		Body:    string(ctx.Body()),
	}
}

func copyParams(ctx fiber.Ctx) map[string]string {
	params := map[string]string{}
	for _, name := range ctx.Route().Params {
		params[strings.Clone(name)] = strings.Clone(ctx.Params(name))
	}
	return params
}

func copyHeaders(headers map[string][]string) map[string][]string {
	copied := make(map[string][]string, len(headers))
	for name, values := range headers {
		if strings.HasPrefix(strings.ToLower(name), protocolHeaderPrefix) {
			continue
		}
		owned := make([]string, len(values))
		for i, value := range values {
			owned[i] = strings.Clone(value)
		}
		copied[strings.Clone(name)] = owned
	}
	return copied
}
