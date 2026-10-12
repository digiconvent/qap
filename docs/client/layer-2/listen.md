# cM6 Listen

Keeps a pending submission up to date without writing a polling loop. It polls now and then and calls the given callback whenever the answer differs from the last one: new votes, a new comment, or the decision. It stops by itself once the submission is no longer [Pending](../states/pending.md), and the function it returns stops it earlier.

---

**Kind:** move.

**Offered by:** [Pending](../states/pending.md) as `listen(onChange, everyMs?)`.

**Sends:** a [Poll](poll.md) every `everyMs` milliseconds, 5000 by default.

**Answers:** a function that stops listening. The callback receives every changed [submission state](../concepts/submission-state.md); an answer equal to the last one is not passed on. A poll that gets no answer at all is retried at the next interval.

```ts
const stop = pending.listen((state) => {
  if (state.kind === 'executed') console.log('carried out', state.result)
})
```

---

**Referenced in:**

- [cS3 Pending](../states/pending.md)
- [cU2 Propose a change](../layer-0/propose-a-change.md)
