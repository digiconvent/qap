# B6 Poll evaluation

Decides a [submission](../concepts/submission.md), after every vote and when its [Deadline timer](deadline-timer.md) fires. The submission leads to [Rejection](rejection.md) when its [Deadline](../layer-4/deadline.md) has passed or when its against votes reach their [Required count](../layer-4/required-count.md). It leads to [Execution](execution.md) when the favour votes in its [Vote count](../layer-4/vote-count.md) reach their [Required count](../layer-4/required-count.md) and enough people took part. Otherwise it stays pending. [Thresholds](../concepts/threshold.md) given as a percentage use the [Voter count](voter-count.md).

---

**Kind:** built from atomic parts.

**In:** a submission, its [policy](../concepts/policy.md) settings, the [Voter count](voter-count.md) stored on it and the current time. **Out:** pending, executed or rejected.

The checks run in this order:

1. Against votes reach the against [Required count](../layer-4/required-count.md): rejected. An against threshold of `0` disables this step.
2. The quorum is met (participation reaches the quorum [Required count](../layer-4/required-count.md), or there is no quorum) and favour votes reach the favour [Required count](../layer-4/required-count.md), capped at the voter count: executed.
3. [Deadline](../layer-4/deadline.md) passed: rejected, unless the policy has an [Extension](../concepts/extension.md) and nobody voted against. Then: executed once the extension has passed, otherwise pending. A submission that meets its criteria at the deadline or during the extension is executed by step 2 first.
4. Otherwise: pending.

---

**Referenced in:**

- [A13 Vote count](../layer-4/vote-count.md)
- [B3 Deadline timer](deadline-timer.md)
- [B10 Voter count](voter-count.md)
- [C4 Extension](../concepts/extension.md)
- [C6 Quorum](../concepts/quorum.md)
- [C10 Submission](../concepts/submission.md)
- [I3 Submit](../layer-2/submit.md)
- [I5 Vote](../layer-2/vote.md)
