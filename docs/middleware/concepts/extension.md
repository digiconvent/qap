# C4 Extension

An extension keeps a group able to act when many voters stop taking part. A [policy](policy.md) with a deadline can add extended time: when the [Deadline](../layer-4/deadline.md) passes without the [Threshold](threshold.md) being met, and nobody has voted against, the [submission](submission.md) stays pending for that extra time. When the extension ends and still nobody has voted against, it passes, even though the threshold was never met. A single vote against, before or during the extension, means it is rejected.

---

**Kind:** concept.

Optional, in seconds after the deadline; without a deadline it changes nothing. [Poll evaluation](../layer-3/poll-evaluation.md) applies it, and the [Deadline timer](../layer-3/deadline-timer.md) fires again at its end.

| At the deadline | Then |
|---|---|
| threshold met | executed, as without an extension |
| a vote against | rejected |
| no vote against | pending for the extension; the threshold met during it executes, a vote against during it rejects, and silence until its end executes |

Example: a policy needs half the council in favour within a week, with an extension of three days. Most of the council has gone quiet, and only one member voted in favour. After a week nobody has objected, so the submission waits three more days. Nobody objects in that time either, and it passes.

---

**Referenced in:**

- [A3 Deadline](../layer-4/deadline.md)
- [B3 Deadline timer](../layer-3/deadline-timer.md)
- [B6 Poll evaluation](../layer-3/poll-evaluation.md)
- [C5 Policy](policy.md)
- [I1 Discover](../layer-2/discover.md)
