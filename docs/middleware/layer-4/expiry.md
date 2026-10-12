# A4 Expiry

A decided [submission](../concepts/submission.md) stays available while people [Poll](../layer-2/poll.md) it, so everyone with an open form sees the outcome. This part tells whether it may be removed: nobody has polled it for one minute. Deciding a submission counts as polling it, so the minute starts at the decision at the latest. A pending submission never expires.

---

**Kind:** atomic part.

**In:** the resolution of a submission, the time it was last polled, and the current time. **Out:** whether it may be removed.

True when the submission is executed or rejected and `now - last polled >= 60` seconds.

---

**Referenced in:**

- [B2 Cleanup](../layer-3/cleanup.md)
- [B9 Store](../layer-3/store.md)
- [C10 Submission](../concepts/submission.md)
- [I4 Poll](../layer-2/poll.md)
