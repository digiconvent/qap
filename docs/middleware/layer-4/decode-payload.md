# A19 Decode payload

Decodes the JSON body of a request into the payload type of its [Endpoint](../host/endpoint.md), so Check and Execute work with a value of the host and not with raw text.

---

**Kind:** atomic part.

**In:** the body. **Out:** the payload, or the problem `{"body": [<the decoding error>]}`.

An empty body is the empty payload, for a request without one such as a `GET` or `DELETE`. A body that does not decode is refused with [Cannot be carried out (422)](../concepts/errors.md#cannot-be-carried-out-422) before Check is called.

---

**Referenced in:**

- [C3 Errors](../concepts/errors.md)
- [C7 Request](../concepts/request.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
- [H3 Endpoint](../host/endpoint.md)
