# cA10 Vote value

Turns a [Vote choice](../concepts/vote-choice.md) into the `X-QAP-Vote` header for [Vote](../layer-2/vote.md).

---

**Kind:** atomic part, `voteValue(choice)`.

| Choice | Header |
|---|---|
| `'favour'` | `1` |
| `'against'` | `-1` |
| `'withhold'` | `0` |

---

**Referenced in:**

- [cC10 Vote choice](../concepts/vote-choice.md)
- [cM7 Vote](../layer-2/vote.md)
