# B4 Eligibility

A user is eligible under a [Policy](../concepts/policy.md) when [Membership](../host/membership.md) says they are a member of at least one of its groups. Eligibility decides who can use a policy to [Submit](../layer-2/submit.md), who can [Vote](../layer-2/vote.md) and [Comment](../layer-2/comment.md), and who can see submissions through [List](../layer-2/list.md) and [Poll](../layer-2/poll.md). It is checked again on every request, so someone who leaves a group loses their right to vote immediately.

---

**Kind:** built from the host.

**In:** a user id and a policy. **Out:** eligible or not.

Asks [Membership](../host/membership.md) about each group of the policy and stops at the first yes. A request without eligibility gets [Not eligible (403)](../concepts/errors.md#not-eligible-403), or [Policy does not include you (403)](../concepts/errors.md#policy-does-not-include-you-403) on [Submit](../layer-2/submit.md).

---

**Referenced in:**

- [C3 Errors](../concepts/errors.md)
- [C5 Policy](../concepts/policy.md)
- [H5 Membership](../host/membership.md)
- [H7 Setup](../host/setup.md)
- [I1 Discover](../layer-2/discover.md)
- [I2 List](../layer-2/list.md)
- [I3 Submit](../layer-2/submit.md)
- [I4 Poll](../layer-2/poll.md)
- [I5 Vote](../layer-2/vote.md)
- [I6 Comment](../layer-2/comment.md)
