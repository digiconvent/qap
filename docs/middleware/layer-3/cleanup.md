# B2 Cleanup

Cleanup removes decided [submissions](../concepts/submission.md) that nobody looks at anymore, so the [Store](store.md) does not grow forever. A submission goes once [Expiry](../layer-4/expiry.md) says that nobody has polled it for one minute. Pending submissions are never removed.

---

**Kind:** built from atomic parts.

Runs every minute and calls "remove expired" on the [Store](store.md). A [Poll](../layer-2/poll.md) on a removed submission answers [Unknown submission (404)](../concepts/errors.md#unknown-submission-404).

---

**Referenced in:**

- [B9 Store](store.md)
- [C3 Errors](../concepts/errors.md)
