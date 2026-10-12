# cS1 Discovered

The answer to [Discover](../layer-2/discover.md) from a server that speaks the same [Version](../concepts/version.md) as the client: the permission behind the route, that version, and the policies that include you, each as [Policy settings](../concepts/policy-settings.md). The only move is [Submit](../layer-2/submit.md) under one of those policies. An empty list means no policy includes you.

---

**Kind:** state, `kind: 'discovered'`.

| Field | Type |
|---|---|
| `endpoint` | `string` |
| `permission` | `string` |
| `version` | `string`, the version that both the server and the client speak |
| `policies` | [Policy settings](../concepts/policy-settings.md)`[]` |

| Move | Leads to |
|---|---|
| `submit(policy, body?)` | [Pending](pending.md), [Executed](executed.md), [Rejected](rejected.md) or [Refused](refused.md) |

```ts
{ kind: 'discovered', endpoint: '/api/council/ring/decide', permission: 'council.ring.decide', version: '-1.-1.-1',
  policies: [{ id: "the council's consent", name: "the council's consent", favourThreshold: 4, againstThreshold: 2, quorum: null, deadlineSeconds: null, extensionSeconds: null, withholdPermitted: false }] }
```

---

**Referenced in:**

- [cC5 Policy settings](../concepts/policy-settings.md)
- [cE1 Endpoint](../layer-1/endpoint.md)
- [cM1 Discover](../layer-2/discover.md)
- [cM3 Submit](../layer-2/submit.md)
