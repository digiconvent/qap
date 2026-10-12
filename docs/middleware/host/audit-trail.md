# H1 Audit trail

The host application receives every [Audit summary](../layer-4/audit-summary.md) and keeps it, for example in a database. QAP only writes to it. Reading the audit trail back is up to the host.

---

**Kind:** supplied by the host.

**Go:** `a value with the method Record(summary qap.AuditSummary)`. Handed to QAP during [Setup](setup.md).

**Called with:** an [Audit summary](../layer-4/audit-summary.md). **Returns:** nothing that QAP waits for. An error while writing does not stop or undo the decision.

---

**Referenced in:**

- [A1 Audit summary](../layer-4/audit-summary.md)
- [B5 Execution](../layer-3/execution.md)
- [B7 Rejection](../layer-3/rejection.md)
- [B11 Withdrawal](../layer-3/withdrawal.md)
- [H7 Setup](setup.md)
