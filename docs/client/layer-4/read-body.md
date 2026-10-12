# cA5 Read body

Reads the body of a response for [Ask](../layer-3/ask.md): parsed JSON, `null` when empty, and the plain text wrapped as an error when it is not JSON.

---

**Kind:** atomic part, `readBody(response)`.

| Body | Result |
|---|---|
| empty | `null` |
| JSON | the parsed value |
| other text | `{ error: text }` |

---

**Referenced in:**

- [cB1 Ask](../layer-3/ask.md)
