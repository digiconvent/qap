# cA9 Version mismatch

Compares the `X-QAP-Version` header of a [Discover](../layer-2/discover.md) answer with the [Version](../concepts/version.md) of the client. The client only talks to a server of exactly its own version; any other answer becomes [Refused](../states/refused.md).

---

**Kind:** atomic part, `versionMismatch(answer)`.

Nothing when the versions are equal. Otherwise, with the status the server sent:

```ts
{ kind: 'refused', status: 200, message: 'QAP version mismatch: the server speaks 1.1.0, this client 1.0.0', hint: undefined, acceptablePolicies: [], problems: {} }
```

A server that sends no version is named as speaking `no version`.

---

**Referenced in:**

- [cC9 Version](../concepts/version.md)
- [cM1 Discover](../layer-2/discover.md)
- [cS7 Refused](../states/refused.md)
