# I5 Vote

Someone with [Eligibility](../layer-3/eligibility.md) votes on a [submission](../concepts/submission.md) they found through [List](list.md) or received as a [Submission id](../layer-4/submission-id.md) from a colleague. The [Vote value](../layer-4/vote-value.md) says favour, against or withhold, and the [Withhold rule](../layer-4/withhold-rule.md) decides whether withholding is allowed. After [Vote recording](../layer-4/vote-recording.md), [Poll evaluation](../layer-3/poll-evaluation.md) may decide the submission. The answer is the [Submission view](../layer-4/submission-view.md), so the voter sees the effect of their vote right away.

---

**Kind:** intent, `X-QAP-Intent: vote`.

**Headers:** `X-QAP-ID`: the submission. `X-QAP-Vote`: the [Vote value](../layer-4/vote-value.md).

**Response:** `200` with the [Submission view](../layer-4/submission-view.md) after the vote. Bob votes in favour (`X-QAP-Vote: 1`). That is the second favour vote, which meets the threshold of the policy, so the submission is executed and `result` holds what the [Endpoint](../host/endpoint.md) of the route returned.

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
    "resolution": "executed",
    "result": {
      "status": 204,
      "headers": {},
      "body": null
    }
  },
  "poll": {
    "favour": 2,
    "against": 0,
    "withhold": 0,
    "comments": [
      {
        "author": "bob",
        "at": 1760000300,
        "text": "This copy is damaged beyond repair."
      }
    ],
    "resolution": "executed"
  }
}
```

**Errors:** [Unauthorized (401)](../concepts/errors.md#unauthorized-401), [Missing submission id (406)](../concepts/errors.md#missing-submission-id-406), [Unknown submission (404)](../concepts/errors.md#unknown-submission-404), [Already resolved (409)](../concepts/errors.md#already-resolved-409), [Not eligible (403)](../concepts/errors.md#not-eligible-403), [Invalid vote (400)](../concepts/errors.md#invalid-vote-400), [Withhold not permitted (400)](../concepts/errors.md#withhold-not-permitted-400).

The id alone names the submission. The endpoint the vote is sent to does not matter, and neither does the HTTP method.

---

**Referenced in:**

- [A10 Submission id](../layer-4/submission-id.md)
- [A11 Submission view](../layer-4/submission-view.md)
- [A15 Vote value](../layer-4/vote-value.md)
- [A16 Withhold rule](../layer-4/withhold-rule.md)
- [B3 Deadline timer](../layer-3/deadline-timer.md)
- [B4 Eligibility](../layer-3/eligibility.md)
- [C9 Sensitive read](../concepts/sensitive-read.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
- [I3 Submit](submit.md)
- [U1 Approve or reject](../layer-0/approve-or-reject.md)
