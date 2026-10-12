# cM4 Follow

Opens a submission by its id, for example one received in a link. It is [Poll](poll.md) without a state to start from.

---

**Kind:** move.

**Offered by:** [Endpoint](../layer-1/endpoint.md) as `follow(id)`.

**Sends:** the route's method and path with `X-QAP-Intent: poll` and `X-QAP-ID`, no body.

**Answers:** the same as [Poll](poll.md).

```ts
const opened = await endpoint.follow('qap_9f2c4e1a7b3d4c5e8f6a0b1c2d3e4f5a')
```

---

**Referenced in:**

- [cC7 Submission facts](../concepts/submission-facts.md)
- [cE1 Endpoint](../layer-1/endpoint.md)
- [cH2 Setup](../app/setup.md)
- [cU1 Approve or reject](../layer-0/approve-or-reject.md)
