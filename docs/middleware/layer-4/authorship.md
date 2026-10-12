# A2 Authorship

Tells whether a user is the [author](../concepts/author.md) of a [submission](../concepts/submission.md): the person who created it with [Submit](../layer-2/submit.md). Only the author may [Withdraw](../layer-2/withdraw.md) a submission, and only the author sees the data of a [sensitive read](../concepts/sensitive-read.md).

---

**Kind:** atomic part.

**In:** the author of a submission and the current user. **Out:** whether they are the same person.

Someone else trying to [Withdraw](../layer-2/withdraw.md): [Not the author (403)](../concepts/errors.md#not-the-author-403).

---

**Referenced in:**

- [C2 Author](../concepts/author.md)
- [C3 Errors](../concepts/errors.md)
- [I7 Withdraw](../layer-2/withdraw.md)
