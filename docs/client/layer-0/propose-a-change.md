# cU2 Propose a change

A person wants to do something that needs approval, such as deciding the fate of the Ring. They use [Discover](../layer-2/discover.md) to see which policies they may use, pick one and [Submit](../layer-2/submit.md) the request. Then they share its id with colleagues and [Listen](../layer-2/listen.md) until it is decided.

---

**Kind:** user feature.

**Moves a client makes:**

```ts
import { openEndpoint } from '@digiconvent/qap-client'

const endpoint = openEndpoint({ path: '/api/council/ring/decide', method: 'POST', send: fetch })

const discovered = await endpoint.discover()
if (discovered.kind === 'discovered' && discovered.policies.length > 0) {
  const body = JSON.stringify({ ring: 'the One Ring', fate: 'destroy it in the fires of Mount Doom' })
  const proposed = await discovered.submit(discovered.policies[0], body)
  if (proposed.kind === 'pending') {
    proposed.listen((state) => console.log(state.kind))
  }
}
```

The answer to the submit is [Executed](../states/executed.md) when one vote is enough, [Pending](../states/pending.md) when it waits for votes, or [Refused](../states/refused.md). A submit without a policy is refused with 406, and `acceptablePolicies` lists the ones to submit again with. A request the route cannot carry out is refused with 422, and `problems` says what is wrong with it.

The body is a string, left out for a request without one, such as a `GET`. The client adds no `Content-Type`; [Send](../app/send.md) adds one when the route needs it.

The server does not push changes, so listening polls in the background and tells the person about every change, until the decision. A policy with `deadlineSeconds` decides by itself when the deadline passes. A decided submission is removed a minute after the last time anyone polled it, so listening stops once it is decided. For a proposed `GET`, the [Result](../concepts/result.md) holds the data for the author only.
