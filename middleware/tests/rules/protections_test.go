package rules_test

import (
	"slices"
	"testing"

	. "github.com/digiconvent/qap/middleware/tests/harness"
)

// anonymous lets everyone see the public events, logged in or not
func TestProtection_Anonymous(t *testing.T) {
	realm := New(t)
	for _, name := range []string{"", "samwise"} {
		got := realm.As(name).Request(PublicEvents, nil, "")
		Expect(t, got.Status == 200 && got.Raw == "reached", "%q: %d %s", name, got.Status, got.Raw)
	}
}

// authenticated lets every logged-in member read the news and refuses everyone else
func TestProtection_Authenticated(t *testing.T) {
	realm := New(t)
	Expect(t, realm.As("samwise").Request(LatestNews, nil, "").Raw == "reached", "samwise should read the news")
	Expect(t, realm.As("").Request(LatestNews, nil, "").Status == 401, "nobody logged in should be refused")
}

// frodo changes his own password with a plain request, while a request with an intent goes through qap
func TestProtection_AuthenticatedOrRequiresPermission(t *testing.T) {
	realm := New(t)
	Expect(t, realm.As("frodo").Request(SetFrodosPassword, nil, "").Raw == "reached", "frodo should reach the handler of the host")
	Expect(t, realm.As("").Request(SetFrodosPassword, nil, "").Status == 401, "nobody logged in should be refused")

	discovered := realm.As("samwise").Discovers(SetFrodosPassword)
	Expect(t, discovered.Status == 200 && slices.Equal(discovered.PolicyNames(), []string{"peer review"}), "discover: %d %s", discovered.Status, discovered.Raw)
}

// requires permission refuses a request without an intent, an intent without a user, and an intent it does not know
func TestProtection_RequiresPermission(t *testing.T) {
	realm := New(t)

	withoutIntent := realm.As("samwise").Request(CreateUser, nil, "")
	Expect(t, withoutIntent.Status == 406 && withoutIntent.Message() == "QAP intent header required", "without intent: %d %s", withoutIntent.Status, withoutIntent.Raw)

	withoutUser := realm.As("").Discovers(CreateUser)
	Expect(t, withoutUser.Status == 401, "without user: %d %s", withoutUser.Status, withoutUser.Raw)

	unknown := realm.As("samwise").Intent(CreateUser, "audit", nil, "")
	Expect(t, unknown.Status == 400 && unknown.Message() == "Unknown intent: 'audit'", "unknown intent: %d %s", unknown.Status, unknown.Raw)
}

// a permission placeholder without a matching route parameter is a mistake of the host
func TestProtection_MisconfiguredRoute(t *testing.T) {
	realm := New(t)
	got := realm.As("samwise").Discovers(MisconfiguredRoute)
	Expect(t, got.Status == 500 && got.Message() == "misconfigured permission template", "got %d %s", got.Status, got.Raw)
}
