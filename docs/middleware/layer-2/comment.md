# I6 Comment

Someone with [Eligibility](../layer-3/eligibility.md) adds a remark to a [submission](../concepts/submission.md), for example to explain why they vote against. The comment is stored on the submission with its [author](../concepts/author.md) and time, and the answer is the [Submission view](../layer-4/submission-view.md).

---

**Kind:** intent, `X-QAP-Intent: comment`.

**Headers:** `X-QAP-ID`: the submission. `X-QAP-Comment`: the comment text, percent-encoded, at most 500 characters, so it works on every method, `GET` included. The body of the request is ignored.

**Response:** `200` with the [Submission view](../layer-4/submission-view.md), including the new comment. Bob explains why he supports the deletion.

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
    "comments": [
      {
        "author": "bob",
        "at": 1760000300,
        "text": "This copy is damaged beyond repair."
      }
    ],
    "resolution": "pending"
  }
}
```

**Errors:** [Unauthorized (401)](../concepts/errors.md#unauthorized-401), [Missing submission id (406)](../concepts/errors.md#missing-submission-id-406), [Unknown submission (404)](../concepts/errors.md#unknown-submission-404), [Not eligible (403)](../concepts/errors.md#not-eligible-403), [Missing comment (406)](../concepts/errors.md#missing-comment-406), [Invalid comment (400)](../concepts/errors.md#invalid-comment-400), [Comment too long (400)](../concepts/errors.md#comment-too-long-400).

---

**Referenced in:**

- [A10 Submission id](../layer-4/submission-id.md)
- [A11 Submission view](../layer-4/submission-view.md)
- [B4 Eligibility](../layer-3/eligibility.md)
- [C3 Errors](../concepts/errors.md)
- [C10 Submission](../concepts/submission.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
- [U1 Approve or reject](../layer-0/approve-or-reject.md)
