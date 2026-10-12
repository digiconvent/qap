# cH2 Setup

The app installs the client once, and opens an [Endpoint](../layer-1/endpoint.md) for every guarded route it talks to, with the path and method of the route and a [Send](send.md) function that identifies the user. In a component, it keeps the state that each move returns.

---

**Kind:** done by the app.

The package is not published to a registry yet. Build it from the `client` folder of the QAP repository, and depend on it by path:

```sh
cd qap/client && npm install && npm run build
cd your-app && npm install ../path/to/qap/client
```

Use the client of the same release as the middleware of the server, such as the tag `v1.0.0` for both: the client refuses a server of another [Version](../concepts/version.md).

It is an ES module only, with its TypeScript types included: use `"type": "module"` or a bundler, and `moduleResolution` set to `"Bundler"` or `"NodeNext"`. It needs a `fetch`-compatible function: any browser, or Node 18 and later.

```ts
import { openEndpoint, type Send } from '@digiconvent/qap-client'

declare function currentToken(): string

const sendWithToken: Send = (url, init) => {
  const headers = new Headers(init.headers)
  headers.set('Authorization', `Bearer ${currentToken()}`)
  return fetch(url, { ...init, headers })
}

const endpoint = openEndpoint({ path: '/api/council/ring/decide', method: 'POST', send: sendWithToken })
```

The path and method are the ones a normal request to the route would use. The token is read inside `send`, so a refreshed token is picked up. For cookies, `send` passes `credentials: 'include'`; in Node, or for an API on another host, it puts a base URL in front of the path.

### In a React component

Each state is an immutable object with its moves attached. The component keeps the whole state, replaces it with the state every move returns, and renders by `kind`.

```tsx
import { useState } from 'react'
import { openEndpoint, type SubmissionState } from '@digiconvent/qap-client'

const endpoint = openEndpoint({ path: '/api/council/ring/decide', method: 'POST', send: fetch })

export function Proposal({ id }: { id: string }) {
  const [state, setState] = useState<SubmissionState | null>(null)

  async function open() {
    setState(await endpoint.follow(id))
  }

  async function voteFor() {
    if (state?.kind === 'pending') setState(await state.vote('favour'))
  }

  async function voteAgainst() {
    if (state?.kind === 'pending') setState(await state.vote('against'))
  }

  if (!state) return <button onClick={open}>open</button>

  switch (state.kind) {
    case 'pending':
      return (
        <div>
          <p>{state.counts.favour} for, {state.counts.against} against</p>
          <button onClick={voteFor}>vote for</button>
          <button onClick={voteAgainst}>vote against</button>
        </div>
      )
    case 'executed':
      return <p>carried out</p>
    case 'rejected':
      return <p>rejected</p>
    case 'withdrawn':
      return <p>withdrawn</p>
    case 'refused':
      return <p>{state.status}: {state.message}</p>
  }
}
```

- The endpoint is opened once per route, outside the component.
- Buttons are disabled while a move runs: a second click acts on the same state twice.
- A state is never serialised into a store, because its moves are functions and get lost. The store keeps the `id`, and [Follow](../layer-2/follow.md) opens it again.
