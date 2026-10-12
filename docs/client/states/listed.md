# cS2 Listed

The answer to [List](../layer-2/list.md): the pending submissions of the route template and method of this endpoint, for every id, that you may vote on. The `endpoint` of each one tells which id it is about. Each one is a [Pending](pending.md) state, so its moves can be used directly. The listed state has no moves of its own.

---

**Kind:** state, `kind: 'listed'`.

| Field | Type |
|---|---|
| `submissions` | [Pending](pending.md)`[]` |

---

**Referenced in:**

- [cE1 Endpoint](../layer-1/endpoint.md)
- [cM2 List](../layer-2/list.md)
- [cU3 Review pending work](../layer-0/review-pending-work.md)
