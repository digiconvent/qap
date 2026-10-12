# cU3 Review pending work

A person checks what waits for their decision. They use [List](../layer-2/list.md) to see the pending submissions of a kind of route that they may vote on, such as every pending password reset, whichever user it is for. Each one is a [Pending](../states/pending.md) state, ready for [Approve or reject](approve-or-reject.md).

---

**Kind:** user feature.

**Moves a client makes:**

```ts
const listed = await endpoint.list()
if (listed.kind === 'listed') {
  for (const proposal of listed.submissions) {
    console.log(proposal.author, proposal.body, proposal.counts)
  }
}
```

[Listed](../states/listed.md) holds the submissions of the route template and method, for every id, whose policy includes the person, their own among them. Decided and withdrawn submissions do not appear.
