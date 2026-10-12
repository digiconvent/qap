# cA8 Read refusal

Turns an answer that is not a success into [Refused](../states/refused.md).

---

**Kind:** atomic part, `readRefusal(answer)`.

| Server field | Becomes |
|---|---|
| `error` | `message`, empty when missing |
| `hint` | `hint` |
| `acceptable_policies` | `acceptablePolicies` as [Policy reference](../concepts/policy-reference.md)s, empty when missing |
| `problems` | `problems`, empty when missing |

---

**Referenced in:**

- [cB4 Read state](../layer-3/read-state.md)
- [cS7 Refused](../states/refused.md)
