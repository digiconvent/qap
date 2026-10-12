package layer1

import (
	"slices"
	"strings"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer2"
	"github.com/digiconvent/qap/middleware/internal/layer4"
	"github.com/gofiber/fiber/v3"
)

const intentHeader = "X-QAP-Intent"

var permissionIntents = []string{"discover", "submit"}

func (g *guard) handle(ctx fiber.Ctx, intent string, handler layer2.Handler, permissions []string, endpoint concepts.Handler) error {
	user := strings.Clone(g.dependencies.CurrentUser(ctx))
	if user == "" {
		return concepts.Unauthorized.Send(ctx)
	}
	if !slices.Contains(permissionIntents, intent) {
		return handler(ctx, layer2.Call{User: user, Endpoint: endpoint})
	}

	param := func(key string) string { return ctx.Params(key) }
	resolved := make([]string, 0, len(permissions))
	for _, permission := range permissions {
		value, ok := layer4.ResolvePermission(permission, param)
		if !ok {
			return concepts.MisconfiguredRoute.Send(ctx)
		}
		resolved = append(resolved, strings.Clone(value))
	}

	return handler(ctx, layer2.Call{User: user, Permissions: resolved, Endpoint: endpoint})
}
