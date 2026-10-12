# C5 Policy

A policy is a set of approval rules for a [Permission](../layer-4/permission.md). It says who may vote, through its groups and [Eligibility](../layer-3/eligibility.md); how many votes approve or reject a [submission](submission.md), through its [Threshold](threshold.md)s and [Quorum](quorum.md); whether a submission has a [Deadline](../layer-4/deadline.md); and whether people may withhold their vote. A permission can have several policies, and the [Author](author.md) picks one when they [Submit](../layer-2/submit.md). The host supplies them through [Policies](../host/policy-lookup.md). A policy applies equally to everyone it is linked to: there are no personal policies. A permission can still contain an id, such as library `123` in `library.123.add-book`, so that one item gets its own policies.

---

**Kind:** concept.

| Field | Meaning |
|---|---|
| id, name | Identify the policy |
| groups | Decide [Eligibility](../layer-3/eligibility.md) |
| favour threshold | The [Threshold](threshold.md) of votes in favour that approves a submission |
| against threshold | The [Threshold](threshold.md) of votes against that rejects a submission (optional) |
| quorum | The [Quorum](quorum.md) needed before approval (optional) |
| deadline | Seconds until the [Deadline timer](../layer-3/deadline-timer.md) fires (optional) |
| extension | Seconds of extended time after the deadline, see [Extension](extension.md) (optional) |
| withhold permitted | See [Withhold rule](../layer-4/withhold-rule.md) |
| comments in audit | Whether the [Audit summary](../layer-4/audit-summary.md) includes comments |
| anonymous | Whether the poll is anonymous, see [Anonymity](anonymity.md) |

A pending submission keeps the settings of the policy it was created under. A changed policy only applies to new submissions.

---

**Referenced in:**

- [A1 Audit summary](../layer-4/audit-summary.md)
- [A3 Deadline](../layer-4/deadline.md)
- [A11 Submission view](../layer-4/submission-view.md)
- [A16 Withhold rule](../layer-4/withhold-rule.md)
- [B3 Deadline timer](../layer-3/deadline-timer.md)
- [B4 Eligibility](../layer-3/eligibility.md)
- [B6 Poll evaluation](../layer-3/poll-evaluation.md)
- [B10 Voter count](../layer-3/voter-count.md)
- [C1 Anonymity](anonymity.md)
- [C3 Errors](errors.md)
- [C4 Extension](extension.md)
- [C10 Submission](submission.md)
- [H4 Member count](../host/member-count.md)
- [H5 Membership](../host/membership.md)
- [H6 Policy lookup](../host/policy-lookup.md)
- [I3 Submit](../layer-2/submit.md)
