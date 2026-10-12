# U3 Review pending work

A person checks what waits for their decision. They use [List](../layer-2/list.md) to see the pending [submissions](../concepts/submission.md) they may vote on, and [Poll](../layer-2/poll.md) to look at one of them in detail.

---

**Kind:** user feature.

**Requests a client sends:**

1. `GET <endpoint>` with `X-QAP-Intent: list`
2. `GET <endpoint>` with `X-QAP-Intent: poll` and `X-QAP-ID: <submission>` for the one to look at
