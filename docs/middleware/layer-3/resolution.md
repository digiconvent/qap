# B8 Resolution

Settles a pending [submission](../concepts/submission.md) as executed or rejected. It happens exactly once: when two deciding votes arrive at the same moment, only one of them succeeds, so a request is never carried out twice. It also only settles the votes that were evaluated: when another vote arrived in between, the outcome may be wrong, so it stands down and that newer vote decides instead. After resolution, a submission never changes again.

---

**Kind:** built from atomic parts.

**In:** a submission id, the revision that was evaluated, and an outcome. **Out:** the settled submission, "already resolved", or "outdated".

Every vote raises the revision of a submission. The [store](store.md) changes `pending` to the outcome in one step, and only when the submission is still pending and still at the evaluated revision. An outdated submission stays pending. An intent on a submission that is already resolved: [Already resolved (409)](../concepts/errors.md#already-resolved-409).

---

**Referenced in:**

- [B3 Deadline timer](deadline-timer.md)
- [B5 Execution](execution.md)
- [B7 Rejection](rejection.md)
- [B9 Store](store.md)
- [C3 Errors](../concepts/errors.md)
- [C10 Submission](../concepts/submission.md)
