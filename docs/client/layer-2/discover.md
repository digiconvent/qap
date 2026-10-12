# cM1 Discover

Asks the route which policies you may use. The answer is [Discovered](../states/discovered.md), listing only the policies that include you.

---

**Kind:** move.

**Offered by:** [Endpoint](../layer-1/endpoint.md) as `discover()`.

**Sends:** the route's method and path with `X-QAP-Intent: discover`, no body.

**Answers:** `200` becomes [Discovered](../states/discovered.md) through [Read policy](../layer-4/read-policy.md), unless [Is discovery](../layer-4/is-discovery.md) rejects it and it becomes [Not QAP](../layer-4/not-qap.md), or [Version mismatch](../layer-4/version-mismatch.md) finds that the server speaks another version; anything else becomes [Refused](../states/refused.md).

```ts
const discovered = await endpoint.discover()
```

---

**Referenced in:**

- [cA1 Is discovery](../layer-4/is-discovery.md)
- [cA7 Read policy](../layer-4/read-policy.md)
- [cA9 Version mismatch](../layer-4/version-mismatch.md)
- [cC9 Version](../concepts/version.md)
- [cE1 Endpoint](../layer-1/endpoint.md)
- [cS1 Discovered](../states/discovered.md)
- [cU2 Propose a change](../layer-0/propose-a-change.md)
