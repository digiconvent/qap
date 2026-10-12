# B5 Execution

[Resolution](resolution.md) settles the [submission](../concepts/submission.md) as executed, and [Carry out](carry-out.md) has the [Endpoint](../host/endpoint.md) of its route check and execute the stored request. The outcome, or the failure if the request did not succeed, is stored on the submission, so everyone who [Poll](../layer-2/poll.md)s it sees what happened. An [Audit summary](../layer-4/audit-summary.md) goes to the [Audit trail](../host/audit-trail.md).

---

**Kind:** built from atomic parts and the host.

1. [Resolution](resolution.md) sets the submission to `executed`. Only the request that succeeds here continues; an outdated one leaves the submission pending.
2. [Carry out](carry-out.md) runs with the stored request. An endpoint that panics gives a result with status `500`, and nothing else is affected.
3. Its [result](../concepts/result.md) is stored in `result` of the [Submission view](../layer-4/submission-view.md). For a `GET`, the result body is the requested data.
4. The [Audit summary](../layer-4/audit-summary.md) goes to the [Audit trail](../host/audit-trail.md), with the changed fields that [Carry out](carry-out.md) found, unless the poll is anonymous.

---

**Referenced in:**

- [A11 Submission view](../layer-4/submission-view.md)
- [B1 Carry out](carry-out.md)
- [B6 Poll evaluation](poll-evaluation.md)
- [C8 Result](../concepts/result.md)
- [C10 Submission](../concepts/submission.md)
