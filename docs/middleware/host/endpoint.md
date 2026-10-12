# H3 Endpoint

The host application supplies one endpoint for every route that QAP guards: a handler. Execute is required: it carries the [Request](../concepts/request.md) out, with the stored body, the user and, for an approved submission, the submission itself. Check and Read are optional, depending on the handler. Check says, field by field, whether the body can be carried out at all: a body that fails it is refused, because a submission of it would be useless. With Read, QAP reads the entity of the route before and after Execute, so the [Audit summary](../layer-4/audit-summary.md) records which fields changed.

---

**Kind:** supplied by the host.

**Go:**

```go
type Handler struct {
	Execute func(request qap.Request) qap.Result
	Check   func(body string) map[string][]error
	Read    func(id string) any
}
```

Handed to QAP per route during [Setup](setup.md), with `Run`: `guard.RequiresPermission("library.book.{id}.delete").Run(deleteBook(library))`.

**Execute:** called with the [Request](../concepts/request.md): the method, url, route, route parameters, headers and body, the user, and the submission. For a plain request the user is the current user and the submission is `nil`; for an approved submission the user is its [author](../concepts/author.md), and the submission holds its id, resolved permission, policy, votes and comments. **Returns:** a [Result](../concepts/result.md).

**Check**, optional: called with the body. **Returns:** the problems per field, and `nil`, an empty map or only `nil` errors when the body can be carried out. A handler without Check passes every body.

**Read**, optional: called with the `id` of the route, which is always a UUID (see [Is UUID](../layer-4/is-uuid.md)), just before and just after Execute of an approved submission. **Returns:** the entity, or `nil` when there is none. [Changes](../layer-4/changes.md) compares both. A handler without Read records no changes.

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

func renameBook(library *Library) qap.Handler {
	return qap.Handler{
		Check: func(body string) map[string][]error {
			var r renaming
			if json.Unmarshal([]byte(body), &r) != nil || r.Title == "" {
				return map[string][]error{"title": {errors.New("is required")}}
			}
			return nil
		},
		Execute: func(request qap.Request) qap.Result {
			var r renaming
			json.Unmarshal([]byte(request.Body), &r)
			library.Rename(request.Params["id"], r.Title)
			return qap.Result{Status: 204}
		},
		Read: func(id string) any {
			return library.Find(id)
		},
	}
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
