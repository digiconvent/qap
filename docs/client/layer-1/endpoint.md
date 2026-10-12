# cE1 Endpoint

The starting state for one guarded route. `openEndpoint`, the only entry point of the library, creates it from [Endpoint options](../concepts/endpoint-options.md), and it stays valid for as long as you keep it. From here you can [Discover](../layer-2/discover.md), [Submit](../layer-2/submit.md), [List](../layer-2/list.md) or [Follow](../layer-2/follow.md). Every state without moves of its own, such as [Refused](../states/refused.md) and [Withdrawn](../states/withdrawn.md), leads back here.

---

**Kind:** endpoint, `openEndpoint(options)` returns the `Endpoint` interface.

| Move | Leads to |
|---|---|
| `discover()` | [Discovered](../states/discovered.md) or [Refused](../states/refused.md) |
| `submit(policyId, body?)` | [Pending](../states/pending.md), [Executed](../states/executed.md), [Rejected](../states/rejected.md) or [Refused](../states/refused.md) |
| `list()` | [Listed](../states/listed.md) or [Refused](../states/refused.md) |
| `follow(id)` | [Pending](../states/pending.md), [Executed](../states/executed.md), [Rejected](../states/rejected.md) or [Refused](../states/refused.md) |

```ts
const endpoint = openEndpoint({ path: '/api/council/ring/decide', method: 'POST', send: fetch })
```

---

**Referenced in:**

- [cC3 Endpoint options](../concepts/endpoint-options.md)
- [cH2 Setup](../app/setup.md)
- [cM1 Discover](../layer-2/discover.md)
- [cM2 List](../layer-2/list.md)
- [cM3 Submit](../layer-2/submit.md)
- [cM4 Follow](../layer-2/follow.md)
- [cS6 Withdrawn](../states/withdrawn.md)
- [cS7 Refused](../states/refused.md)
