# H2 Current user

The host application tells who sends a request, for example from a session or a token. QAP never logs anyone in itself. An empty answer means nobody is logged in, and the [QAP endpoint](../layer-1/qap-endpoint.md) refuses the request.

---

**Kind:** supplied by the host.

**Go:** `func(ctx fiber.Ctx) string`. Handed to QAP during [Setup](setup.md).

**Called with:** the incoming request. **Returns:** the user id, or `""` when nobody is logged in.

`""` on a [QAP endpoint](../layer-1/qap-endpoint.md): [Unauthorized (401)](../concepts/errors.md#unauthorized-401).

---

**Referenced in:**

- [C2 Author](../concepts/author.md)
- [C3 Errors](../concepts/errors.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
- [H7 Setup](setup.md)
- [I1 Discover](../layer-2/discover.md)
- [I2 List](../layer-2/list.md)
- [I3 Submit](../layer-2/submit.md)
- [I7 Withdraw](../layer-2/withdraw.md)
