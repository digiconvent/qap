# A5 Intent

Every request to a [QAP endpoint](../layer-1/qap-endpoint.md) says what it wants in the `X-QAP-Intent` header. This part reads that header and names the intent: discover, submit, vote, comment, list, poll or delete-[submission](../concepts/submission.md). A request without the header, or with a value that is not one of these, is reported as missing or unknown, and the endpoint decides what that means for its protection.

---

**Kind:** atomic part.

**In:** the value of the `X-QAP-Intent` header. **Out:** one of `discover`, `submit`, `vote`, `comment`, `list`, `poll`, `delete-submission`, or missing, or unknown.

- Header missing: [Missing intent (406)](../concepts/errors.md#missing-intent-406) under "requires permission"; under "authenticated or requires permission" the request passes like a plain request.
- Unknown value: [Unknown intent (400)](../concepts/errors.md#unknown-intent-400).

---

**Referenced in:**

- [C3 Errors](../concepts/errors.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
