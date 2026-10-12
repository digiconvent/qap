# A17 Write result

Writes the [Result](../concepts/result.md) of an [Endpoint](../host/endpoint.md) as the answer to a plain request: its status, its headers and its body.

---

**Kind:** atomic part.

**In:** the request being answered and a result. **Out:** the answer.

| Body | Written as |
|---|---|
| `nil` | no body |
| a string | that text |
| bytes | those bytes |
| anything else | JSON |

---

**Referenced in:**

- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
