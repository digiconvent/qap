# A7 Problems

Turns what the Check of an [Endpoint](../host/endpoint.md) returns into the problems a client sees: the messages of the errors, per field. A check that returned no errors has no problems, so the request can be carried out.

---

**Kind:** atomic part.

**In:** `map[string][]error`. **Out:** `map[string][]string`, or nothing when there are no errors. `nil` errors and fields without errors are left out.

`{"title": [errors.New("is required"), nil], "year": nil}` gives `{"title": ["is required"]}`. The client receives them in [Cannot be carried out (422)](../concepts/errors.md#cannot-be-carried-out-422).

---

**Referenced in:**

- [B1 Carry out](../layer-3/carry-out.md)
- [C3 Errors](../concepts/errors.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
