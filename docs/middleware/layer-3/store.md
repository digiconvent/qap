# B9 Store

The store keeps every [Submission](../concepts/submission.md) while QAP runs; the [Deadline timer](deadline-timer.md) keeps the timers. It lives in memory, so a restart starts empty. Every change to a submission happens in one step, so two requests never overwrite each other, and [Resolution](resolution.md) can settle a submission exactly once. [Cleanup](cleanup.md) removes submissions that are no longer needed.

---

**Kind:** built from atomic parts.

| Operation | Does |
|---|---|
| create | Stores a new submission under its [Submission id](../layer-4/submission-id.md) |
| get | Returns one submission |
| vote | Applies [Vote recording](../layer-4/vote-recording.md) to a pending submission, and raises its revision |
| comment | Adds a comment to a submission |
| resolve | Changes `pending` to `executed` or `rejected`, only once, and only at the revision that was evaluated (see [Resolution](resolution.md)) |
| set result | Stores the [Result](../concepts/result.md) |
| delete | Removes a pending submission (see [Withdraw](../layer-2/withdraw.md)) |
| list pending | Returns the pending submissions of one route template and method, for every id, in creation order |
| touch | Records the time of a [Poll](../layer-2/poll.md), for [Expiry](../layer-4/expiry.md) |
| remove expired | Removes every submission that [Expiry](../layer-4/expiry.md) allows (see [Cleanup](cleanup.md)) |

Every operation is safe when many requests arrive at the same time. The store keeps its own copy of every submission, the policy included, so the host changing its policies afterwards does not reach a stored submission.

---

**Referenced in:**

- [B2 Cleanup](cleanup.md)
- [B8 Resolution](resolution.md)
- [C10 Submission](../concepts/submission.md)
