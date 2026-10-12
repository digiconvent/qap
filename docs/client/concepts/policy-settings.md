# cC5 Policy settings

A [Policy reference](policy-reference.md) together with the rules that decide a submission under it. [Discovered](../states/discovered.md) lists them, so a user can see what a proposal will take before making it. [Read policy](../layer-4/read-policy.md) builds them.

---

**Kind:** concept, `PolicySettings`.

| Field | Type | Meaning |
|---|---|---|
| `id`, `name` | `string` | The policy. |
| `favourThreshold` | `number` | Votes in favour that approve. Negative means a percentage of those who may vote: `-50` is half, rounded up. |
| `againstThreshold` | `number \| null` | Votes against that reject; `null` when objections never reject. Negative means a percentage. |
| `quorum` | `number \| null` | How many must vote for or against before anything is decided; `null` for none. |
| `deadlineSeconds` | `number \| null` | Seconds after which a submission without enough votes is rejected; `null` for none. |
| `extensionSeconds` | `number \| null` | Extra seconds after the deadline: when nobody voted against, the submission waits that long, and passes at the end if still nobody did; `null` for none. |
| `withholdPermitted` | `boolean` | Whether abstaining is allowed. |

---

**Referenced in:**

- [cA7 Read policy](../layer-4/read-policy.md)
- [cM3 Submit](../layer-2/submit.md)
- [cS1 Discovered](../states/discovered.md)
