# B11 Withdrawal

Records a [Withdraw](../layer-2/withdraw.md) in the [Audit trail](../host/audit-trail.md), so taking back a proposal leaves the same kind of trace as deciding one.

---

**Kind:** built from atomic parts and the host.

**In:** the withdrawn submission. **Out:** nothing; an [Audit summary](../layer-4/audit-summary.md) with the resolution `withdrawn` and the time of the withdrawal goes to the audit trail. Under an anonymous policy the summary holds as little as for a decision, see [Anonymity](../concepts/anonymity.md).

---

**Referenced in:**

- [I7 Withdraw](../layer-2/withdraw.md)
