# A18 Changes

Compares an entity before and after an approved submission was carried out, field by field, so the audit trail tells who changed which data, and when.

---

**Kind:** atomic part.

**In:** the entity before and after, as the Read of the [Endpoint](../host/endpoint.md) returns them. **Out:** every field that differs, with its value before and after, or nothing when no field changed.

Both are turned into their JSON fields first. A field that appears has `null` before, a field that disappears has `null` after, and an entity that is gone has `null` for every field after. A value that is not an object is compared as a whole, under the field `""`.

```json
{ "title": { "before": "There and Back Again", "after": "The Red Book" } }
```

---

**Referenced in:**

- [A1 Audit summary](audit-summary.md)
- [B1 Carry out](../layer-3/carry-out.md)
- [H3 Endpoint](../host/endpoint.md)
