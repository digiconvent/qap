# C1 Anonymity

A [policy](policy.md) can make its polls anonymous. Voters never see who voted how, in any poll: the [Submission view](../layer-4/submission-view.md) only shows counts. What changes is the record that stays behind: the [Audit summary](../layer-4/audit-summary.md) of an anonymous poll only holds the [author](author.md), the time of the decision and the outcome. Only a poll that is not anonymous records who voted in favour, against and withhold.

---

**Kind:** concept.

A flag on the policy, off by default.

| Audit summary of | Holds |
|---|---|
| an anonymous poll | `anonymous: true`, the author, the time of the decision and the resolution |
| any other poll | everything, including the favour, against and withhold lists, the changed fields and, when the policy includes them, the comments |

---

**Referenced in:**

- [A1 Audit summary](../layer-4/audit-summary.md)
- [B11 Withdrawal](../layer-3/withdrawal.md)
- [C5 Policy](policy.md)
