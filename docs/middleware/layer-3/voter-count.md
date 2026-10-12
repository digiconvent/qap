# B10 Voter count

The number of people who are eligible under a [Policy](../concepts/policy.md) when a [submission](../concepts/submission.md) is created: the [Member count](../host/member-count.md) of all its groups, each person counted once. A [Threshold](../concepts/threshold.md) or [Quorum](../concepts/quorum.md) given as a percentage needs it, so that "half of the librarians" becomes an actual number of votes. [Submit](../layer-2/submit.md) stores it on the submission, like the policy itself, so the host is asked once per submission and the number does not change while people vote.

---

**Kind:** built from the host.

**In:** a policy. **Out:** a whole number.

Calls [Member count](../host/member-count.md) with the group ids of the policy. [Required count](../layer-4/required-count.md) uses the stored result on every later [Poll evaluation](poll-evaluation.md), including when the [Deadline timer](deadline-timer.md) fires.

---

**Referenced in:**

- [A9 Required count](../layer-4/required-count.md)
- [B6 Poll evaluation](poll-evaluation.md)
- [C10 Submission](../concepts/submission.md)
- [H4 Member count](../host/member-count.md)
