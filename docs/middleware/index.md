# QAP middleware plan

Every name below has its own page and its own code: the letter names the layer, the number the place in the summary of that layer. Each page describes its name in a few sentences, followed by technical details below a horizontal line, and a list of the pages that refer to it. Each heading below links to a summary of that layer. The [flow](../flow.md) shows in which order the pages take part, from the first request to the audit trail. Start at the top and follow the links down.

## [Layer 0: what a person does](layer-0/summary.md)

[U1 Approve or reject](layer-0/approve-or-reject.md) · [U2 Propose a change](layer-0/propose-a-change.md) · [U3 Review pending work](layer-0/review-pending-work.md) · [U4 Take back a proposal](layer-0/take-back-a-proposal.md)

## [Layer 1: the endpoint](layer-1/summary.md)

[E1 QAP endpoint](layer-1/qap-endpoint.md)

## [Layer 2: the intents](layer-2/summary.md)

[I1 Discover](layer-2/discover.md) · [I2 List](layer-2/list.md) · [I3 Submit](layer-2/submit.md) · [I4 Poll](layer-2/poll.md) · [I5 Vote](layer-2/vote.md) · [I6 Comment](layer-2/comment.md) · [I7 Withdraw](layer-2/withdraw.md)

## [Layer 3: built from atomic parts](layer-3/summary.md)

[B1 Carry out](layer-3/carry-out.md) · [B2 Cleanup](layer-3/cleanup.md) · [B3 Deadline timer](layer-3/deadline-timer.md) · [B4 Eligibility](layer-3/eligibility.md) · [B5 Execution](layer-3/execution.md) · [B6 Poll evaluation](layer-3/poll-evaluation.md) · [B7 Rejection](layer-3/rejection.md) · [B8 Resolution](layer-3/resolution.md) · [B9 Store](layer-3/store.md) · [B10 Voter count](layer-3/voter-count.md) · [B11 Withdrawal](layer-3/withdrawal.md)

## [Layer 4: atomic](layer-4/summary.md)

[A1 Audit summary](layer-4/audit-summary.md) · [A2 Authorship](layer-4/authorship.md) · [A3 Deadline](layer-4/deadline.md) · [A4 Expiry](layer-4/expiry.md) · [A5 Intent](layer-4/intent.md) · [A6 Permission](layer-4/permission.md) · [A7 Problems](layer-4/problems.md) · [A8 Request copy](layer-4/request-copy.md) · [A9 Required count](layer-4/required-count.md) · [A10 Submission id](layer-4/submission-id.md) · [A11 Submission view](layer-4/submission-view.md) · [A12 Version](layer-4/version.md) · [A13 Vote count](layer-4/vote-count.md) · [A14 Vote recording](layer-4/vote-recording.md) · [A15 Vote value](layer-4/vote-value.md) · [A16 Withhold rule](layer-4/withhold-rule.md) · [A17 Write result](layer-4/write-result.md) · [A18 Changes](layer-4/changes.md) · [A19 Decode payload](layer-4/decode-payload.md) · [A20 Is UUID](layer-4/is-uuid.md)

## [Concepts](concepts/summary.md)

[C1 Anonymity](concepts/anonymity.md) · [C2 Author](concepts/author.md) · [C3 Errors](concepts/errors.md) · [C4 Extension](concepts/extension.md) · [C5 Policy](concepts/policy.md) · [C6 Quorum](concepts/quorum.md) · [C7 Request](concepts/request.md) · [C8 Result](concepts/result.md) · [C9 Sensitive read](concepts/sensitive-read.md) · [C10 Submission](concepts/submission.md) · [C11 Threshold](concepts/threshold.md)

## [Supplied by the host](host/summary.md)

[H1 Audit trail](host/audit-trail.md) · [H2 Current user](host/current-user.md) · [H3 Endpoint](host/endpoint.md) · [H4 Member count](host/member-count.md) · [H5 Membership](host/membership.md) · [H6 Policy lookup](host/policy-lookup.md) · [H7 Setup](host/setup.md)
