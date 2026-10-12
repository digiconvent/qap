# cU1 Approve or reject

A person looks at the submission of someone else and casts a [Vote](../layer-2/vote.md): in favour, against or withhold. They can add a [Comment](../layer-2/comment.md) to explain their choice. The submission comes from [List](../layer-2/list.md), or from an id a colleague shared, opened with [Follow](../layer-2/follow.md).

---

**Kind:** user feature.

**Moves a client makes:**

```ts
const opened = await endpoint.follow('qap_9f2c4e1a7b3d4c5e8f6a0b1c2d3e4f5a')
if (opened.kind === 'pending') {
  const commented = await opened.comment('I will take the Ring, though I do not know the way.')
  if (commented.kind === 'pending') await commented.vote('favour')
}
```

A new vote replaces the earlier one. `'withhold'` is only accepted when the policy has `withholdPermitted`. A comment does not count as a vote. It travels in a header, so it works on every route, a proposed `GET` included.

Each move acts on the state the last move returned. An older [Pending](../states/pending.md) state can be decided already: acting on it gives [Refused](../states/refused.md) with status 409, and [Follow](../layer-2/follow.md) shows what happened.

---

**Referenced in:**

- [cU3 Review pending work](review-pending-work.md)
