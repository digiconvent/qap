package layer2

import "github.com/gofiber/fiber/v3"

type policySettings struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	FavourThreshold   int    `json:"favour_threshold"`
	AgainstThreshold  *int   `json:"against_threshold"`
	Quorum            *int   `json:"quorum"`
	DeadlineSeconds   *int   `json:"deadline_seconds"`
	ExtensionSeconds  *int   `json:"extension_seconds"`
	WithholdPermitted bool   `json:"withhold_permitted"`
}

func (i Intents) Discover(ctx fiber.Ctx, call Call) error {
	policies := []policySettings{}
	for _, eligible := range i.eligiblePolicies(call) {
		p := eligible.policy
		policies = append(policies, policySettings{p.Id, p.Name, p.FavourThreshold, p.AgainstThreshold, p.Quorum, p.DeadlineSeconds, p.ExtensionSeconds, p.WithholdPermitted})
	}
	ctx.Set(versionHeader, i.dependencies.Version)
	return ctx.JSON(fiber.Map{"endpoint": ctx.Path(), "permission": call.Permissions[0], "policies": policies})
}
