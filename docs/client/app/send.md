# cH1 Send

The function every request goes through. It is the only place the client meets the outside world, so it is where you add a login token, cookies or a base URL. The browser's `fetch` can be passed as it is. [Ask](../layer-3/ask.md) calls it.

---

**Kind:** supplied by the app, `Send`.

```ts
type Send = (url: string, init: RequestInit) => Promise<Response>
```

- It is called on its own, not as a method, so `fetch` works without binding.
- It must keep the headers in `init`.
- Errors it throws, such as no network, DNS or CORS, are not caught; they reach the caller of the move. Only an answer becomes a state.

---

**Referenced in:**

- [cB1 Ask](../layer-3/ask.md)
- [cC3 Endpoint options](../concepts/endpoint-options.md)
- [cH2 Setup](setup.md)
- [cS7 Refused](../states/refused.md)
- [cU2 Propose a change](../layer-0/propose-a-change.md)
