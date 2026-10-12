package concepts_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/gofiber/fiber/v3"
)

func sent(t *testing.T, e concepts.Error) (int, map[string]any) {
	t.Helper()
	app := fiber.New()
	app.Get("/", func(ctx fiber.Ctx) error { return e.Send(ctx) })
	resp, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	if err != nil {
		t.Fatal(err)
	}
	raw, _ := io.ReadAll(resp.Body)
	body := map[string]any{}
	if err := json.Unmarshal(raw, &body); err != nil {
		t.Fatalf("body is not json: %s", raw)
	}
	return resp.StatusCode, body
}

// every error is sent with its status and a json body that names the problem
func TestError_SendsStatusAndBody(t *testing.T) {
	cases := []struct {
		err     concepts.Error
		status  int
		message string
	}{
		{concepts.Unauthorized, 401, "unauthorized"},
		{concepts.MissingIntent, 406, "QAP intent header required"},
		{concepts.UnknownIntent("audit"), 400, "Unknown intent: 'audit'"},
		{concepts.MissingSubmissionId, 406, "QAP id header required"},
		{concepts.MissingPolicyId(nil), 406, "QAP policy id header required"},
		{concepts.PolicyDoesNotIncludeYou, 403, "Policy does not include you"},
		{concepts.MisconfiguredRoute, 500, "misconfigured permission template"},
		{concepts.UnknownSubmission, 404, "Submission state does not exist"},
		{concepts.AlreadyResolved, 409, "Submission is already resolved"},
		{concepts.NotEligible, 403, "You cannot vote on this"},
		{concepts.InvalidVote, 400, "X-QAP-Vote must be a whole number"},
		{concepts.WithholdNotPermitted, 400, "withhold is not permitted for this poll"},
		{concepts.NotTheAuthor, 403, "Only the author can delete this submission"},
		{concepts.MissingComment, 406, "QAP comment header required"},
		{concepts.InvalidComment, 400, "X-QAP-Comment must be percent-encoded"},
		{concepts.CommentTooLong, 400, "X-QAP-Comment must be at most 500 characters"},
	}
	for _, c := range cases {
		status, body := sent(t, c.err)
		if status != c.status || body["error"] != c.message {
			t.Errorf("got %d %v, want %d %q", status, body, c.status, c.message)
		}
	}
}

// a missing policy id lists the policies the user may pick, and says how to fix the request
func TestError_MissingPolicyIdListsAcceptablePolicies(t *testing.T) {
	_, body := sent(t, concepts.MissingPolicyId([]concepts.PolicyReference{{Id: "pol_librarians", Name: "two_librarians"}}))

	policies, _ := body["acceptable_policies"].([]any)
	if body["hint"] != "Specify X-QAP-Policy-ID" || len(policies) != 1 {
		t.Fatalf("unexpected body %v", body)
	}
	if first, _ := policies[0].(map[string]any); first["id"] != "pol_librarians" || first["name"] != "two_librarians" {
		t.Fatalf("unexpected policy %v", policies[0])
	}
}
