# cC10 Vote choice

What [Vote](../layer-2/vote.md) takes: in favour, against, or abstain. [Vote value](../layer-4/vote-value.md) turns it into the number the server expects.

---

**Kind:** concept, `VoteChoice`.

```ts
type VoteChoice = 'favour' | 'against' | 'withhold'
```

`'withhold'` is refused with 400 unless the policy has `withholdPermitted`.

---

**Referenced in:**

- [cA10 Vote value](../layer-4/vote-value.md)
- [cM7 Vote](../layer-2/vote.md)
- [cS3 Pending](../states/pending.md)
