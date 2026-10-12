# cS7 Refused

The server said no. Every move can end here. It carries the HTTP status, the message, an optional hint, after a [Submit](../layer-2/submit.md) without a policy the policies you could use, and, when the endpoint cannot carry the request out, the problems per field. It has no moves: go back to the [Endpoint](../layer-1/endpoint.md). [Read refusal](../layer-4/read-refusal.md) builds it from any answer that is not a success.

---

**Kind:** state, `kind: 'refused'`.

| Field | Type |
|---|---|
| `status` | `number` |
| `message` | `string` |
| `hint` | `string \| undefined` |
| `acceptablePolicies` | [Policy reference](../concepts/policy-reference.md)`[]` |
| `problems` | `Record<string, string[]>`, empty unless the status is 422 |

| Status | When |
|---|---|
| 400 | The id in the URL is not a UUID, abstaining is not allowed by the policy, or a comment is not readable or longer than 500 characters. |
| 401 | [Send](../app/send.md) did not identify a user. |
| 403 | The policy does not include you, or you withdrew someone else's submission. |
| 404 | No such submission: never existed, withdrawn, or cleaned up. |
| 406 | The policy was missing when submitting; see `acceptablePolicies`. |
| 409 | The submission was already decided. |
| 422 | The endpoint of the route cannot carry the request out, or its body cannot be read; see `problems`. |
| 500 | The route is set up wrongly on the server. |
| 2xx | The answer was a success but not a QAP answer, for example from a route without a guard, with the message `not a QAP answer`; or the server speaks another [Version](../concepts/version.md), see [Version mismatch](../layer-4/version-mismatch.md). |

A request that gets no answer at all does not become a refused state: the error from [Send](../app/send.md) is thrown.

```ts
{ kind: 'refused', status: 422, message: 'the request cannot be carried out', hint: undefined,
  acceptablePolicies: [], problems: { fate: ['must say what happens to the Ring'] } }
```

---

**Referenced in:**

- [cA2 Is success](../layer-4/is-success.md)
- [cA4 Not QAP](../layer-4/not-qap.md)
- [cA8 Read refusal](../layer-4/read-refusal.md)
- [cA9 Version mismatch](../layer-4/version-mismatch.md)
- [cC4 Policy reference](../concepts/policy-reference.md)
- [cC8 Submission state](../concepts/submission-state.md)
- [cE1 Endpoint](../layer-1/endpoint.md)
- [cM1 Discover](../layer-2/discover.md)
- [cM2 List](../layer-2/list.md)
- [cM3 Submit](../layer-2/submit.md)
- [cM5 Poll](../layer-2/poll.md)
- [cM7 Vote](../layer-2/vote.md)
- [cM8 Comment](../layer-2/comment.md)
- [cM9 Withdraw](../layer-2/withdraw.md)
- [cS1 Discovered](discovered.md)
- [cS3 Pending](pending.md)
- [cU1 Approve or reject](../layer-0/approve-or-reject.md)
- [cU2 Propose a change](../layer-0/propose-a-change.md)
- [cU4 Take back a proposal](../layer-0/take-back-a-proposal.md)
