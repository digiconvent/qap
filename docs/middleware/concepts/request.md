# C7 Request

A request is what the Execute of an [Endpoint](../host/endpoint.md) carries out. For a plain request it is the incoming request; for a [submission](submission.md) it is the stored copy, made by [Request copy](../layer-4/request-copy.md) when it was submitted, so it can be carried out long after the original request is gone.

---

**Kind:** concept, `qap.Request`.

| Field | Meaning |
|---|---|
| method | The HTTP method |
| url | The URL, including the query string |
| route | The route template it matched, such as `/api/library/book/:id/delete` |
| params | The route parameters by name, such as `id` → a UUID |
| headers | The headers, without the `X-QAP-` ones |
| body | The body, as it was sent |
| user | Who asks: the current user for a plain request, the [author](author.md) for a submission |
| submission | `nil` for a plain request; for an approved submission, the [Submission](submission.md) with its id, resolved permission, policy, votes and comments |

The original Fiber context is not part of it: by the time a submission is approved, that request is long gone, and Fiber reuses its context for other requests.

---

**Referenced in:**

- [B1 Carry out](../layer-3/carry-out.md)
- [C10 Submission](submission.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
- [H3 Endpoint](../host/endpoint.md)
