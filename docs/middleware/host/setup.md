# H7 Setup

The host application sets up QAP once, when it starts. It tells QAP how to find the [Current user](current-user.md), how to find the policies of a permission through its [Policy lookup](policy-lookup.md), how to check [Membership](membership.md) of a user in a group, and how to get the [Member count](member-count.md) of a set of groups, and it hands over the [Audit trail](audit-trail.md) that receives every [Audit summary](../layer-4/audit-summary.md). Then it puts a [QAP endpoint](../layer-1/qap-endpoint.md) in front of every route, and gives each route its own [Endpoint](endpoint.md): the check and the handler of that route.

---

**Kind:** done by the host.

```go
guard := qap.New().
	ReadUserFrom(currentUser).
	LoadPoliciesFrom(policyLookup).
	VerifyEligibilityWith(isMemberOf).
	ComputePercentagesWith(countMembersOf).
	RecordAuditTo(auditTrail).
	AsMiddleware()

app.Delete("/api/library/book/:id/delete", guard.RequiresPermission("library.book.{id}.delete").Run(qap.Handle(deleteBook{library})))
app.Post("/api/iam/user/:id/set-password", guard.AuthenticatedOrRequiresPermission("iam.user.set-password", "iam.user.{id}.set-password").Run(qap.Handle(setPassword{users})))
app.Get("/api/forum/latest-news", guard.AllowAuthenticatedOrRequiresPermission().Run(qap.Handle(latestNews{forum})))
app.Get("/api/calendar/public-events", guard.AllowAnonymous().Run(qap.Handle(publicEvents{calendar})))
```

Every step is required, in this order: each step only offers the next one, so a missing step does not compile. Every argument is a named function or value of the host, never an inline function. Every route ends in `Run` with its [Endpoint](endpoint.md), wrapped by `qap.Handle`, so the host writes no Fiber handlers of its own for guarded routes. A route parameter that names an entity is always called `id`, and always holds a UUID.

### The arguments

QAP needs five things from the host application, plus an endpoint per route. Each is a function or value that the host defines and names itself. QAP calls them; it never looks into the database of the host.

#### currentUser

Tells who sends a request. It usually reads what an authentication middleware stored earlier, such as the user id from a verified token. It returns `""` when nobody is logged in. See [Current user](current-user.md).

```go
func currentUser(ctx fiber.Ctx) string {
	userId, _ := ctx.Locals("userId").(string)
	return userId
}
```

#### policyLookup

Returns the policies of a resolved permission, for example from the `permission_has_policy` table of the host. Each policy carries the ids of the groups that may vote. See [Policy lookup](policy-lookup.md).

```go
func policyLookup(permission string) []qap.Policy {
	return policyRepository.ForPermission(permission)
}
```

#### isMemberOf

Answers one question that every host can already answer: is this user a member of this group right now? QAP turns it into [Eligibility](../layer-3/eligibility.md) by asking about each group of a policy. The host keeps its membership rules to itself: start and end dates, inheritance between groups, guests. See [Membership](membership.md).

```go
func isMemberOf(userId string, groupId string) bool {
	return memberships.IsActive(userId, groupId, time.Now())
}
```

#### countMembersOf

Returns how many people are members of a set of groups right now, each person once. QAP needs it for thresholds and quorums given as a percentage, and to cap a threshold at the number of members. See [Member count](member-count.md).

```go
func countMembersOf(groupIds []string) int {
	return memberships.CountDistinctActive(groupIds, time.Now())
}
```

#### auditTrail

Receives every audit summary. It is a value with a `Record` method, so the host decides where summaries are kept. See [Audit trail](audit-trail.md).

```go
type auditLog struct{ database *sql.DB }

func (log auditLog) Record(summary qap.AuditSummary) {
	log.database.Exec("insert into audit (summary) values (?)", summary.JSON())
}

var auditTrail = auditLog{database: database}
```

#### an endpoint per route

Checks a request and carries it out. See [Endpoint](endpoint.md) for the interface and an example.

The names `memberships`, `policyRepository`, `deleteBook`, `setPassword`, `latestNews` and `publicEvents` stand for code of the host application, not of QAP.

---

**Referenced in:**

- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
- [H1 Audit trail](audit-trail.md)
- [H2 Current user](current-user.md)
- [H3 Endpoint](endpoint.md)
- [H4 Member count](member-count.md)
- [H5 Membership](membership.md)
- [H6 Policy lookup](policy-lookup.md)
