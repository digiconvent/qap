# A6 Permission

Every protected route names the action it performs as a permission, such as `library.book.delete`. A permission can be a template with placeholders, so one route can cover many items. This part fills each `{name}` placeholder with the route parameter of the same name: `library.book.{id}.delete` on `/api/library/book/123/delete` becomes `library.book.123.delete`. The [result](../concepts/result.md) is what the [Policies](../host/policy-lookup.md) are looked up for.

---

**Kind:** atomic part.

**In:** a permission template and the parameters of the route. **Out:** the resolved permission.

`library.book.{id}.delete` with route parameter `id=123` gives `library.book.123.delete`. A placeholder without a matching route parameter is a mistake in the host application, and the endpoint answers [Misconfigured route (500)](../concepts/errors.md#misconfigured-route-500).

Only [Discover](../layer-2/discover.md) and [Submit](../layer-2/submit.md) resolve the permission. The intents on an existing submission use the permission stored on it, and [List](../layer-2/list.md) goes by the route template instead.

---

**Referenced in:**

- [C3 Errors](../concepts/errors.md)
- [C5 Policy](../concepts/policy.md)
- [C10 Submission](../concepts/submission.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
- [H6 Policy lookup](../host/policy-lookup.md)
- [I1 Discover](../layer-2/discover.md)
- [I2 List](../layer-2/list.md)
