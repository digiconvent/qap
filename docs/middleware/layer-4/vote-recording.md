# A14 Vote recording

Records the vote of one person on a [submission](../concepts/submission.md): favour, against or withhold. A person has at most one vote per submission, so a new vote from the same person replaces their earlier vote. The [author](../concepts/author.md) of a submission starts with a favour vote.

---

**Kind:** atomic part.

**In:** the votes of a submission, a user id and a vote. **Out:** the new votes.

A submission keeps three lists of user ids: favour, against and withhold. Recording a vote removes the user from all three lists and adds them to the one they chose.

---

**Referenced in:**

- [B9 Store](../layer-3/store.md)
- [C10 Submission](../concepts/submission.md)
- [I3 Submit](../layer-2/submit.md)
- [I5 Vote](../layer-2/vote.md)
