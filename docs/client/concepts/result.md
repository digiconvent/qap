# cC6 Result

What the route returned when an approved submission was carried out. Only [Executed](../states/executed.md) has one. For a proposed `GET`, the body is there only for the author; everyone else gets `null`.

---

**Kind:** concept, `Result`.

| Field | Type |
|---|---|
| `status` | `number` |
| `headers` | `Record<string, string>` |
| `body` | `unknown` |

---

**Referenced in:**

- [cB2 Decided state](../layer-3/decided-state.md)
- [cS4 Executed](../states/executed.md)
- [cU2 Propose a change](../layer-0/propose-a-change.md)
