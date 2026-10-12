# cC7 Submission facts

What every submission state knows about its submission: the request, who proposed it under which policy, the [Counts](counts.md) and the [Comments](comment.md). [Read facts](../layer-4/read-facts.md) builds them. They are a snapshot of the last answer.

---

**Kind:** concept, `SubmissionFacts`.

| Field | Type | Meaning |
|---|---|---|
| `id` | `string` | Use it with [Follow](../layer-2/follow.md). |
| `endpoint`, `method`, `body` | `string` | The proposed request. |
| `permission` | `string` | The permission it needs. |
| `policy` | [Policy reference](policy-reference.md) | The policy it was proposed under. |
| `author` | `string` | Who proposed it. |
| `when` | `number` | When, in Unix seconds. |
| `counts` | [Counts](counts.md) | The votes so far. |
| `comments` | [Comment](comment.md)`[]` | The comments so far. |

---

**Referenced in:**

- [cA6 Read facts](../layer-4/read-facts.md)
- [cB2 Decided state](../layer-3/decided-state.md)
- [cB3 Pending state](../layer-3/pending-state.md)
- [cC1 Comment](comment.md)
- [cC2 Counts](counts.md)
- [cS3 Pending](../states/pending.md)
- [cS4 Executed](../states/executed.md)
- [cS5 Rejected](../states/rejected.md)
