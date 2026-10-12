# cM2 List

Asks for the pending submissions that you may vote on, of the route template and method of this endpoint, for every id: listing on the password reset of one user shows the pending resets of all users. The answer is [Listed](../states/listed.md), each submission a [Pending](../states/pending.md) state.

---

**Kind:** move.

**Offered by:** [Endpoint](../layer-1/endpoint.md) as `list()`.

**Sends:** the route's method and path with `X-QAP-Intent: list`, no body.

**Answers:** `200` becomes [Listed](../states/listed.md) through [Pending state](../layer-3/pending-state.md), unless one of its views fails [Is view](../layer-4/is-view.md) or is not pending, and it becomes [Not QAP](../layer-4/not-qap.md); anything else becomes [Refused](../states/refused.md).

```ts
const listed = await endpoint.list()
```

---

**Referenced in:**

- [cA3 Is view](../layer-4/is-view.md)
- [cE1 Endpoint](../layer-1/endpoint.md)
- [cS2 Listed](../states/listed.md)
- [cU1 Approve or reject](../layer-0/approve-or-reject.md)
- [cU3 Review pending work](../layer-0/review-pending-work.md)
