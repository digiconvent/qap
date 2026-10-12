# cA6 Read facts

Turns the server's view of a submission into [Submission facts](../concepts/submission-facts.md).

---

**Kind:** atomic part, `readFacts(view)`.

The view has a `submission` and a `poll` part. The counts and comments come from `poll`, everything else from `submission`. Missing comments become an empty list.

---

**Referenced in:**

- [cB2 Decided state](../layer-3/decided-state.md)
- [cB3 Pending state](../layer-3/pending-state.md)
- [cC7 Submission facts](../concepts/submission-facts.md)
