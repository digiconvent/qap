# E1 QAP endpoint

Protects one route of the host application and runs its [Endpoint](../host/endpoint.md). The host picks one of four kinds of protection per route, from open to strict, and ends each with `Run(endpoint)`. **Anonymous** lets everyone through, logged in or not. **Authenticated** lets every logged-in [Current user](../host/current-user.md) through. **Requires permission** only accepts requests that carry an [Intent](../layer-4/intent.md): the endpoint works out the [Permission](../layer-4/permission.md) of the route and hands the request to [Discover](../layer-2/discover.md), [Submit](../layer-2/submit.md), [List](../layer-2/list.md), [Poll](../layer-2/poll.md), [Vote](../layer-2/vote.md), [Comment](../layer-2/comment.md) or [Withdraw](../layer-2/withdraw.md). **Authenticated or requires permission** combines the last two: a request without an intent passes like authenticated, and a request with an intent goes through QAP. A request that passes without QAP is checked and executed by the endpoint right away. The host application creates QAP endpoints during [Setup](../host/setup.md).

---

**Kind:** endpoint: Fiber handler for one route.

| Protection | Go | Without an intent | With an intent |
|---|---|---|---|
| anonymous | `guard.AllowAnonymous().Run(qap.Handle(endpoint))` | Runs the endpoint | Runs the endpoint; the intent is ignored |
| authenticated | `guard.AllowAuthenticatedOrRequiresPermission().Run(qap.Handle(endpoint))` | Runs the endpoint when logged in, otherwise [Unauthorized (401)](../concepts/errors.md#unauthorized-401) | Same as without |
| authenticated or requires permission | `guard.AuthenticatedOrRequiresPermission(permissions...).Run(qap.Handle(endpoint))` | Runs the endpoint when logged in, otherwise [Unauthorized (401)](../concepts/errors.md#unauthorized-401) | Handled by QAP, as below |
| requires permission | `guard.RequiresPermission(permissions...).Run(qap.Handle(endpoint))` | [Missing intent (406)](../concepts/errors.md#missing-intent-406) | Handled by QAP, as below |

The two protections that take permissions need at least one. Without, they panic when the route is set up, not on its first request.

Before anything else, a route parameter `id` must pass [Is UUID](../layer-4/is-uuid.md), whatever the protection; otherwise [Invalid id (400)](../concepts/errors.md#invalid-id-400).

**Runs the endpoint:** the incoming [Request](../concepts/request.md), with the current user, is copied; its body goes through [Decode payload](../layer-4/decode-payload.md) to Check. [Problems](../layer-4/problems.md) found: [Cannot be carried out (422)](../concepts/errors.md#cannot-be-carried-out-422). Otherwise Execute runs, and [Write result](../layer-4/write-result.md) turns its result into the answer.

**Handled by QAP:**

1. [Intent](../layer-4/intent.md): [Unknown intent (400)](../concepts/errors.md#unknown-intent-400) when unknown.
2. [Current user](../host/current-user.md): [Unauthorized (401)](../concepts/errors.md#unauthorized-401) when nobody is logged in.
3. [Permission](../layer-4/permission.md), only for [Discover](../layer-2/discover.md) and [Submit](../layer-2/submit.md): [Misconfigured route (500)](../concepts/errors.md#misconfigured-route-500) when the route misses a placeholder. [List](../layer-2/list.md) goes by the route template, and the other intents work on an existing submission, which already holds its permission and policy, so they skip this step. A route can list several permissions; [Discover](../layer-2/discover.md) then offers the policies of all of them, and names the first permission.
4. The intent handles the rest, with the endpoint of the route at hand for [Submit](../layer-2/submit.md).

**Example:** a member lost access to their phone and asks others to reset their password. The route `/api/iam/user/:id/set-password` uses "authenticated or requires permission" with `iam.user.set-password`. The member can change their own password with a plain request, and the Check of the endpoint makes sure the id in the URL is their own. Others reset it through QAP: one submits, the policy decides who votes, and the endpoint carries out the change once it is approved. QAP never checks ownership; that stays in the endpoint.

---

**Referenced in:**

- [A5 Intent](../layer-4/intent.md)
- [H2 Current user](../host/current-user.md)
- [H7 Setup](../host/setup.md)
