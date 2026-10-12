# cA7 Read policy

Turns a policy as the server sends it into [Policy settings](../concepts/policy-settings.md). Used by [Discover](../layer-2/discover.md).

---

**Kind:** atomic part, `readPolicy(policy)`.

| Server field | Becomes |
|---|---|
| `favour_threshold` | `favourThreshold` |
| `against_threshold` | `againstThreshold` |
| `quorum` | `quorum` |
| `deadline_seconds` | `deadlineSeconds` |
| `extension_seconds` | `extensionSeconds`, `null` when missing |
| `withhold_permitted` | `withholdPermitted` |

---

**Referenced in:**

- [cC5 Policy settings](../concepts/policy-settings.md)
- [cM1 Discover](../layer-2/discover.md)
