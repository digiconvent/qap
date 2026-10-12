# cS3 Pending

A submission that waits for votes. It carries the [Submission facts](../concepts/submission-facts.md) as they were at the last answer, and offers [Poll](../layer-2/poll.md), [Listen](../layer-2/listen.md), [Vote](../layer-2/vote.md), [Comment](../layer-2/comment.md) and [Withdraw](../layer-2/withdraw.md). Other people keep voting, so the facts can be out of date: a move on an old pending state can come back as [Refused](refused.md) with status 409.

---

**Kind:** state, `kind: 'pending'`, built by [Pending state](../layer-3/pending-state.md).

**Fields:** the [Submission facts](../concepts/submission-facts.md).

| Move | Leads to |
|---|---|
| `poll()` | any [submission state](../concepts/submission-state.md) |
| `listen(onChange, everyMs?)` | a function that stops listening; `onChange` receives every changed [submission state](../concepts/submission-state.md) |
| `vote(choice)` | any [submission state](../concepts/submission-state.md), with a [Vote choice](../concepts/vote-choice.md) |
| `comment(text)` | any [submission state](../concepts/submission-state.md) |
| `withdraw()` | [Withdrawn](withdrawn.md) or [Refused](refused.md) |

`withdraw` is offered to everyone; only the author succeeds. The server decides, because the client does not know who is asking.

---

**Referenced in:**

- [cB3 Pending state](../layer-3/pending-state.md)
- [cC8 Submission state](../concepts/submission-state.md)
- [cE1 Endpoint](../layer-1/endpoint.md)
- [cM2 List](../layer-2/list.md)
- [cM3 Submit](../layer-2/submit.md)
- [cM5 Poll](../layer-2/poll.md)
- [cM6 Listen](../layer-2/listen.md)
- [cM7 Vote](../layer-2/vote.md)
- [cM8 Comment](../layer-2/comment.md)
- [cM9 Withdraw](../layer-2/withdraw.md)
- [cS1 Discovered](discovered.md)
- [cS2 Listed](listed.md)
- [cU1 Approve or reject](../layer-0/approve-or-reject.md)
- [cU2 Propose a change](../layer-0/propose-a-change.md)
- [cU3 Review pending work](../layer-0/review-pending-work.md)
