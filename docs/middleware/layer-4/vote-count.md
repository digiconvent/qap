# A13 Vote count

Counts the favour, against and withhold votes on a [submission](../concepts/submission.md). It also counts participation, which is favour plus against: a withhold vote does not count as taking part. [Poll evaluation](../layer-3/poll-evaluation.md) compares these numbers with each [Required count](required-count.md).

---

**Kind:** atomic part.

**In:** the votes of a submission. **Out:** the number of favour, against and withhold votes, and participation (favour + against).

---

**Referenced in:**

- [B6 Poll evaluation](../layer-3/poll-evaluation.md)
- [C6 Quorum](../concepts/quorum.md)
