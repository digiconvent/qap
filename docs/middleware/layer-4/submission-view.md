# A11 Submission view

Shapes a [submission](../concepts/submission.md) into the response that [Submit](../layer-2/submit.md), [Vote](../layer-2/vote.md), [Comment](../layer-2/comment.md), [Poll](../layer-2/poll.md) and [List](../layer-2/list.md) return. It contains the submission itself (id, endpoint, method, body, permission, [policy](../concepts/policy.md), [author](../concepts/author.md), time, resolution and [result](../concepts/result.md)) and its votes, vote counts and comments. Voters see the body, so they know what they approve, but only the author sees the data that a read returns.

---

**Kind:** atomic part.

**In:** a submission and the user asking. **Out:** this JSON shape.

```json
{
  "submission": {
    "id": "qap_...",
    "endpoint": "/api/library/book/123/delete",
    "method": "DELETE",
    "body": "",
    "permission": "library.book.123.delete",
    "policy": { "id": "<policy id>", "name": "<policy name>" },
    "author": "<user id>",
    "when": 1760000000,
    "resolution": "pending | executed | rejected",
    "result": null
  },
  "poll": {
    "favour": 1,
    "against": 0,
    "withhold": 0,
    "comments": [{ "author": "<user id>", "at": 1760000000, "text": "..." }],
    "resolution": "pending | executed | rejected"
  }
}
```

`result` stays `null` until [Execution](../layer-3/execution.md), then holds `{"status", "headers", "body"}`. For a `GET` submission, only the author sees `result.body`.

---

**Referenced in:**

- [B5 Execution](../layer-3/execution.md)
- [C1 Anonymity](../concepts/anonymity.md)
- [C8 Result](../concepts/result.md)
- [C9 Sensitive read](../concepts/sensitive-read.md)
- [C10 Submission](../concepts/submission.md)
- [I2 List](../layer-2/list.md)
- [I3 Submit](../layer-2/submit.md)
- [I4 Poll](../layer-2/poll.md)
- [I5 Vote](../layer-2/vote.md)
- [I6 Comment](../layer-2/comment.md)
