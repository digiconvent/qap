# H3 Endpoint

The host application supplies one endpoint for every route that QAP guards. It is generic over its payload: the JSON body of the request, decoded into a type of the host. Check only sees that payload and says, field by field, whether it can be carried out at all: a payload that fails it is refused, because a submission of it would be useless. Execute carries the [Request](../concepts/request.md) out, with the payload, the stored request, the user and, for an approved submission, the submission itself: it is the handler of the route. An endpoint can also be a reader: then QAP reads the entity of the route before and after Execute, so the [Audit summary](../layer-4/audit-summary.md) records which fields changed.

---

**Kind:** supplied by the host.

**Go:**

```go
type Endpoint[Payload any] interface {
	Check(payload Payload) map[string][]error
	Execute(request qap.Request[Payload]) qap.Result
}

type Reader interface {
	Read(id string) any
}
```

Handed to QAP per route during [Setup](setup.md), with `Run` and `qap.Handle`: `guard.RequiresPermission("library.book.{id}.delete").Run(qap.Handle(deleteBook{}))`. Go cannot put a type parameter on a method, so `qap.Handle` turns the generic endpoint into a handler; Go infers the payload type from the methods.

**Check:** called with the payload, decoded by [Decode payload](../layer-4/decode-payload.md); a body that does not decode never reaches it. **Returns:** the problems per field, and `nil`, an empty map or only `nil` errors when the payload can be carried out.

**Execute:** called with the [Request](../concepts/request.md): the payload, method, url, route, route parameters and headers, the user, and the submission. For a plain request the user is the current user and the submission is `nil`; for an approved submission the user is its [author](../concepts/author.md), and the submission holds its id, resolved permission, policy, votes and comments. **Returns:** a [Result](../concepts/result.md).

**Read**, optional: called with the `id` of the route, which is always a UUID (see [Is UUID](../layer-4/is-uuid.md)), just before and just after Execute of an approved submission. **Returns:** the entity, or `nil` when there is none. [Changes](../layer-4/changes.md) compares both.

| When | Check fails | Check passes |
|---|---|---|
| a plain request without an intent, on a route that allows one | [Cannot be carried out (422)](../concepts/errors.md#cannot-be-carried-out-422) | Execute runs, and its result is the answer |
| [Submit](../layer-2/submit.md) | [Cannot be carried out (422)](../concepts/errors.md#cannot-be-carried-out-422), and nothing is stored | the submission is created |
| [Carry out](../layer-3/carry-out.md), after approval | the stored result has status `422` and the problems; Execute does not run | Read, Execute, Read; the result is stored, the changes go to the audit trail |

An endpoint that panics while an approved submission is carried out gives a result with status `500`, and nothing else is affected.

```go
type renaming struct {
	Title string `json:"title"`
}

type renameBook struct{ library *Library }

func (r renameBook) Check(payload renaming) map[string][]error {
	if payload.Title == "" {
		return map[string][]error{"title": {errors.New("is required")}}
	}
	return nil
}

func (r renameBook) Execute(request qap.Request[renaming]) qap.Result {
	r.library.Rename(request.Params["id"], request.Payload.Title)
	return qap.Result{Status: 204}
}

func (r renameBook) Read(id string) any {
	return r.library.Find(id)
}
```

The endpoint calls the logic of the host directly. On a route with "authenticated or requires permission", Execute can let users change only their own account in a plain request, where `request.Submission` is `nil`, while an approved submission changes the account of someone else: there the approval is what allows it.

---

**Referenced in:**

- [A1 Audit summary](../layer-4/audit-summary.md)
- [A7 Problems](../layer-4/problems.md)
- [A8 Request copy](../layer-4/request-copy.md)
- [A17 Write result](../layer-4/write-result.md)
- [A18 Changes](../layer-4/changes.md)
- [A19 Decode payload](../layer-4/decode-payload.md)
- [B1 Carry out](../layer-3/carry-out.md)
- [B5 Execution](../layer-3/execution.md)
- [C3 Errors](../concepts/errors.md)
- [C7 Request](../concepts/request.md)
- [C8 Result](../concepts/result.md)
- [C9 Sensitive read](../concepts/sensitive-read.md)
- [C10 Submission](../concepts/submission.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
- [H7 Setup](setup.md)
- [I3 Submit](../layer-2/submit.md)
- [I5 Vote](../layer-2/vote.md)
