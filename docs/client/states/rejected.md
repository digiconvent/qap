# cS5 Rejected

A submission that was voted down, or whose deadline passed without enough votes. Nothing was carried out. It carries the [Submission facts](../concepts/submission-facts.md), and the only move is [Poll](../layer-2/poll.md).

---

**Kind:** state, `kind: 'rejected'`, built by [Decided state](../layer-3/decided-state.md).

| Field | Type |
|---|---|
| the [Submission facts](../concepts/submission-facts.md) | |
| `result` | `null` |

| Move | Leads to |
|---|---|
| `poll()` | any [submission state](../concepts/submission-state.md) |

---

**Referenced in:**

- [cB2 Decided state](../layer-3/decided-state.md)
- [cC8 Submission state](../concepts/submission-state.md)
- [cE1 Endpoint](../layer-1/endpoint.md)
- [cM5 Poll](../layer-2/poll.md)
- [cM7 Vote](../layer-2/vote.md)
- [cS1 Discovered](discovered.md)
