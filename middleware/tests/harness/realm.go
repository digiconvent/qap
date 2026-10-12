package harness

import (
	"slices"
	"sync"
	"testing"
	"time"

	qap "github.com/digiconvent/qap/middleware"
	"github.com/digiconvent/qap/middleware/internal/host"
	"github.com/gofiber/fiber/v3"
)

type Realm struct {
	t           testing.TB
	app         *fiber.App
	clock       *clock
	mutex       sync.Mutex
	memberships map[string][]string
	policies    map[string][]qap.Policy
	people      map[string]map[string]any
	executed    []qap.Request[Payload]
	audited     []qap.AuditSummary
	refusing    bool
}

func New(t testing.TB) *Realm {
	r := &Realm{
		t:           t,
		clock:       &clock{now: time.Unix(1760000000, 0), ticks: make(chan time.Time)},
		memberships: memberships(),
		policies:    policies(),
		people:      map[string]map[string]any{FrodoId: {"name": "Frodo Baggins"}, SamwiseId: {"name": "Samwise Gamgee"}},
	}
	guard := host.Assemble(host.Host{
		CurrentUser:    r.CurrentUser,
		PolicyLookup:   r.PolicyLookup,
		IsMemberOf:     r.IsMemberOf,
		CountMembersOf: r.CountMembersOf,
		AuditTrail:     r,
	}, host.Clock{Now: r.clock.Now, AfterFunc: r.clock.AfterFunc, Ticks: r.clock.ticks})

	r.app = fiber.New()
	r.app.Get(PublicEvents.Path, guard.AllowAnonymous().Run(r.reached()))
	r.app.Get(LatestNews.Path, guard.AllowAuthenticatedOrRequiresPermission().Run(r.reached()))
	r.app.Post("/api/iam/user/:id/set-password", guard.AuthenticatedOrRequiresPermission("iam.user.set-password").Run(r.reached()))
	r.app.Post("/api/iam/user/create", guard.RequiresPermission("iam.user.create").Run(r.carriedOut()))
	r.app.Post("/api/iam/user/:id/update", guard.RequiresPermission("iam.user.update").Run(r.updated()))
	r.app.Post("/api/iam/user/:id/rename", guard.RequiresPermission("iam.user.update").Run(r.carriedOut()))
	r.app.Get("/api/iam/user/:id/history", guard.RequiresPermission("iam.user.history").Run(r.history()))
	r.app.Delete("/api/iam/user/:id/delete", guard.RequiresPermission("iam.user.delete").Run(r.carriedOut()))
	r.app.Post("/api/iam/group/create", guard.RequiresPermission("iam.group.create").Run(r.carriedOut()))
	r.app.Delete("/api/iam/group/:id/delete", guard.RequiresPermission("iam.group.delete").Run(r.carriedOut()))
	r.app.Post("/api/iam/group/:id/add-user", guard.RequiresPermission("iam.group.add-user").Run(r.carriedOut()))
	r.app.Post("/api/iam/group/:id/remove-user", guard.RequiresPermission("iam.group.remove-user").Run(r.carriedOut()))
	r.app.Post("/api/iam/policy/:id/update", guard.RequiresPermission("iam.policy.update").Run(r.carriedOut()))
	r.app.Delete("/api/iam/policy/:id/delete", guard.RequiresPermission("iam.policy.delete").Run(r.carriedOut()))
	r.app.Post("/api/iam/policy/:id/add-permission", guard.RequiresPermission("iam.policy.add-permission").Run(r.carriedOut()))
	r.app.Post(MusterTheArmy.Path, guard.RequiresPermission("rohan.army.muster").Run(r.carriedOut()))
	r.app.Post("/api/library/:id/add-book", guard.RequiresPermission("library.{id}.add-book").Run(r.carriedOut()))
	r.app.Post(MisconfiguredRoute.Path, guard.RequiresPermission("iam.group.{id}.remove-user").Run(r.carriedOut()))
	return r
}

func (r *Realm) CurrentUser(ctx fiber.Ctx) string {
	return ctx.Get("X-User")
}

func (r *Realm) PolicyLookup(permission string) []qap.Policy {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.policies[permission]
}

func (r *Realm) IsMemberOf(userId string, groupId string) bool {
	return slices.Contains(r.memberships[userId], groupId)
}

func (r *Realm) CountMembersOf(groupIds []string) int {
	count := 0
	for _, groups := range r.memberships {
		for _, group := range groups {
			if slices.Contains(groupIds, group) {
				count++
				break
			}
		}
	}
	return count
}

func (r *Realm) Record(summary qap.AuditSummary) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.audited = append(r.audited, summary)
}

func (r *Realm) RefusesEveryRequest() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.refusing = true
}

func (r *Realm) Waits(duration time.Duration) {
	r.clock.advance(duration)
}

func (r *Realm) CleansUp() {
	r.clock.ticks <- r.clock.Now()
	r.clock.ticks <- r.clock.Now()
}

func (r *Realm) SetPolicies(permission string, policies ...qap.Policy) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.policies[permission] = policies
}

func (r *Realm) Executed() []qap.Request[Payload] {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return slices.Clone(r.executed)
}

func (r *Realm) Audited() []qap.AuditSummary {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return slices.Clone(r.audited)
}

func (r *Realm) As(name string) Member {
	return Member{realm: r, name: name}
}

func Expect(t testing.TB, condition bool, format string, values ...any) {
	t.Helper()
	if !condition {
		t.Fatalf(format, values...)
	}
}
