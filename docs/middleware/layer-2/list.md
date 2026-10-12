# I2 List

Shows what waits for a decision. It returns the pending [submissions](../concepts/submission.md) of the route template and method of the endpoint, for every id, that the [Current user](../host/current-user.md) has [Eligibility](../layer-3/eligibility.md) for, each as a [Submission view](../layer-4/submission-view.md). Listing on `POST /api/iam/user/frodo/set-password`, of the route `/api/iam/user/:id/set-password`, shows the pending password resets of Frodo and of Sam. Another route stays apart, even under the same [Permission](../layer-4/permission.md). Decided submissions do not appear.

---

**Kind:** intent, `X-QAP-Intent: list`.

**Headers:** no others.

**Response:** `200` with a JSON array of [Submission view](../layer-4/submission-view.md)s: every pending submission of the route template and method of this endpoint that the current user may vote on. The permission is not resolved, so a route that misses a placeholder still lists. Eligibility is checked against the policy of each submission, so a submission for another id only appears when its policy includes the user. Bob lists the pending deletions of book 123 and finds the one by Alice. The array is empty when nothing is pending.

```json
[
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
]
```

**Errors:** [Unauthorized (401)](../concepts/errors.md#unauthorized-401).

---

**Referenced in:**

- [A6 Permission](../layer-4/permission.md)
- [A8 Request copy](../layer-4/request-copy.md)
- [A11 Submission view](../layer-4/submission-view.md)
- [B4 Eligibility](../layer-3/eligibility.md)
- [C10 Submission](../concepts/submission.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
- [I5 Vote](vote.md)
- [U3 Review pending work](../layer-0/review-pending-work.md)
