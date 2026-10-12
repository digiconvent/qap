# cM9 Withdraw

Takes back a pending submission. Only its author succeeds, and only while it is pending.

---

**Kind:** move.

**Offered by:** [Pending](../states/pending.md) as `withdraw()`.

**Sends:** the route's method and path with `X-QAP-Intent: delete-submission` and `X-QAP-ID`, no body.

**Answers:** any success becomes [Withdrawn](../states/withdrawn.md); anything else becomes [Refused](../states/refused.md): 403 for someone else, 409 when already decided.

```ts
const after = await pending.withdraw()
```

---

**Referenced in:**

- [cB3 Pending state](../layer-3/pending-state.md)
- [cS3 Pending](../states/pending.md)
- [cS6 Withdrawn](../states/withdrawn.md)
- [cU4 Take back a proposal](../layer-0/take-back-a-proposal.md)
