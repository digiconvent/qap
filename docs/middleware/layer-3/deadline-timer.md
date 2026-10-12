# B3 Deadline timer

Starts a timer when a [submission](../concepts/submission.md) is created under a [policy](../concepts/policy.md) with a [Deadline](../layer-4/deadline.md). When the timer fires, [Poll evaluation](poll-evaluation.md) decides the submission. A submission still pending in an [Extension](../concepts/extension.md) gets a new timer for the end of that extension; any other submission that is still pending leads to [Rejection](rejection.md). A submission that is decided or withdrawn before the deadline cancels its timer.

---

**Kind:** built from atomic parts.

Started by [Submit](../layer-2/submit.md) when the policy has a deadline, with a delay of that many seconds. Cancelled when [Submit](../layer-2/submit.md) or [Vote](../layer-2/vote.md) has actually decided the submission through [Resolution](resolution.md), and by [Withdraw](../layer-2/withdraw.md). A decision that stood down as outdated leaves the timer running. When the timer fires at the moment a vote arrives, whichever decides first wins, and the other finds the submission resolved. Timers live in memory, like everything else in QAP, so a restart loses them together with the submissions.

---

**Referenced in:**

- [A3 Deadline](../layer-4/deadline.md)
- [B6 Poll evaluation](poll-evaluation.md)
- [B9 Store](store.md)
- [B10 Voter count](voter-count.md)
- [C4 Extension](../concepts/extension.md)
- [C5 Policy](../concepts/policy.md)
- [I3 Submit](../layer-2/submit.md)
- [I7 Withdraw](../layer-2/withdraw.md)
