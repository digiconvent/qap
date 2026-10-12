# C8 Result

The result is the outcome of carrying out a submission: what the endpoint answered. It is empty until [Execution](../layer-3/execution.md), and stays empty after [Rejection](../layer-3/rejection.md). A failed execution has a result too, so everyone can see what went wrong. For a [Sensitive read](sensitive-read.md), the result holds the requested data, and only the [Author](author.md) sees it.

---

**Kind:** concept.

```json
{ "status": 204, "headers": { "content-type": "application/json" }, "body": null }
```

Produced by the [Endpoint](../host/endpoint.md), or by [Carry out](../layer-3/carry-out.md) when the request no longer passes its check, stored on the [Submission](submission.md), shown in the [Submission view](../layer-4/submission-view.md) and written into the [Audit summary](../layer-4/audit-summary.md).

---

**Referenced in:**

- [A1 Audit summary](../layer-4/audit-summary.md)
- [A6 Permission](../layer-4/permission.md)
- [A11 Submission view](../layer-4/submission-view.md)
- [A17 Write result](../layer-4/write-result.md)
- [B1 Carry out](../layer-3/carry-out.md)
- [B5 Execution](../layer-3/execution.md)
- [B9 Store](../layer-3/store.md)
- [C9 Sensitive read](sensitive-read.md)
- [C10 Submission](submission.md)
- [H3 Endpoint](../host/endpoint.md)
