# cM8 Comment

Adds a [Comment](../concepts/comment.md) to a pending submission. It does not count as a vote. The text travels in a header, so it works on every route, a proposed `GET` included.

---

**Kind:** move.

**Offered by:** [Pending](../states/pending.md) as `comment(text)`.

**Sends:** the route's method and path with `X-QAP-Intent: comment`, `X-QAP-ID`, and `X-QAP-Comment` with the text percent-encoded by `encodeURIComponent`, no body. The server accepts at most 500 characters.

**Answers:** `200` becomes a [submission state](../concepts/submission-state.md) through [Read state](../layer-3/read-state.md); anything else becomes [Refused](../states/refused.md).

```ts
const after = await pending.comment('Do not be hasty.')
```

---

**Referenced in:**

- [cB3 Pending state](../layer-3/pending-state.md)
- [cC1 Comment](../concepts/comment.md)
- [cS3 Pending](../states/pending.md)
- [cU1 Approve or reject](../layer-0/approve-or-reject.md)
