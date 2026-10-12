# H6 Policy lookup

The host application gives the policies that govern a [Permission](../layer-4/permission.md). A permission can have several policies, for example "one librarian" for small changes and "three librarians" for rare books. Each [Policy](../concepts/policy.md) says which groups may vote, how many votes are needed, and whether there is a deadline.

---

**Kind:** supplied by the host.

**Go:** `func(permission string) []qap.Policy`, named `policyLookup` in [Setup](setup.md).

**Called with:** a resolved [Permission](../layer-4/permission.md). **Returns:** the [Policy](../concepts/policy.md) values of that permission, in the order they are offered.

---

**Referenced in:**

- [A6 Permission](../layer-4/permission.md)
- [C5 Policy](../concepts/policy.md)
- [H7 Setup](setup.md)
- [I1 Discover](../layer-2/discover.md)
