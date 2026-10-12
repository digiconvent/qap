# cA4 Not QAP

The [Refused](../states/refused.md) state for an answer that was a success but not a QAP answer, for example from a route without a guard, or a page of a proxy.

---

**Kind:** atomic part, `notQap(answer)`.

The status is the one the server sent:

```ts
{ kind: 'refused', status: 200, message: 'not a QAP answer', hint: undefined, acceptablePolicies: [] }
```

---

**Referenced in:**

- [cA1 Is discovery](is-discovery.md)
- [cA3 Is view](is-view.md)
- [cB4 Read state](../layer-3/read-state.md)
- [cM1 Discover](../layer-2/discover.md)
- [cM2 List](../layer-2/list.md)
