# C10 Submission

A submission is a request that waits for approval before it is carried out. It is created by [Submit](../layer-2/submit.md) and holds the request itself, the [Policy](policy.md) it was submitted under, its [Author](author.md), the votes and comments on it, and, once it is decided, its [Result](result.md). Every submission has a [Submission id](../layer-4/submission-id.md). It stays pending until [Poll evaluation](../layer-3/poll-evaluation.md) decides it, and then it is executed or rejected for good. Until then, the author can [Withdraw](../layer-2/withdraw.md) it.

---

**Kind:** concept.

| Field | Meaning |
|---|---|
| id | The [Submission id](../layer-4/submission-id.md) |
| endpoint | The URL of the request, including the query string |
| method | The HTTP method of the request |
| headers | The headers of the request, without the `X-QAP-` ones (see [Request copy](../layer-4/request-copy.md)) |
| body | The body of the request |
| route | The route template the request matched, such as `/api/library/book/:id/delete`, for [List](../layer-2/list.md) |
| params | The route parameters by name, such as `id` → `123` |
| handler | The [Endpoint](../host/endpoint.md) of the route, which checks and carries out the request |
| permission | The resolved [Permission](../layer-4/permission.md), such as `library.book.123.delete` |
| policy | The [Policy](policy.md), with the settings it had when the submission was created |
| author | The user id of the [Author](author.md) |
| when | Creation time, in Unix seconds |
| resolution | `pending`, `executed` or `rejected` |
| result | The [Result](result.md), `null` until [Execution](../layer-3/execution.md) |
| votes | Three lists of user ids: favour, against and withhold (see [Vote recording](../layer-4/vote-recording.md)) |
| comments | Author, time and text of every comment (see [Comment](../layer-2/comment.md)) |
| last polled | The time of the last [Poll](../layer-2/poll.md) or of the decision, for [Expiry](../layer-4/expiry.md) |
| revision | Raised by every vote, so [Resolution](../layer-3/resolution.md) only settles the votes it evaluated |
| voter count | The [Voter count](../layer-3/voter-count.md) of the policy when the submission was created |

A submission lives in the [Store](../layer-3/store.md) and is shown through the [Submission view](../layer-4/submission-view.md). Its request, with the author as user, is the [Request](request.md) that the endpoint carries out.

---

**Referenced in:**

- [A1 Audit summary](../layer-4/audit-summary.md)
- [A2 Authorship](../layer-4/authorship.md)
- [A3 Deadline](../layer-4/deadline.md)
- [A4 Expiry](../layer-4/expiry.md)
- [A5 Intent](../layer-4/intent.md)
- [A8 Request copy](../layer-4/request-copy.md)
- [A10 Submission id](../layer-4/submission-id.md)
- [A11 Submission view](../layer-4/submission-view.md)
- [A13 Vote count](../layer-4/vote-count.md)
- [A14 Vote recording](../layer-4/vote-recording.md)
- [A16 Withhold rule](../layer-4/withhold-rule.md)
- [B1 Carry out](../layer-3/carry-out.md)
- [B2 Cleanup](../layer-3/cleanup.md)
- [B3 Deadline timer](../layer-3/deadline-timer.md)
- [B5 Execution](../layer-3/execution.md)
- [B6 Poll evaluation](../layer-3/poll-evaluation.md)
- [B7 Rejection](../layer-3/rejection.md)
- [B8 Resolution](../layer-3/resolution.md)
- [B9 Store](../layer-3/store.md)
- [B10 Voter count](../layer-3/voter-count.md)
- [C2 Author](author.md)
- [C3 Errors](errors.md)
- [C4 Extension](extension.md)
- [C5 Policy](policy.md)
- [C6 Quorum](quorum.md)
- [C7 Request](request.md)
- [C8 Result](result.md)
- [C9 Sensitive read](sensitive-read.md)
- [C11 Threshold](threshold.md)
- [I2 List](../layer-2/list.md)
- [I3 Submit](../layer-2/submit.md)
- [I4 Poll](../layer-2/poll.md)
- [I5 Vote](../layer-2/vote.md)
- [I6 Comment](../layer-2/comment.md)
- [I7 Withdraw](../layer-2/withdraw.md)
- [U1 Approve or reject](../layer-0/approve-or-reject.md)
- [U2 Propose a change](../layer-0/propose-a-change.md)
- [U3 Review pending work](../layer-0/review-pending-work.md)
- [U4 Take back a proposal](../layer-0/take-back-a-proposal.md)
