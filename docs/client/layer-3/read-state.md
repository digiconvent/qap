# cB4 Read state

Turns an answer about one submission into a [submission state](../concepts/submission-state.md). An answer that is not a success becomes [Read refusal](../layer-4/read-refusal.md). A success that [Is view](../layer-4/is-view.md) rejects becomes [Not QAP](../layer-4/not-qap.md). Every other answer goes to [Read view](read-view.md).

---

**Kind:** built from atomic parts, `readState(options, answer)`.

Whether the answer is a success is decided by [Is success](../layer-4/is-success.md).

---

**Referenced in:**

- [cA2 Is success](../layer-4/is-success.md)
- [cA3 Is view](../layer-4/is-view.md)
- [cM3 Submit](../layer-2/submit.md)
- [cM5 Poll](../layer-2/poll.md)
- [cM7 Vote](../layer-2/vote.md)
- [cM8 Comment](../layer-2/comment.md)
