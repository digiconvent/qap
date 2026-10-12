# C3 Errors

An error is the answer QAP gives when it refuses a request. Every error has an HTTP status and a JSON body with an `error` message, sometimes with a `hint` that says how to fix the request. This page lists every error QAP can return and where it comes from.

---

**Kind:** concept.

Body shape: `{"error": "<message>", "hint": "<how to fix it>"}`. `hint` is optional.

### Unauthorized (401)

`{"error": "unauthorized"}`. Nobody is logged in: the [Current user](../host/current-user.md) is empty.

### Missing intent (406)

`{"error": "QAP intent header required"}`. The request has no `X-QAP-Intent` header (see [Intent](../layer-4/intent.md)).

### Unknown intent (400)

`{"error": "Unknown intent: '<value>'"}`. The `X-QAP-Intent` header names no known [Intent](../layer-4/intent.md).

### Missing submission id (406)

`{"error": "QAP id header required", "hint": "Specify X-QAP-ID"}`. An intent that works on one [submission](submission.md) has no `X-QAP-ID` header.

### Missing policy id (406)

`{"error": "QAP policy id header required", "hint": "Specify X-QAP-Policy-ID", "acceptable_policies": [...]}`. A [Submit](../layer-2/submit.md) without `X-QAP-Policy-ID`. `acceptable_policies` lists the policies the user may pick.

### Policy does not include you (403)

`{"error": "Policy does not include you"}`. A [Submit](../layer-2/submit.md) with a [Policy](policy.md) the user has no [Eligibility](../layer-3/eligibility.md) for.

### Misconfigured route (500)

`{"error": "misconfigured permission template"}`. A [Permission](../layer-4/permission.md) placeholder has no matching route parameter. A mistake in the host application.

### Unknown submission (404)

`{"error": "Submission state does not exist"}`. No submission has this id, or it was withdrawn or removed by [Cleanup](../layer-3/cleanup.md).

### Already resolved (409)

`{"error": "Submission is already resolved"}`. The submission was decided by [Resolution](../layer-3/resolution.md) and cannot change.

### Not eligible (403)

`{"error": "You cannot vote on this"}`. The user has no [Eligibility](../layer-3/eligibility.md) under the policy of the submission.

### Invalid vote (400)

`{"error": "X-QAP-Vote must be a whole number"}`. See [Vote value](../layer-4/vote-value.md).

### Withhold not permitted (400)

`{"error": "withhold is not permitted for this poll"}`. See [Withhold rule](../layer-4/withhold-rule.md).

### Not the author (403)

`{"error": "Only the author can delete this submission"}`. See [Authorship](../layer-4/authorship.md).

### Missing comment (406)

`{"error": "QAP comment header required", "hint": "Specify X-QAP-Comment"}`. A [Comment](../layer-2/comment.md) without the `X-QAP-Comment` header.

### Invalid comment (400)

`{"error": "X-QAP-Comment must be percent-encoded"}`. The `X-QAP-Comment` header of a [Comment](../layer-2/comment.md) is not valid percent-encoding.

### Comment too long (400)

`{"error": "X-QAP-Comment must be at most 500 characters"}`. The decoded text of a [Comment](../layer-2/comment.md) has more than 500 characters. Characters are counted, not bytes.

### Invalid id (400)

`{"error": "the id must be a UUID"}`. The route parameter `id` is not a UUID; see [Is UUID](../layer-4/is-uuid.md).

### Cannot be carried out (422)

`{"error": "the request cannot be carried out", "problems": {"<field>": ["<message>"]}}`. The body did not pass [Decode payload](../layer-4/decode-payload.md), or the Check of the [Endpoint](../host/endpoint.md) found [Problems](../layer-4/problems.md), with a plain request or a [Submit](../layer-2/submit.md). After approval, [Carry out](../layer-3/carry-out.md) stores the same body as the result instead.

---

**Referenced in:**

- [A2 Authorship](../layer-4/authorship.md)
- [A5 Intent](../layer-4/intent.md)
- [A6 Permission](../layer-4/permission.md)
- [A7 Problems](../layer-4/problems.md)
- [A15 Vote value](../layer-4/vote-value.md)
- [A16 Withhold rule](../layer-4/withhold-rule.md)
- [A19 Decode payload](../layer-4/decode-payload.md)
- [A20 Is UUID](../layer-4/is-uuid.md)
- [B1 Carry out](../layer-3/carry-out.md)
- [B2 Cleanup](../layer-3/cleanup.md)
- [B4 Eligibility](../layer-3/eligibility.md)
- [B8 Resolution](../layer-3/resolution.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
- [H2 Current user](../host/current-user.md)
- [H3 Endpoint](../host/endpoint.md)
- [I1 Discover](../layer-2/discover.md)
- [I2 List](../layer-2/list.md)
- [I3 Submit](../layer-2/submit.md)
- [I4 Poll](../layer-2/poll.md)
- [I5 Vote](../layer-2/vote.md)
- [I6 Comment](../layer-2/comment.md)
- [I7 Withdraw](../layer-2/withdraw.md)
