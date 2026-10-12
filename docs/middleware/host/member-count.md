# H4 Member count

The host application tells how many people are members of a set of groups, right now. Someone who is a member of several of the groups counts once. QAP uses it for the [Voter count](../layer-3/voter-count.md), which turns a [Threshold](../concepts/threshold.md) or [Quorum](../concepts/quorum.md) given as a percentage into a number of votes. QAP asks once per submission, when it is created.

---

**Kind:** supplied by the host.

**Go:** `func(groupIds []string) int`, named `countMembersOf` in [Setup](setup.md).

**Called with:** the group ids of a [Policy](../concepts/policy.md). **Returns:** the number of distinct members.

---

**Referenced in:**

- [B10 Voter count](../layer-3/voter-count.md)
- [H7 Setup](setup.md)
