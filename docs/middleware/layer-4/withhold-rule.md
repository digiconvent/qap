# A16 Withhold rule

A [policy](../concepts/policy.md) decides whether people may withhold their vote, which means taking no side. This part tells whether the policy of a [submission](../concepts/submission.md) allows it. When it does not, a withhold [Vote](../layer-2/vote.md) is refused.

---

**Kind:** atomic part.

**In:** the policy of a submission. **Out:** whether withhold votes are allowed.

A refused withhold vote: [Withhold not permitted (400)](../concepts/errors.md#withhold-not-permitted-400).

---

**Referenced in:**

- [C3 Errors](../concepts/errors.md)
- [C5 Policy](../concepts/policy.md)
- [I5 Vote](../layer-2/vote.md)
