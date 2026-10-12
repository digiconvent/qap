# cC3 Endpoint options

What opening an [Endpoint](../layer-1/endpoint.md) needs: the path and HTTP method of the guarded route, and the [Send](../app/send.md) function that carries every request. Every move of every state made from it uses the same options.

---

**Kind:** concept, `EndpointOptions`.

| Field | Type | Meaning |
|---|---|---|
| `path` | `string` | The route, as a normal request would use it. Passed to [Send](../app/send.md) as the URL. |
| `method` | `string` | The route's HTTP method, used for every move. |
| `send` | [Send](../app/send.md) | Carries every request. |

---

**Referenced in:**

- [cB1 Ask](../layer-3/ask.md)
- [cE1 Endpoint](../layer-1/endpoint.md)
