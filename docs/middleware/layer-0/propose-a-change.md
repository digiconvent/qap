# U2 Propose a change

A person wants to do something that needs approval, such as deleting a book. They use [Discover](../layer-2/discover.md) to see which policies they may use, pick one and send a [Submit](../layer-2/submit.md). Then they share the [submission](../concepts/submission.md) with colleagues and keep using [Poll](../layer-2/poll.md) until it is decided.

---

**Kind:** user feature.

**Requests a client sends**, for deleting book 123:

1. `DELETE /api/library/book/123/delete` with `X-QAP-Intent: discover`
2. `DELETE /api/library/book/123/delete` with `X-QAP-Intent: submit` and `X-QAP-Policy-ID: <chosen policy>`
3. Repeatedly: `GET /api/library/book/123/delete` with `X-QAP-Intent: poll` and `X-QAP-ID: <id from step 2>`, until `resolution` is no longer `pending`
