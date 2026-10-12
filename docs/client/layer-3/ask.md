# cB1 Ask

Sends one request through [Send](../app/send.md) and reads the answer into a status, its headers and a body. Every move uses it.

---

**Kind:** built from atomic parts and the app, `ask(options, headers, body?)`.

- Uses the method and path from the [Endpoint options](../concepts/endpoint-options.md).
- Leaves the body out when it is empty, so a `GET` can be proposed.
- Calls `send` on its own, not as a method.
- Returns `{ status, headers, body }`, the body read by [Read body](../layer-4/read-body.md).

---

**Referenced in:**

- [cA5 Read body](../layer-4/read-body.md)
- [cH1 Send](../app/send.md)
