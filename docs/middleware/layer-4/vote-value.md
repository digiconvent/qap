# A15 Vote value

A [Vote](../layer-2/vote.md) carries its choice in the `X-QAP-Vote` header as a number. A positive number means favour, `0` means withhold and a negative number means against. Anything that is not a number is invalid, and the vote is refused.

---

**Kind:** atomic part.

**In:** the value of the `X-QAP-Vote` header. **Out:** favour, withhold or against.

| Header | Vote |
|---|---|
| `1`, `2`, ... | favour |
| `0` | withhold |
| `-1`, `-2`, ... | against |
| anything else | invalid: [Invalid vote (400)](../concepts/errors.md#invalid-vote-400) |

---

**Referenced in:**

- [C3 Errors](../concepts/errors.md)
- [I5 Vote](../layer-2/vote.md)
