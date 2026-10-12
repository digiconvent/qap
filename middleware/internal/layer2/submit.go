package layer2

import (
	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/gofiber/fiber/v3"
)

func (i Intents) Submit(ctx fiber.Ctx, call Call) error {
	eligible := i.eligiblePolicies(call)
	policyId := ctx.Get(policyIdHeader)
	if policyId == "" {
		acceptable := []concepts.PolicyReference{}
		for _, e := range eligible {
			acceptable = append(acceptable, concepts.PolicyReference{Id: e.policy.Id, Name: e.policy.Name})
		}
		return concepts.MissingPolicyId(acceptable).Send(ctx)
	}

	var chosen *permittedPolicy
	for _, e := range eligible {
		if e.policy.Id == policyId {
			chosen = &e
			break
		}
	}
	if chosen == nil {
		return concepts.PolicyDoesNotIncludeYou.Send(ctx)
	}

	request := i.dependencies.CopyRequest(ctx)
	request.User = call.User
	if problems := i.dependencies.Problems(call.Endpoint.Check(request.Body)); problems != nil {
		return concepts.CannotBeCarriedOut(problems).Send(ctx)
	}
	now := i.dependencies.Now().Unix()
	s := concepts.Submission{
		Id:         i.dependencies.NewId(),
		Endpoint:   request.Url,
		Method:     request.Method,
		Headers:    request.Headers,
		Body:       request.Body,
		Route:      request.Route,
		Params:     request.Params,
		Handler:    call.Endpoint,
		Permission: chosen.permission,
		Policy:     chosen.policy,
		Author:     call.User,
		When:       now,
		Resolution: concepts.Pending,
		Favour:     []string{call.User},
		Against:    []string{},
		Withhold:   []string{},
		Comments:   []concepts.Comment{},
		LastPolled: now,
		VoterCount: i.dependencies.VoterCount(chosen.policy),
	}
	i.dependencies.Store.Create(s)

	s = i.settle(s)
	if s.Resolution == concepts.Pending && s.Policy.DeadlineSeconds != nil {
		i.dependencies.Deadlines.Start(s)
	}
	return ctx.Status(fiber.StatusAccepted).JSON(i.dependencies.View(s, call.User))
}
