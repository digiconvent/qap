# A1 Audit summary

When a [submission](../concepts/submission.md) is decided or withdrawn, this part shapes it into an audit entry: the permission used, its author, the time, the URL, the body that was sent, the outcome or failure, and the votes with stable user ids. The entry stays meaningful even after the users in it are deleted. A poll under an anonymous policy records far less, see [Anonymity](../concepts/anonymity.md). The entry is handed to the [Audit trail](../host/audit-trail.md).

---

**Kind:** atomic part.

**In:** a decided or withdrawn submission. **Out:** an audit entry with:

- the submission id, permission and [policy](../concepts/policy.md)
- the [author](../concepts/author.md)
- the time it was created and the time it was decided
- the URL and the body that were sent
- the resolution: `executed`, `rejected` or `withdrawn`, and the [result](../concepts/result.md) or failure
- the favour, against and withhold lists, as user ids
- the comments, when the policy includes them in audits
- the changed fields with their values before and after, from [Changes](changes.md), when the [Endpoint](../host/endpoint.md) reads its entity

For an anonymous policy, only `anonymous: true`, the author, the time it was decided and the resolution; every other field stays empty.

---

**Referenced in:**

- [B5 Execution](../layer-3/execution.md)
- [B7 Rejection](../layer-3/rejection.md)
- [B11 Withdrawal](../layer-3/withdrawal.md)
- [C1 Anonymity](../concepts/anonymity.md)
- [C5 Policy](../concepts/policy.md)
- [C8 Result](../concepts/result.md)
- [H1 Audit trail](../host/audit-trail.md)
- [H3 Endpoint](../host/endpoint.md)
- [H7 Setup](../host/setup.md)
