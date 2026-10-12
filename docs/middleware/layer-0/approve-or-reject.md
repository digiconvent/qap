# U1 Approve or reject

A person looks at the [submission](../concepts/submission.md) of someone else and casts a [Vote](../layer-2/vote.md): in favour, against or withhold. They can add a [Comment](../layer-2/comment.md) to explain their choice.

---

**Kind:** user feature.

**Requests a client sends:**

1. `POST <endpoint>` with `X-QAP-Intent: vote`, `X-QAP-ID: <submission>` and `X-QAP-Vote: 1` (favour), `0` (withhold) or `-1` (against)
2. Optionally `POST <endpoint>` with `X-QAP-Intent: comment`, `X-QAP-ID: <submission>` and `X-QAP-Comment: <percent-encoded text>`
