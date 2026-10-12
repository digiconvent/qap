# cM5 Poll

Asks for the current state of a submission. Use it to bring an old state up to date, or to wait for a decision.

---

**Kind:** move.

**Offered by:** [Pending](../states/pending.md), [Executed](../states/executed.md) and [Rejected](../states/rejected.md) as `poll()`.

**Sends:** the route's method and path with `X-QAP-Intent: poll` and `X-QAP-ID`, no body.

**Answers:** `200` becomes a [submission state](../concepts/submission-state.md) through [Read state](../layer-3/read-state.md); anything else becomes [Refused](../states/refused.md).

```ts
const now = await pending.poll()
```

---

**Referenced in:**

- [cB2 Decided state](../layer-3/decided-state.md)
- [cB3 Pending state](../layer-3/pending-state.md)
- [cM4 Follow](follow.md)
- [cM6 Listen](listen.md)
- [cS3 Pending](../states/pending.md)
- [cS4 Executed](../states/executed.md)
- [cS5 Rejected](../states/rejected.md)
