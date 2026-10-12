# QAP demo

Ten Middle-earth stories, played step by step against the [middleware](../middleware) through the [client](../client). Each person has a lane showing their current state, the moves that state offers, and the exact HTTP request and response of every step.

```sh
npm ci
npm start
```

Then open `http://localhost:4280`. Set `PORT` to use another port.

| Story | Shows |
|---|---|
| 1. Bilbo leaves Bag End to Frodo | who a policy includes; one vote is enough |
| 2. Boromir pleads to take the Ring to Gondor | rejection by objections |
| 3. The Council decides the Ring must be destroyed | votes adding up, comments, acting on an outdated view |
| 4. The Fellowship is formed | percentage threshold, abstaining |
| 5. Boromir tries to take the Ring | refused outsiders, being offered the right policy |
| 6. Frodo goes on alone | withdrawing, only by the author |
| 7. Gríma gains the king's authority | changing a vote; a decision that changes what someone may do |
| 8. The Entmoot marches on Isengard | quorum |
| 9. Saruman offers peace at Orthanc | a deadline running out |
| 10. Sam reads the Red Book | an answer only the asker sees |

## Layout

| Path | Contents |
|---|---|
| `server/` | A small Go server that guards the stories' routes with the middleware, using only its public setup chain. The person is taken from an `X-User` header. `main.go` and `server.go` hold the routes and the reset; `story/` holds the cast, groups, policies and the Red Book; `realm/` holds what the server supplies to QAP: memberships, the executor and the audit trail. |
| `web/` | The page. `web/src/main.ts` runs it and `types.ts` holds the shared types; `stories/` holds the scenarios and the cast, `play/` performs the moves and records each exchange, and `render/` turns states and exchanges into text. |
| `verify/` | Plays every story against a running server and fails when a step goes differently. |

## Check the stories

With the server running:

```sh
npm run build
npm run verify
```

Set `DEMO_URL` when the server is not on `http://localhost:4280`. One story waits for a 45-second deadline, so a full run takes about a minute.
