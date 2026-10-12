package rules_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	qap "github.com/digiconvent/qap/middleware"
	. "github.com/digiconvent/qap/middleware/tests/harness"

	"github.com/gofiber/fiber/v3"
)

// the public setup chain wires the same assembly, so a host that only imports qap gets a working guard
func TestThroughQapNew(t *testing.T) {
	realm := New(t)
	guard := qap.New().
		ReadUserFrom(realm.CurrentUser).
		LoadPoliciesFrom(realm.PolicyLookup).
		VerifyEligibilityWith(realm.IsMemberOf).
		ComputePercentagesWith(realm.CountMembersOf).
		RecordAuditTo(realm).
		AsMiddleware()
	app := fiber.New()
	app.Post("/api/iam/user/create", guard.RequiresPermission("iam.user.create").Run(realm.RecordingEndpoint()))

	req := httptest.NewRequest(http.MethodPost, "/api/iam/user/create", nil)
	req.Header.Set("X-User", "samwise")
	req.Header.Set("X-QAP-Intent", "submit")
	req.Header.Set("X-QAP-Policy-ID", "fast track")
	resp, err := app.Test(req)

	Expect(t, err == nil && resp.StatusCode == 202 && len(realm.Executed()) == 1, "got %v %v, executions %d", resp, err, len(realm.Executed()))
}
