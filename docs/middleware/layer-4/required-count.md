# A9 Required count

A [threshold](../concepts/threshold.md) says how many votes are needed. This part turns it into an actual number of votes. A positive threshold is that number of votes. A negative threshold is a percentage of the eligible voters, rounded up, so `-50` with 5 eligible voters needs 3 votes. The number of eligible voters comes from the [Voter count](../layer-3/voter-count.md).

---

**Kind:** atomic part.

**In:** a threshold (whole number) and the number of eligible voters. **Out:** the number of votes needed.

- Positive threshold `t`: `t` votes. For against votes and quorum, a positive threshold larger than the number of eligible voters cannot be met.
- Negative threshold `t`: `ceil(abs(t) / 100 * eligible voters)` votes.
- For favour votes the number is capped at the number of eligible voters, so a favour threshold of 3 with 2 eligible voters needs 2.

---

**Referenced in:**

- [A13 Vote count](vote-count.md)
- [B6 Poll evaluation](../layer-3/poll-evaluation.md)
- [B10 Voter count](../layer-3/voter-count.md)
- [C6 Quorum](../concepts/quorum.md)
- [C11 Threshold](../concepts/threshold.md)
