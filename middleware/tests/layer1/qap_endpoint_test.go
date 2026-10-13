package layer1_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/digiconvent/qap/middleware/internal/concepts"
	"github.com/digiconvent/qap/middleware/internal/layer1"
	"github.com/digiconvent/qap/middleware/internal/layer2"
	"github.com/gofiber/fiber/v3"
)

const (
	book = "11111111-1111-4111-8111-111111111111"
	user = "22222222-2222-4222-8222-222222222222"
)

var intentNames = []string{"discover", "submit", "list", "poll", "vote", "comment", "delete-submission"}

func stubCurrentUser(ctx fiber.Ctx) string {
	return ctx.Get("X-User")
}

func stubIntent(name string) layer2.Handler {
	return func(ctx fiber.Ctx, call layer2.Call) error {
		endpoint := "without endpoint"
		if call.Endpoint.Execute != nil {
			endpoint = "with endpoint"
		}
		return ctx.Status(299).SendString(name + "|" + call.User + "|" + strings.Join(call.Permissions, ",") + "|" + endpoint)
	}
}

func newGuard() layer1.Guard {
	intents := map[string]layer2.Handler{}
	for _, name := range intentNames {
		intents[name] = stubIntent(name)
	}
	return layer1.New(layer1.Dependencies{CurrentUser: stubCurrentUser, Intents: intents})
}

func passing(refuses bool) concepts.Handler {
	return concepts.Handler{
		Check: func(request concepts.Request) map[string][]error {
			if refuses {
				return map[string][]error{"body": {errors.New("refused")}}
			}
			return nil
		},
		Execute: func(request concepts.Request) concepts.Result {
			return concepts.Result{Status: 200, Body: "passed as " + request.User}
		},
	}
}

func send(t *testing.T, protection layer1.Protection, route string, path string, user string, intentHeader string) (int, string) {
	t.Helper()
	return sendTo(t, protection.Run(passing(false)), route, path, user, intentHeader)
}

func sendTo(t *testing.T, handler fiber.Handler, route string, path string, user string, intentHeader string) (int, string) {
	t.Helper()
	app := fiber.New()
	app.Get(route, handler)
	req := httptest.NewRequest(http.MethodGet, path, nil)
	if user != "" {
		req.Header.Set("X-User", user)
	}
	if intentHeader != "" {
		req.Header.Set("X-QAP-Intent", intentHeader)
	}
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(body)
}

// anonymous runs the endpoint for every request, with or without a user and with or without an intent
func TestEndpoint_AnonymousLetsEveryoneThrough(t *testing.T) {
	guard := newGuard()
	for _, c := range [][2]string{{"", ""}, {"alice", ""}, {"", "vote"}} {
		if status, body := send(t, guard.AllowAnonymous(), "/news", "/news", c[0], c[1]); status != 200 || body != "passed as "+c[0] {
			t.Errorf("user %q intent %q: got %d %q", c[0], c[1], status, body)
		}
	}
}

// authenticated runs the endpoint for a logged-in user and ignores any intent
func TestEndpoint_AuthenticatedLetsLoggedInUsersThrough(t *testing.T) {
	guard := newGuard()
	for _, intentHeader := range []string{"", "vote"} {
		if status, body := send(t, guard.AllowAuthenticatedOrRequiresPermission(), "/news", "/news", "alice", intentHeader); status != 200 || body != "passed as alice" {
			t.Errorf("intent %q: got %d %q", intentHeader, status, body)
		}
	}
}

// authenticated refuses a request without a logged-in user
func TestEndpoint_AuthenticatedRefusesAnonymous(t *testing.T) {
	if status, _ := send(t, newGuard().AllowAuthenticatedOrRequiresPermission(), "/news", "/news", "", ""); status != 401 {
		t.Errorf("got %d, want 401", status)
	}
}

// authenticated or requires permission runs the endpoint for a logged-in user when the request carries no intent
func TestEndpoint_AuthenticatedOrRequiresPermissionWithoutIntentPasses(t *testing.T) {
	protection := newGuard().AuthenticatedOrRequiresPermission("iam.user.set-password")
	if status, body := send(t, protection, "/users/:id", "/users/"+user, "alice", ""); status != 200 || body != "passed as alice" {
		t.Errorf("got %d %q", status, body)
	}
	if status, _ := send(t, protection, "/users/:id", "/users/"+user, "", ""); status != 401 {
		t.Errorf("anonymous: got %d, want 401", status)
	}
}

// authenticated or requires permission hands a request with an intent to qap, together with the endpoint
func TestEndpoint_AuthenticatedOrRequiresPermissionWithIntentGoesToQap(t *testing.T) {
	protection := newGuard().AuthenticatedOrRequiresPermission("iam.user.set-password")
	if status, body := send(t, protection, "/users/:id", "/users/"+user, "bob", "discover"); status != 299 || body != "discover|bob|iam.user.set-password|with endpoint" {
		t.Errorf("got %d %q", status, body)
	}
}

// a plain request whose check fails is refused with the problems, and the endpoint does not execute it
func TestEndpoint_PlainRequestIsChecked(t *testing.T) {
	handler := newGuard().AllowAnonymous().Run(passing(true))
	if status, body := sendTo(t, handler, "/news", "/news", "alice", ""); status != 422 || !strings.Contains(body, `"problems":{"body":["refused"]}`) {
		t.Errorf("got %d %q", status, body)
	}
}

// a route parameter id that is not a uuid is refused, whatever the protection and the intent
func TestEndpoint_IdMustBeUuid(t *testing.T) {
	guard := newGuard()
	for _, protection := range []layer1.Protection{guard.AllowAnonymous(), guard.RequiresPermission("library.book.{id}.delete")} {
		for _, intentHeader := range []string{"", "discover"} {
			if status, body := send(t, protection, "/books/:id", "/books/123", "alice", intentHeader); status != 400 || !strings.Contains(body, "the id must be a UUID") {
				t.Errorf("intent %q: got %d %q", intentHeader, status, body)
			}
		}
	}
}

// requires permission refuses a request without an intent
func TestEndpoint_RequiresPermissionWithoutIntentIsRefused(t *testing.T) {
	if status, body := send(t, newGuard().RequiresPermission("library.book.{id}.delete"), "/books/:id", "/books/"+book, "alice", ""); status != 406 || !strings.Contains(body, "QAP intent header required") {
		t.Errorf("got %d %q", status, body)
	}
}

// requires permission refuses an intent from someone who is not logged in
func TestEndpoint_RequiresPermissionWithoutUserIsRefused(t *testing.T) {
	if status, _ := send(t, newGuard().RequiresPermission("library.book.{id}.delete"), "/books/:id", "/books/"+book, "", "discover"); status != 401 {
		t.Errorf("got %d, want 401", status)
	}
}

// requires permission refuses an intent it does not know
func TestEndpoint_UnknownIntentIsRefused(t *testing.T) {
	if status, body := send(t, newGuard().RequiresPermission("library.book.{id}.delete"), "/books/:id", "/books/"+book, "alice", "audit"); status != 400 || !strings.Contains(body, "Unknown intent: 'audit'") {
		t.Errorf("got %d %q", status, body)
	}
}

// a permission placeholder without a matching route parameter is a mistake of the host
func TestEndpoint_MisconfiguredRouteFails(t *testing.T) {
	if status, _ := send(t, newGuard().RequiresPermission("library.book.{id}.delete"), "/books", "/books", "alice", "discover"); status != 500 {
		t.Errorf("got %d, want 500", status)
	}
}

// every known intent reaches its own handler with the user and the endpoint, and only discover and submit with the resolved permissions
func TestEndpoint_EveryIntentReachesItsHandler(t *testing.T) {
	protection := newGuard().RequiresPermission("library.book.{id}.delete", "library.admin")
	for _, name := range intentNames {
		want := name + "|alice|"
		if name == "discover" || name == "submit" {
			want += "library.book." + book + ".delete,library.admin"
		}
		want += "|with endpoint"
		if status, body := send(t, protection, "/books/:id", "/books/"+book, "alice", name); status != 299 || body != want {
			t.Errorf("%s: got %d %q", name, status, body)
		}
	}
}

// list and the intents on an existing submission do not resolve the permission, so a misconfigured template does not refuse them
func TestEndpoint_IntentOnSubmissionIgnoresTheTemplate(t *testing.T) {
	protection := newGuard().RequiresPermission("library.book.{id}.delete")
	for _, name := range []string{"list", "poll", "vote", "comment", "delete-submission"} {
		if status, body := send(t, protection, "/books", "/books", "alice", name); status != 299 || body != name+"|alice||with endpoint" {
			t.Errorf("%s: got %d %q", name, status, body)
		}
	}
}

// authenticated or requires permission refuses an intent it does not know
func TestEndpoint_AuthenticatedOrRequiresPermissionUnknownIntentIsRefused(t *testing.T) {
	protection := newGuard().AuthenticatedOrRequiresPermission("iam.user.set-password")
	if status, body := send(t, protection, "/users/:id", "/users/"+user, "bob", "audit"); status != 400 || !strings.Contains(body, "Unknown intent: 'audit'") {
		t.Errorf("got %d %q", status, body)
	}
}

// a guard that names no permission is refused when the route is set up, not on its first request
func TestEndpoint_GuardWithoutPermissionIsRefused(t *testing.T) {
	guard := newGuard()
	for name, protect := range map[string]func(...string) layer1.Protection{
		"RequiresPermission":                guard.RequiresPermission,
		"AuthenticatedOrRequiresPermission": guard.AuthenticatedOrRequiresPermission,
	} {
		func() {
			defer func() {
				if recover() == nil {
					t.Fatalf("%s without permissions did not panic", name)
				}
			}()
			protect()
		}()
	}
}
