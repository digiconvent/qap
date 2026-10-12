# QAP

QAP, the Distributed Trust Authorization Protocol, prevents unilateral actions on sensitive API endpoints. Instead of running straight away, a guarded request becomes a proposal that other people vote on, under a policy that says how many must agree.

All documentation lives in [docs](docs/index.md).

| Folder | Contents |
|---|---|
| [middleware](middleware) | The server side: Go middleware for Fiber v3 that guards routes with QAP. Module `github.com/digiconvent/qap/middleware`. |
| [client](client) | The client side: a TypeScript library, shaped as a state machine, for talking to a guarded route. Package `@digiconvent/qap-client`. |
| [demo](demo) | Middle-earth stories played step by step against the middleware through the client, one lane per person. |

## Development

```sh
go vet ./... && go test -race ./...
cd client && npm ci && npm test && npm run typecheck
cd demo && npm ci && npm start
```
