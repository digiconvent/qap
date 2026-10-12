# cA3 Is view

Says whether a successful answer really is the view of a submission. [Read state](../layer-3/read-state.md) and [List](../layer-2/list.md) use it, so an answer from a route without a guard becomes [Not QAP](not-qap.md) instead of a broken state.

---

**Kind:** atomic part, `isView(body)`.

A view passes when `submission` is an object with a string `id`, a `resolution` of `pending`, `executed` or `rejected`, and a `policy` object, and `poll` is an object with numbers for `favour`, `against` and `withhold`, and `comments` as a list or `null`. A resolution this client does not know fails, so it is never read as rejected.

---

**Referenced in:**

- [cB4 Read state](../layer-3/read-state.md)
- [cM2 List](../layer-2/list.md)
