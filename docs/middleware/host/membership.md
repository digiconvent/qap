# H5 Membership

The host application tells whether a user is a member of a group, right now. QAP uses this to decide [Eligibility](../layer-3/eligibility.md): it asks about each group of a [Policy](../concepts/policy.md). The host keeps its own membership rules to itself, such as when a membership starts and ends, which groups inherit from which, and guests in generated groups.

---

**Kind:** supplied by the host.

**Go:** `func(userId string, groupId string) bool`, named `isMemberOf` in [Setup](setup.md).

**Called with:** a user id and a group id. **Returns:** true when the user is a member of the group at this moment.

---

**Referenced in:**

- [B4 Eligibility](../layer-3/eligibility.md)
- [H7 Setup](setup.md)
