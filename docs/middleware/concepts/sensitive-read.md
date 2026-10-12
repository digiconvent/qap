# C9 Sensitive read

Some data may only be read after approval, such as the borrowing history of a patron. A person asks for it with [Submit](../layer-2/submit.md) on the `GET` route, and others [Vote](../layer-2/vote.md). Once it is approved, the data is stored in the [Result](result.md), and the [Author](author.md) reads it there with [Poll](../layer-2/poll.md). Voters never see the data. Data that is not sensitive is read without QAP.

---

**Kind:** concept.

A `GET` [submission](submission.md) is carried out by the [Endpoint](../host/endpoint.md) of its route like any other submission. The [Submission view](../layer-4/submission-view.md) shows `result.body` only to the author. There is no second request to fetch the data.

---

**Referenced in:**

- [A2 Authorship](../layer-4/authorship.md)
- [C2 Author](author.md)
- [C8 Result](result.md)
- [I4 Poll](../layer-2/poll.md)
