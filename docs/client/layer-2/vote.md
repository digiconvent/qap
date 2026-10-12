# cM7 Vote

Votes on a pending submission with a [Vote choice](../concepts/vote-choice.md). A new vote replaces your earlier one. The answer shows the effect right away: still [Pending](../states/pending.md), or decided as [Executed](../states/executed.md) or [Rejected](../states/rejected.md).

---

**Kind:** move.

**Offered by:** [Pending](../states/pending.md) as `vote(choice)`.

**Sends:** the route's method and path with `X-QAP-Intent: vote`, `X-QAP-ID` and `X-QAP-Vote` from [Vote value](../layer-4/vote-value.md), no body.

**Answers:** `200` becomes a [submission state](../concepts/submission-state.md) through [Read state](../layer-3/read-state.md); anything else becomes [Refused](../states/refused.md).

```ts
const after = await pending.vote('against')
```

---

**Referenced in:**

- [cA10 Vote value](../layer-4/vote-value.md)
- [cB3 Pending state](../layer-3/pending-state.md)
- [cC10 Vote choice](../concepts/vote-choice.md)
- [cS3 Pending](../states/pending.md)
- [cU1 Approve or reject](../layer-0/approve-or-reject.md)
