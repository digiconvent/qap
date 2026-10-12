# cA1 Is discovery

Says whether a successful answer to [Discover](../layer-2/discover.md) really lists policies. When it does not, the answer becomes [Not QAP](not-qap.md).

---

**Kind:** atomic part, `isDiscovery(body)`.

A discovery passes when `permission` is a string and `policies` is a list in which every policy has a string `id` and `name` and a number `favour_threshold`.

---

**Referenced in:**

- [cM1 Discover](../layer-2/discover.md)
