# I3 Submit

The [Current user](../host/current-user.md) sends the request they want approved, together with the id of the [policy](../concepts/policy.md) they picked. They find the policies they may use through [Discover](discover.md); a policy they have no [Eligibility](../layer-3/eligibility.md) for is refused. After a [Request copy](../layer-4/request-copy.md), the [Endpoint](../host/endpoint.md) of the route checks the request: one it cannot carry out is refused, because a submission of it would be useless. Then the [submission](../concepts/submission.md) keeps the endpoint, gets a [Submission id](../layer-4/submission-id.md) and its [author](../concepts/author.md) gets a favour vote through [Vote recording](../layer-4/vote-recording.md). [Poll evaluation](../layer-3/poll-evaluation.md) may decide it straight away, when one vote is enough, and a [Deadline timer](../layer-3/deadline-timer.md) starts when the policy has a deadline. The answer is the [Submission view](../layer-4/submission-view.md).

---

**Kind:** intent, `X-QAP-Intent: submit`.

**Headers:** `X-QAP-Policy-ID`: the id of the chosen policy. The body of the request is the body of the request being submitted.

**Response:** `202` with the [Submission view](../layer-4/submission-view.md) of the new submission. Alice proposes deleting book 123 under the policy found with Discover. Her own vote counts as the first favour vote; the policy needs 2, so the submission is pending. The new id is `submission.id`.

```json
{
  "submission": {
    "id": "qap_3f9a0c2e7b1d4e8fa6c5b2d1e0f9a8b7",
    "endpoint": "/api/library/book/123/delete",
    "method": "DELETE",
    "body": "",
    "permission": "library.book.123.delete",
    "policy": {
      "id": "pol_librarians",
      "name": "two_librarians"
    },
    "author": "alice",
    "when": 1760000000,
    "resolution": "pending",
    "result": null
  },
  "poll": {
    "favour": 1,
    "against": 0,
    "withhold": 0,
    "comments": [],
    "resolution": "pending"
  }
}
```

When one vote is already enough, the same response comes back with `resolution` set to `executed` and `result` filled in, as in the example on [Vote](vote.md).

**Errors:** [Unauthorized (401)](../concepts/errors.md#unauthorized-401), [Missing policy id (406)](../concepts/errors.md#missing-policy-id-406), [Policy does not include you (403)](../concepts/errors.md#policy-does-not-include-you-403), [Cannot be carried out (422)](../concepts/errors.md#cannot-be-carried-out-422).

---

**Referenced in:**

- [A2 Authorship](../layer-4/authorship.md)
- [A6 Permission](../layer-4/permission.md)
- [A10 Submission id](../layer-4/submission-id.md)
- [A11 Submission view](../layer-4/submission-view.md)
- [B3 Deadline timer](../layer-3/deadline-timer.md)
- [B4 Eligibility](../layer-3/eligibility.md)
- [B10 Voter count](../layer-3/voter-count.md)
- [C2 Author](../concepts/author.md)
- [C3 Errors](../concepts/errors.md)
- [C5 Policy](../concepts/policy.md)
- [C9 Sensitive read](../concepts/sensitive-read.md)
- [C10 Submission](../concepts/submission.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
- [H3 Endpoint](../host/endpoint.md)
- [I1 Discover](discover.md)
- [U2 Propose a change](../layer-0/propose-a-change.md)
