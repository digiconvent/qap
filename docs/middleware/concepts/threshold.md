# C11 Threshold

A threshold says how many votes it takes to decide a [submission](submission.md). The favour threshold approves it, and the optional against threshold rejects it. A threshold is either a fixed number of votes or a percentage of the people who may vote. [Required count](../layer-4/required-count.md) turns it into an actual number of votes.

---

**Kind:** concept.

A whole number. A positive value is a number of votes; a negative value is a percentage, from `-1` to `-100`. A favour threshold above the number of eligible voters is capped at that number, so a group that shrank can still approve when all its members agree. An against threshold of `0` disables the veto: objections never reject.

| Threshold | Eligible voters | Favour votes needed |
|---|---|---|
| `3` | 5 | 3 |
| `3` | 2 | 2 |
| `-50` | 7 | 4 |
| `-100` | 4 | 4 |

---

**Referenced in:**

- [A9 Required count](../layer-4/required-count.md)
- [B6 Poll evaluation](../layer-3/poll-evaluation.md)
- [B10 Voter count](../layer-3/voter-count.md)
- [C4 Extension](extension.md)
- [C5 Policy](policy.md)
- [C6 Quorum](quorum.md)
- [H4 Member count](../host/member-count.md)
