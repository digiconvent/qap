package host_test

import (
	"reflect"
	"testing"

	qap "github.com/digiconvent/qap/middleware"
	"github.com/gofiber/fiber/v3"
)

func stubCurrentUser(ctx fiber.Ctx) string {
	return ""
}

func stubPolicyLookup(permission string) []qap.Policy {
	return nil
}

func stubIsMemberOf(userId string, groupId string) bool {
	return false
}

func stubCountMembersOf(groupIds []string) int {
	return 0
}

type stubAuditTrail struct{}

func (stubAuditTrail) Record(summary qap.AuditSummary) {}

var stubEndpoint = qap.Handler{Execute: func(request qap.Request) qap.Result {
	return qap.Result{Status: 204}
}}

func configuredGuard() qap.Guard {
	return qap.New().
		ReadUserFrom(stubCurrentUser).
		LoadPoliciesFrom(stubPolicyLookup).
		VerifyEligibilityWith(stubIsMemberOf).
		ComputePercentagesWith(stubCountMembersOf).
		RecordAuditTo(stubAuditTrail{}).
		AsMiddleware()
}

// the setup chain takes the five host functions in order and produces a guard
func TestSetup_ChainProducesGuard(t *testing.T) {
	if configuredGuard() == nil {
		t.Fatal("expected a guard from a fully configured setup chain")
	}
}

// each setup step offers exactly one next step, so a missing or misplaced step does not compile
func TestSetup_EachStepOffersOnlyTheNextStep(t *testing.T) {
	steps := []struct {
		step reflect.Type
		next string
	}{
		{reflect.TypeFor[qap.NeedsCurrentUser](), "ReadUserFrom"},
		{reflect.TypeFor[qap.NeedsPolicyLookup](), "LoadPoliciesFrom"},
		{reflect.TypeFor[qap.NeedsMembership](), "VerifyEligibilityWith"},
		{reflect.TypeFor[qap.NeedsMemberCount](), "ComputePercentagesWith"},
		{reflect.TypeFor[qap.NeedsAuditTrail](), "RecordAuditTo"},
		{reflect.TypeFor[qap.ReadyForMiddleware](), "AsMiddleware"},
	}

	for _, s := range steps {
		if s.step.NumMethod() != 1 || s.step.Method(0).Name != s.next {
			t.Errorf("%s should offer only %s", s.step.Name(), s.next)
		}
	}
}

// the guard offers the four kinds of protection, and each runs an endpoint as a fiber handler
func TestSetup_GuardOffersFourProtections(t *testing.T) {
	guard := configuredGuard()

	protections := map[string]fiber.Handler{
		"anonymous":                            guard.AllowAnonymous().Run(stubEndpoint),
		"authenticated":                        guard.AllowAuthenticatedOrRequiresPermission().Run(stubEndpoint),
		"authenticated or requires permission": guard.AuthenticatedOrRequiresPermission("iam.user.set-password").Run(stubEndpoint),
		"requires permission":                  guard.RequiresPermission("library.book.{id}.delete").Run(stubEndpoint),
	}

	for name, handler := range protections {
		if handler == nil {
			t.Errorf("expected a handler for %s", name)
		}
	}
}

// qap.New starts the setup chain, so a host only needs to import the qap package
func TestNew_StartsSetupChain(t *testing.T) {
	if qap.New() == nil {
		t.Fatal("expected the first step of the setup chain")
	}
}
