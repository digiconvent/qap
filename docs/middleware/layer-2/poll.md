# I4 Poll

Returns the current [Submission view](../layer-4/submission-view.md) of one [submission](../concepts/submission.md) to someone with [Eligibility](../layer-3/eligibility.md). Clients poll while a form is open, so they see the votes come in and the outcome when it is decided. Every poll restarts the [Expiry](../layer-4/expiry.md) clock. Only the [author](../concepts/author.md) sees the data of a [sensitive read](../concepts/sensitive-read.md).

---

**Kind:** intent, `X-QAP-Intent: poll`.

**Headers:** `X-QAP-ID`: the submission.

**Response:** `200` with the [Submission view](../layer-4/submission-view.md) of one submission. Bob polls the submission of Alice while it waits for votes.

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

A [sensitive read](../concepts/sensitive-read.md) that was approved carries the requested data in `result.body`. This is what Alice, its author, sees:

```json
{
  "submission": {
    "id": "qap_8c1e4b7a2f9d0e3c6b5a4d2e1f0c9b8a",
    "endpoint": "/api/library/patron/456/history",
    "method": "GET",
    "body": "",
    "permission": "library.patron.456.history",
    "policy": {
      "id": "pol_librarians",
      "name": "two_librarians"
    },
    "author": "alice",
    "when": 1760000600,
    "resolution": "executed",
    "result": {
      "status": 200,
      "headers": {
        "content-type": "application/json"
      },
      "body": {
        "patron": "456",
        "borrowed": [
          "book 123",
          "book 77"
        ]
      }
    }
  },
  "poll": {
    "favour": 2,
    "against": 0,
    "withhold": 0,
    "comments": [],
    "resolution": "executed"
  }
}
```

Anyone else polling the same submission sees `"body": null` inside `result`.

**Errors:** [Unauthorized (401)](../concepts/errors.md#unauthorized-401), [Missing submission id (406)](../concepts/errors.md#missing-submission-id-406), [Unknown submission (404)](../concepts/errors.md#unknown-submission-404), [Not eligible (403)](../concepts/errors.md#not-eligible-403).

---

**Referenced in:**

- [A4 Expiry](../layer-4/expiry.md)
- [A10 Submission id](../layer-4/submission-id.md)
- [A11 Submission view](../layer-4/submission-view.md)
- [B2 Cleanup](../layer-3/cleanup.md)
- [B4 Eligibility](../layer-3/eligibility.md)
- [B5 Execution](../layer-3/execution.md)
- [B7 Rejection](../layer-3/rejection.md)
- [B9 Store](../layer-3/store.md)
- [C9 Sensitive read](../concepts/sensitive-read.md)
- [C10 Submission](../concepts/submission.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
- [I7 Withdraw](withdraw.md)
- [U2 Propose a change](../layer-0/propose-a-change.md)
- [U3 Review pending work](../layer-0/review-pending-work.md)
