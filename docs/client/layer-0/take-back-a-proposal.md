# cU4 Take back a proposal

A person changes their mind about something they proposed, before it is decided. They use [Withdraw](../layer-2/withdraw.md), and the submission disappears from the list of everyone.

---

**Kind:** user feature.

**Moves a client makes:**

```ts
const proposed = await endpoint.submit('the fellowship agrees', JSON.stringify({ route: 'by way of Minas Tirith' }))
if (proposed.kind === 'pending') await proposed.withdraw()
```

The answer is [Withdrawn](../states/withdrawn.md), or [Refused](../states/refused.md): 403 for the submission of someone else, 409 when it was decided first.
