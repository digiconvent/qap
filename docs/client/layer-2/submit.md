# cM3 Submit

Proposes the request under a policy. The answer is the submission in its new state: [Executed](../states/executed.md) when one vote is enough, otherwise [Pending](../states/pending.md). Without a policy, or with one that does not include you, it is [Refused](../states/refused.md).

---

**Kind:** move.

**Offered by:** [Endpoint](../layer-1/endpoint.md) as `submit(policyId, body?)`, and [Discovered](../states/discovered.md) as `submit(policy, body?)` with one of its [Policy settings](../concepts/policy-settings.md).

**Sends:** the route's method and path with `X-QAP-Intent: submit` and `X-QAP-Policy-ID`, and the body when it is not empty.

**Answers:** `202` becomes a [submission state](../concepts/submission-state.md) through [Read state](../layer-3/read-state.md); anything else becomes [Refused](../states/refused.md).

```ts
const proposed = await endpoint.submit("the council's consent", JSON.stringify({ fate: 'destroy it' }))
```

---

**Referenced in:**

- [cE1 Endpoint](../layer-1/endpoint.md)
- [cS1 Discovered](../states/discovered.md)
- [cS7 Refused](../states/refused.md)
- [cU2 Propose a change](../layer-0/propose-a-change.md)
