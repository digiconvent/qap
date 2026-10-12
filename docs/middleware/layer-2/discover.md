# I1 Discover

The first step before proposing anything. It returns the [Policies](../host/policy-lookup.md) of the [Permission](../layer-4/permission.md) that the [Current user](../host/current-user.md) may use, according to [Eligibility](../layer-3/eligibility.md), each with its settings: how many votes are needed, whether there is a deadline and an [Extension](../concepts/extension.md), and whether withholding is allowed. The answer also names the [Version](../layer-4/version.md) of QAP. The user picks one of these policies for [Submit](submit.md).

---

**Kind:** intent, `X-QAP-Intent: discover`.

**Headers:** no others.

**Response:** `200`, with the [Version](../layer-4/version.md) in the `X-QAP-Version` header. The body lists the policies the current user may use for this endpoint, each with its settings. The example policy needs 2 votes in favour, is rejected by 1 vote against, and runs out after 48 hours, without an extension.

```json
{
  "endpoint": "/api/library/book/123/delete",
  "permission": "library.book.123.delete",
  "policies": [
    {
      "id": "pol_librarians",
      "name": "two_librarians",
      "favour_threshold": 2,
      "against_threshold": 1,
      "quorum": null,
      "deadline_seconds": 172800,
      "extension_seconds": null,
      "withhold_permitted": false
    }
  ]
}
```

**Errors:** [Unauthorized (401)](../concepts/errors.md#unauthorized-401).

---

**Referenced in:**

- [A6 Permission](../layer-4/permission.md)
- [A12 Version](../layer-4/version.md)
- [E1 QAP endpoint](../layer-1/qap-endpoint.md)
- [I3 Submit](submit.md)
- [U2 Propose a change](../layer-0/propose-a-change.md)
