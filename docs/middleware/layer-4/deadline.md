# A3 Deadline

A [policy](../concepts/policy.md) can give [submissions](../concepts/submission.md) a deadline, in seconds after they are created, and an [Extension](../concepts/extension.md) after it. This part tells whether the deadline has passed, and whether the extension has, by comparing the creation time plus those seconds with the current time, in Unix seconds. The [Deadline timer](../layer-3/deadline-timer.md) uses it to decide a submission on time.

---

**Kind:** atomic part.

**In:** the creation time of the submission, the deadline and extension of its policy in seconds, and the current time. **Out:** passed or not.

- The deadline passes when `now >= created + deadline`. A policy without a deadline never passes.
- The extension passes when `now >= created + deadline + extension`. It never passes without both a deadline and an extension.

---

**Referenced in:**

- [B3 Deadline timer](../layer-3/deadline-timer.md)
- [B6 Poll evaluation](../layer-3/poll-evaluation.md)
- [C4 Extension](../concepts/extension.md)
- [C5 Policy](../concepts/policy.md)
