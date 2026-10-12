# cC8 Submission state

Any state a move on a submission can lead to: [Pending](../states/pending.md), [Executed](../states/executed.md), [Rejected](../states/rejected.md), [Withdrawn](../states/withdrawn.md) or [Refused](../states/refused.md). Check `kind` to find out which.

---

**Kind:** concept, `SubmissionState`.

```ts
type SubmissionState = Pending | Executed | Rejected | Withdrawn | Refused
```

---

**Referenced in:**

- [cB4 Read state](../layer-3/read-state.md)
- [cM3 Submit](../layer-2/submit.md)
- [cM5 Poll](../layer-2/poll.md)
- [cM6 Listen](../layer-2/listen.md)
- [cM7 Vote](../layer-2/vote.md)
- [cM8 Comment](../layer-2/comment.md)
- [cS3 Pending](../states/pending.md)
- [cS4 Executed](../states/executed.md)
- [cS5 Rejected](../states/rejected.md)
