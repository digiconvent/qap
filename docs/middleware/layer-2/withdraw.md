# I7 Withdraw

The delete-[submission](../concepts/submission.md) intent. The [author](../concepts/author.md) takes back a submission that is still pending, for example after a mistake. [Authorship](../layer-4/authorship.md) confirms that the [Current user](../host/current-user.md) wrote it, the submission is removed, its [Deadline timer](../layer-3/deadline-timer.md) is cancelled, and [Withdrawal](../layer-3/withdrawal.md) records it in the audit trail. A decided submission cannot be withdrawn.

---

**Kind:** intent, `X-QAP-Intent: delete-submission`.

**Headers:** `X-QAP-ID`: the submission.

**Response:** `200` with an empty body. The submission no longer exists, so a later [Poll](poll.md) on its id answers [Unknown submission (404)](../concepts/errors.md#unknown-submission-404).

**Errors:** [Unauthorized (401)](../concepts/errors.md#unauthorized-401), [Missing submission id (406)](../concepts/errors.md#missing-submission-id-406), [Unknown submission (404)](../concepts/errors.md#unknown-submission-404), [Not the author (403)](../concepts/errors.md#not-the-author-403), [Already resolved (409)](../concepts/errors.md#already-resolved-409).

---

**Referenced in:**

- [A2 Authorship](../layer-4/authorship.md)
- [A10 Submission id](../layer-4/submission-id.md)
- [B3 Deadline timer](../layer-3/deadline-timer.md)
- [B9 Store](../layer-3/store.md)
- [B11 Withdrawal](../layer-3/withdrawal.md)
- [C2 Author](../concepts/author.md)
- [C10 Submission](../concepts/submission.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
- [U4 Take back a proposal](../layer-0/take-back-a-proposal.md)
