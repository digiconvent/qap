# cS4 Executed

A submission that was approved and carried out. It carries the [Submission facts](../concepts/submission-facts.md) and the [Result](../concepts/result.md) of running the request. The only move is [Poll](../layer-2/poll.md).

---

**Kind:** state, `kind: 'executed'`, built by [Decided state](../layer-3/decided-state.md).

| Field | Type |
|---|---|
| the [Submission facts](../concepts/submission-facts.md) | |
| `result` | [Result](../concepts/result.md)` \| null` |

| Move | Leads to |
|---|---|
| `poll()` | any [submission state](../concepts/submission-state.md) |

---

**Referenced in:**

- [cB2 Decided state](../layer-3/decided-state.md)
- [cC6 Result](../concepts/result.md)
- [cC8 Submission state](../concepts/submission-state.md)
- [cE1 Endpoint](../layer-1/endpoint.md)
- [cM3 Submit](../layer-2/submit.md)
- [cM5 Poll](../layer-2/poll.md)
- [cM7 Vote](../layer-2/vote.md)
- [cS1 Discovered](discovered.md)
- [cU2 Propose a change](../layer-0/propose-a-change.md)
