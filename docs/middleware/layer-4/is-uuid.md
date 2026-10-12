# A19 Is UUID

Ids are always UUIDs. Every route parameter called `id` must be one, on every request to the route, whatever the protection and the intent; anything else is refused before QAP does anything else.

---

**Kind:** atomic part.

**In:** the `id` route parameter. **Out:** whether it is a UUID: 32 hexadecimal digits in groups of 8, 4, 4, 4 and 12, separated by dashes, in either case.

Not a UUID: [Invalid id (400)](../concepts/errors.md#invalid-id-400).

---

**Referenced in:**

- [C3 Errors](../concepts/errors.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
- [H3 Endpoint](../host/endpoint.md)
