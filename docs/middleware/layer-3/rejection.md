# B7 Rejection

[Resolution](resolution.md) settles the [submission](../concepts/submission.md) as rejected, and the request is never carried out. An [Audit summary](../layer-4/audit-summary.md) goes to the [Audit trail](../host/audit-trail.md). People who [Poll](../layer-2/poll.md) the submission see that it was rejected.

---

**Kind:** built from atomic parts and the host.

1. [Resolution](resolution.md) sets the submission to `rejected`. Only the request that succeeds here continues; an outdated one leaves the submission pending.
2. The [Audit summary](../layer-4/audit-summary.md) goes to the [Audit trail](../host/audit-trail.md).

`result` stays `null`.

---

**Referenced in:**

- [B3 Deadline timer](deadline-timer.md)
- [B6 Poll evaluation](poll-evaluation.md)
- [C8 Result](../concepts/result.md)
