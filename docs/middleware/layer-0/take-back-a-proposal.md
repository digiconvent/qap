# U4 Take back a proposal

A person changes their mind about something they proposed, before it is decided. They use [Withdraw](../layer-2/withdraw.md), and the [submission](../concepts/submission.md) disappears.

---

**Kind:** user feature.

**Requests a client sends:** `DELETE <endpoint>` with `X-QAP-Intent: delete-submission` and `X-QAP-ID: <submission>`.
