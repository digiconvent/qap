# A8 Request copy

A [submission](../concepts/submission.md) stores the request it wants carried out: its method, URL, headers and body, the route template it matched, which [List](../layer-2/list.md) goes by, and the route parameters, which the [Endpoint](../host/endpoint.md) reads. The headers that start with `X-QAP-` belong to the protocol, not to the request, so they are left out. The web framework reuses its memory for the next request, so this part copies those values first. Without the copy, a later request could overwrite a stored submission.

---

**Kind:** atomic part.

**In:** the incoming request. **Out:** its method, its URL including the query string, the route template such as `/api/library/book/:id/delete`, the route parameters by name, its headers without the `X-QAP-` ones, each with all its values, and its body, as copies that the submission owns.

Fiber reuses request buffers between requests, so strings and bytes taken from a request are copied before they are stored.

---

**Referenced in:**

- [C7 Request](../concepts/request.md)
- [C10 Submission](../concepts/submission.md)
- [I3 Submit](../layer-2/submit.md)
