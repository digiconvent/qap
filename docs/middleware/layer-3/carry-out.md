# B1 Carry out

Carries out an approved [submission](../concepts/submission.md) through the [Endpoint](../host/endpoint.md) of its route. It checks the stored payload again first, because the world may have changed while people voted; a request that no longer passes is not carried out, and its result says why. When the endpoint is a reader and the route has an id, it reads the entity before and after, so the changes can go to the audit trail.

---

**Kind:** built from atomic parts and the host.

**In:** a submission. **Out:** a [Result](../concepts/result.md), and the changed fields.

1. The stored [Request](../concepts/request.md) is rebuilt, with the author as user and the submission as context.
2. The endpoint checks its payload. [Problems](../layer-4/problems.md) found: the result has status `422` and the body of [Cannot be carried out (422)](../concepts/errors.md#cannot-be-carried-out-422), and nothing else happens.
3. Without a reader or without an `id` in the route: the endpoint executes it, and there are no changes.
4. Otherwise: the endpoint reads the entity, executes, reads it again, and [Changes](../layer-4/changes.md) compares both.

[Execution](execution.md) stores the result, puts the changes into the audit summary, and catches a panic.

---

**Referenced in:**

- [B5 Execution](execution.md)
- [C3 Errors](../concepts/errors.md)
- [C8 Result](../concepts/result.md)
- [H3 Endpoint](../host/endpoint.md)
