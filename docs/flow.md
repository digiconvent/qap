# Flow

This is every flow of layer 0, from the person to the host and back, in chronological order. Each section is one flow; each row is one part of the system that is addressed, in the lane of the actor it belongs to, with what it does at that moment. Parts that recur in every flow, such as a request on its way to the server, have their own section below the flows, and a flow names them in one row.

---

## Setup

Before any person acts, the host sets QAP up on the server and the app opens the route in the browser.

| # | Person | App | Client | Endpoint | Intent | Built from atomic parts | Atomic | Host |
|---|---|---|---|---|---|---|---|---|
| 1 |  |  |  |  |  |  |  | [H7 Setup](middleware/host/setup.md) assembles QAP and guards every route |
| 2 |  |  |  |  |  |  |  | [H3 Endpoint](middleware/host/endpoint.md) is given to the route, with Execute and optionally Check and Read |
| 3 |  | [cH2 Setup](client/app/setup.md) installs the client |  |  |  |  |  |  |
| 4 |  | [cH1 Send](client/app/send.md) is the function that sends every request |  |  |  |  |  |  |
| 5 |  |  | [cE1 Endpoint](client/layer-1/endpoint.md) is opened with the path and method of the route |  |  |  |  |  |

## Propose a change

The person proposes a change and waits for the decision: [U2 Propose a change](middleware/layer-0/propose-a-change.md) and [cU2 Propose a change](client/layer-0/propose-a-change.md).

| # | Person | App | Client | Endpoint | Intent | Built from atomic parts | Atomic | Host |
|---|---|---|---|---|---|---|---|---|
| 1 | [cU2 Propose a change](client/layer-0/propose-a-change.md): the person wants something that needs approval |  |  |  |  |  |  |  |
| 2 |  |  | [cM1 Discover](client/layer-2/discover.md) asks which policies the person may use |  |  |  |  |  |
| 3 |  |  | [cB1 Ask](client/layer-3/ask.md) sends the route with `X-QAP-Intent: discover`, see [A request](#a-request) |  |  |  |  |  |
| 4 |  |  |  | [E1 QAP endpoint](middleware/layer-1/qap-endpoint.md) checks the id, the intent and the user, and hands the request to the intent |  |  |  |  |
| 5 |  |  |  |  | [I1 Discover](middleware/layer-2/discover.md) collects the policies the person may use |  |  |  |
| 6 |  |  |  |  |  |  |  | [H6 Policy lookup](middleware/host/policy-lookup.md) gives the policies of each permission of the route |
| 7 |  |  |  |  |  | [B4 Eligibility](middleware/layer-3/eligibility.md) keeps the policies that include the person |  |  |
| 8 |  |  |  |  |  |  |  | [H5 Membership](middleware/host/membership.md) says whether the person is in one of their groups |
| 9 |  |  |  |  |  |  | [A12 Version](middleware/layer-4/version.md) goes along in the answer as `X-QAP-Version` |  |
| 10 |  |  | [cA5 Read body](client/layer-4/read-body.md) reads the answer |  |  |  |  |  |
| 11 |  |  | [cA2 Is success](client/layer-4/is-success.md) checks the status; a refusal goes to [cA8 Read refusal](client/layer-4/read-refusal.md) and ends in [cS7 Refused](client/states/refused.md) |  |  |  |  |  |
| 12 |  |  | [cA1 Is discovery](client/layer-4/is-discovery.md) checks it is a discovery; otherwise [cA4 Not QAP](client/layer-4/not-qap.md) |  |  |  |  |  |
| 13 |  |  | [cA9 Version mismatch](client/layer-4/version-mismatch.md) compares the versions of client and server |  |  |  |  |  |
| 14 |  |  | [cA7 Read policy](client/layer-4/read-policy.md) reads each policy |  |  |  |  |  |
| 15 |  |  | [cS1 Discovered](client/states/discovered.md) shows the policies |  |  |  |  |  |
| 16 | the person picks a policy and gives the body |  |  |  |  |  |  |  |
| 17 |  |  | [cM3 Submit](client/layer-2/submit.md) sends the request under that policy |  |  |  |  |  |
| 18 |  |  | [cB1 Ask](client/layer-3/ask.md) sends the route with `X-QAP-Intent: submit`, `X-QAP-Policy-ID` and the body, see [A request](#a-request) |  |  |  |  |  |
| 19 |  |  |  | [E1 QAP endpoint](middleware/layer-1/qap-endpoint.md) checks the id, the intent and the user, and hands the request to the intent |  |  |  |  |
| 20 |  |  |  |  | [I3 Submit](middleware/layer-2/submit.md) finds the chosen policy among the eligible ones; without one: 406, or 403 |  |  |  |
| 21 |  |  |  |  |  |  |  | [H6 Policy lookup](middleware/host/policy-lookup.md) gives the policies of each permission of the route |
| 22 |  |  |  |  |  | [B4 Eligibility](middleware/layer-3/eligibility.md) keeps the policies that include the person |  |  |
| 23 |  |  |  |  |  |  |  | [H5 Membership](middleware/host/membership.md) says whether the person is in one of their groups |
| 24 |  |  |  |  |  |  | [A8 Request copy](middleware/layer-4/request-copy.md) copies the request, so it can be carried out later |  |
| 25 |  |  |  |  |  |  |  | [H3 Endpoint](middleware/host/endpoint.md) checks the body, when it has Check |
| 26 |  |  |  |  |  |  | [A7 Problems](middleware/layer-4/problems.md) turns the problems into the answer: 422, and nothing is stored |  |
| 27 |  |  |  |  |  |  | [A10 Submission id](middleware/layer-4/submission-id.md) gives the submission its id |  |
| 28 |  |  |  |  |  | [B10 Voter count](middleware/layer-3/voter-count.md) counts who may vote |  |  |
| 29 |  |  |  |  |  |  |  | [H4 Member count](middleware/host/member-count.md) counts the members of the groups of the policy |
| 30 |  |  |  |  |  | [B9 Store](middleware/layer-3/store.md) keeps the submission, with the vote in favour of its author |  |  |
| 31 |  |  |  |  |  | [B6 Poll evaluation](middleware/layer-3/poll-evaluation.md) decides whether the votes settle it |  |  |
| 32 |  |  |  |  |  |  | [A13 Vote count](middleware/layer-4/vote-count.md) counts the votes in favour, against and withheld |  |
| 33 |  |  |  |  |  |  | [A9 Required count](middleware/layer-4/required-count.md) gives how many votes each threshold needs |  |
| 34 |  |  |  |  |  |  | [A3 Deadline](middleware/layer-4/deadline.md) says whether the deadline has passed |  |
| 35 |  |  |  |  |  | executed: [B5 Execution](middleware/layer-3/execution.md) carries it out, see [Carrying out](#carrying-out) |  |  |
| 36 |  |  |  |  |  | rejected: [B7 Rejection](middleware/layer-3/rejection.md) rejects it, see [Rejecting](#rejecting) |  |  |
| 37 |  |  |  |  |  | still pending with a deadline: [B3 Deadline timer](middleware/layer-3/deadline-timer.md) starts |  |  |
| 38 |  |  |  |  |  |  | [A11 Submission view](middleware/layer-4/submission-view.md) is the answer, as the person may see it |  |
| 39 |  |  | [cB4 Read state](client/layer-3/read-state.md) turns the answer into [cS3 Pending](client/states/pending.md), [cS4 Executed](client/states/executed.md) or [cS7 Refused](client/states/refused.md), see [An answer](#an-answer) |  |  |  |  |  |
| 40 | the person shares the id with colleagues |  |  |  |  |  |  |  |
| 41 |  |  | [cM6 Listen](client/layer-2/listen.md) polls in the background |  |  |  |  |  |
| 42 |  |  | [cM5 Poll](client/layer-2/poll.md) asks for the submission |  |  |  |  |  |
| 43 |  |  | [cB1 Ask](client/layer-3/ask.md) sends the route with `X-QAP-Intent: poll` and `X-QAP-ID`, see [A request](#a-request) |  |  |  |  |  |
| 44 |  |  |  | [E1 QAP endpoint](middleware/layer-1/qap-endpoint.md) checks the id, the intent and the user, and hands the request to the intent |  |  |  |  |
| 45 |  |  |  |  | [I4 Poll](middleware/layer-2/poll.md) answers with the submission |  |  |  |
| 46 |  |  |  |  |  | [B9 Store](middleware/layer-3/store.md) finds the submission by `X-QAP-ID` |  |  |
| 47 |  |  |  |  |  | [B4 Eligibility](middleware/layer-3/eligibility.md) checks that its policy includes the person |  |  |
| 48 |  |  |  |  |  |  |  | [H5 Membership](middleware/host/membership.md) says whether the person is in one of its groups |
| 49 |  |  |  |  |  | [B9 Store](middleware/layer-3/store.md) notes the time of this poll, which keeps it from [cleanup](#without-a-person) |  |  |
| 50 |  |  |  |  |  |  | [A11 Submission view](middleware/layer-4/submission-view.md) is the answer, as the person may see it |  |
| 51 |  |  | [cB4 Read state](client/layer-3/read-state.md) turns the answer into a state, see [An answer](#an-answer) |  |  |  |  |  |
| 52 |  |  | [cM6 Listen](client/layer-2/listen.md) tells the person when it changed, and stops at [cS4 Executed](client/states/executed.md) or [cS5 Rejected](client/states/rejected.md) |  |  |  |  |  |

## Review pending work

A colleague looks at what waits for their decision: [U3 Review pending work](middleware/layer-0/review-pending-work.md) and [cU3 Review pending work](client/layer-0/review-pending-work.md).

| # | Person | App | Client | Endpoint | Intent | Built from atomic parts | Atomic | Host |
|---|---|---|---|---|---|---|---|---|
| 1 | [cU3 Review pending work](client/layer-0/review-pending-work.md): the person checks what waits for them |  |  |  |  |  |  |  |
| 2 |  |  | [cM2 List](client/layer-2/list.md) asks for the pending submissions of the route |  |  |  |  |  |
| 3 |  |  | [cB1 Ask](client/layer-3/ask.md) sends the route with `X-QAP-Intent: list`, see [A request](#a-request) |  |  |  |  |  |
| 4 |  |  |  | [E1 QAP endpoint](middleware/layer-1/qap-endpoint.md) checks the id, the intent and the user, and hands the request to the intent |  |  |  |  |
| 5 |  |  |  |  | [I2 List](middleware/layer-2/list.md) lists the pending submissions |  |  |  |
| 6 |  |  |  |  |  | [B9 Store](middleware/layer-3/store.md) gives the pending submissions of the route template and method |  |  |
| 7 |  |  |  |  |  | [B4 Eligibility](middleware/layer-3/eligibility.md) keeps the ones whose policy includes the person |  |  |
| 8 |  |  |  |  |  |  |  | [H5 Membership](middleware/host/membership.md) says whether the person is in one of its groups |
| 9 |  |  |  |  |  |  | [A11 Submission view](middleware/layer-4/submission-view.md) is the answer for each, as the person may see it |  |
| 10 |  |  | [cA5 Read body](client/layer-4/read-body.md) reads the answer |  |  |  |  |  |
| 11 |  |  | [cA2 Is success](client/layer-4/is-success.md) checks the status; a refusal goes to [cA8 Read refusal](client/layer-4/read-refusal.md) and ends in [cS7 Refused](client/states/refused.md) |  |  |  |  |  |
| 12 |  |  | [cA3 Is view](client/layer-4/is-view.md) checks each one is a view; otherwise [cA4 Not QAP](client/layer-4/not-qap.md) |  |  |  |  |  |
| 13 |  |  | [cB3 Pending state](client/layer-3/pending-state.md) makes each one a pending state |  |  |  |  |  |
| 14 |  |  | [cA6 Read facts](client/layer-4/read-facts.md) reads the facts of each |  |  |  |  |  |
| 15 |  |  | [cS2 Listed](client/states/listed.md) shows them; the person picks one, and goes on with [Approve or reject](#approve-or-reject) |  |  |  |  |  |

## Approve or reject

A colleague comments on the submission of someone else and votes: [U1 Approve or reject](middleware/layer-0/approve-or-reject.md) and [cU1 Approve or reject](client/layer-0/approve-or-reject.md).

| # | Person | App | Client | Endpoint | Intent | Built from atomic parts | Atomic | Host |
|---|---|---|---|---|---|---|---|---|
| 1 | [cU1 Approve or reject](client/layer-0/approve-or-reject.md): the person opens a submission, from the list or from a shared id |  |  |  |  |  |  |  |
| 2 |  |  | [cM4 Follow](client/layer-2/follow.md) opens it by its id |  |  |  |  |  |
| 3 |  |  | [cM5 Poll](client/layer-2/poll.md) asks for the submission |  |  |  |  |  |
| 4 |  |  | [cB1 Ask](client/layer-3/ask.md) sends the route with `X-QAP-Intent: poll` and `X-QAP-ID`, see [A request](#a-request) |  |  |  |  |  |
| 5 |  |  |  | [E1 QAP endpoint](middleware/layer-1/qap-endpoint.md) checks the id, the intent and the user, and hands the request to the intent |  |  |  |  |
| 6 |  |  |  |  | [I4 Poll](middleware/layer-2/poll.md) answers with the submission |  |  |  |
| 7 |  |  |  |  |  | [B9 Store](middleware/layer-3/store.md) finds the submission by `X-QAP-ID` |  |  |
| 8 |  |  |  |  |  | [B4 Eligibility](middleware/layer-3/eligibility.md) checks that its policy includes the person |  |  |
| 9 |  |  |  |  |  |  |  | [H5 Membership](middleware/host/membership.md) says whether the person is in one of its groups |
| 10 |  |  |  |  |  | [B9 Store](middleware/layer-3/store.md) notes the time of this poll, which keeps it from [cleanup](#without-a-person) |  |  |
| 11 |  |  |  |  |  |  | [A11 Submission view](middleware/layer-4/submission-view.md) is the answer, as the person may see it |  |
| 12 |  |  | [cB4 Read state](client/layer-3/read-state.md) turns the answer into [cS3 Pending](client/states/pending.md), see [An answer](#an-answer) |  |  |  |  |  |
| 13 | the person explains their choice |  |  |  |  |  |  |  |
| 14 |  |  | [cM8 Comment](client/layer-2/comment.md) sends the text |  |  |  |  |  |
| 15 |  |  | [cB1 Ask](client/layer-3/ask.md) sends the route with `X-QAP-Intent: comment`, `X-QAP-ID` and `X-QAP-Comment`, see [A request](#a-request) |  |  |  |  |  |
| 16 |  |  |  | [E1 QAP endpoint](middleware/layer-1/qap-endpoint.md) checks the id, the intent and the user, and hands the request to the intent |  |  |  |  |
| 17 |  |  |  |  | [I6 Comment](middleware/layer-2/comment.md) checks the text: present, decodable and at most 500 characters |  |  |  |
| 18 |  |  |  |  |  | [B9 Store](middleware/layer-3/store.md) finds the submission by `X-QAP-ID` |  |  |
| 19 |  |  |  |  |  | [B4 Eligibility](middleware/layer-3/eligibility.md) checks that its policy includes the person |  |  |
| 20 |  |  |  |  |  |  |  | [H5 Membership](middleware/host/membership.md) says whether the person is in one of its groups |
| 21 |  |  |  |  |  | [B9 Store](middleware/layer-3/store.md) adds the comment |  |  |
| 22 |  |  |  |  |  |  | [A11 Submission view](middleware/layer-4/submission-view.md) is the answer, as the person may see it |  |
| 23 |  |  | [cB4 Read state](client/layer-3/read-state.md) turns the answer into [cS3 Pending](client/states/pending.md), see [An answer](#an-answer) |  |  |  |  |  |
| 24 | the person chooses favour, against or withhold |  |  |  |  |  |  |  |
| 25 |  |  | [cM7 Vote](client/layer-2/vote.md) sends the choice |  |  |  |  |  |
| 26 |  |  | [cA10 Vote value](client/layer-4/vote-value.md) turns the choice into `X-QAP-Vote` |  |  |  |  |  |
| 27 |  |  | [cB1 Ask](client/layer-3/ask.md) sends the route with `X-QAP-Intent: vote`, `X-QAP-ID` and `X-QAP-Vote`, see [A request](#a-request) |  |  |  |  |  |
| 28 |  |  |  | [E1 QAP endpoint](middleware/layer-1/qap-endpoint.md) checks the id, the intent and the user, and hands the request to the intent |  |  |  |  |
| 29 |  |  |  |  | [I5 Vote](middleware/layer-2/vote.md) refuses a submission that is decided already: 409 |  |  |  |
| 30 |  |  |  |  |  | [B9 Store](middleware/layer-3/store.md) finds the submission by `X-QAP-ID` |  |  |
| 31 |  |  |  |  |  | [B4 Eligibility](middleware/layer-3/eligibility.md) checks that its policy includes the person |  |  |
| 32 |  |  |  |  |  |  |  | [H5 Membership](middleware/host/membership.md) says whether the person is in one of its groups |
| 33 |  |  |  |  |  |  | [A15 Vote value](middleware/layer-4/vote-value.md) reads `X-QAP-Vote` |  |
| 34 |  |  |  |  |  |  | [A16 Withhold rule](middleware/layer-4/withhold-rule.md) refuses withhold when the policy does not permit it |  |
| 35 |  |  |  |  |  | [B9 Store](middleware/layer-3/store.md) records the vote |  |  |
| 36 |  |  |  |  |  |  | [A14 Vote recording](middleware/layer-4/vote-recording.md) replaces an earlier vote of the person |  |
| 37 |  |  |  |  |  | [B6 Poll evaluation](middleware/layer-3/poll-evaluation.md) decides whether the votes settle it |  |  |
| 38 |  |  |  |  |  |  | [A13 Vote count](middleware/layer-4/vote-count.md) counts the votes in favour, against and withheld |  |
| 39 |  |  |  |  |  |  | [A9 Required count](middleware/layer-4/required-count.md) gives how many votes each threshold needs |  |
| 40 |  |  |  |  |  |  | [A3 Deadline](middleware/layer-4/deadline.md) says whether the deadline has passed |  |
| 41 |  |  |  |  |  | executed: [B5 Execution](middleware/layer-3/execution.md) carries it out, see [Carrying out](#carrying-out) |  |  |
| 42 |  |  |  |  |  | rejected: [B7 Rejection](middleware/layer-3/rejection.md) rejects it, see [Rejecting](#rejecting) |  |  |
| 43 |  |  |  |  |  | decided: [B3 Deadline timer](middleware/layer-3/deadline-timer.md) is cancelled |  |  |
| 44 |  |  |  |  |  |  | [A11 Submission view](middleware/layer-4/submission-view.md) is the answer, as the person may see it |  |
| 45 |  |  | [cB4 Read state](client/layer-3/read-state.md) turns the answer into [cS3 Pending](client/states/pending.md), [cS4 Executed](client/states/executed.md) or [cS5 Rejected](client/states/rejected.md), see [An answer](#an-answer) |  |  |  |  |  |

## Take back a proposal

The author withdraws their submission before it is decided: [U4 Take back a proposal](middleware/layer-0/take-back-a-proposal.md) and [cU4 Take back a proposal](client/layer-0/take-back-a-proposal.md).

| # | Person | App | Client | Endpoint | Intent | Built from atomic parts | Atomic | Host |
|---|---|---|---|---|---|---|---|---|
| 1 | [cU4 Take back a proposal](client/layer-0/take-back-a-proposal.md): the author changes their mind |  |  |  |  |  |  |  |
| 2 |  |  | [cM9 Withdraw](client/layer-2/withdraw.md) withdraws the submission |  |  |  |  |  |
| 3 |  |  | [cB1 Ask](client/layer-3/ask.md) sends the route with `X-QAP-Intent: delete-submission` and `X-QAP-ID`, see [A request](#a-request) |  |  |  |  |  |
| 4 |  |  |  | [E1 QAP endpoint](middleware/layer-1/qap-endpoint.md) checks the id, the intent and the user, and hands the request to the intent |  |  |  |  |
| 5 |  |  |  |  | [I7 Withdraw](middleware/layer-2/withdraw.md) withdraws the submission |  |  |  |
| 6 |  |  |  |  |  | [B9 Store](middleware/layer-3/store.md) finds the submission by `X-QAP-ID` |  |  |
| 7 |  |  |  |  |  |  | [A2 Authorship](middleware/layer-4/authorship.md) refuses anyone but the author: 403 |  |
| 8 |  |  |  |  |  | [B9 Store](middleware/layer-3/store.md) deletes it, unless it was decided first: 409 |  |  |
| 9 |  |  |  |  |  | [B3 Deadline timer](middleware/layer-3/deadline-timer.md) is cancelled |  |  |
| 10 |  |  |  |  |  | [B11 Withdrawal](middleware/layer-3/withdrawal.md) records the withdrawal |  |  |
| 11 |  |  |  |  |  |  | [A1 Audit summary](middleware/layer-4/audit-summary.md) summarizes it |  |
| 12 |  |  |  |  |  |  |  | [H1 Audit trail](middleware/host/audit-trail.md) records the summary |
| 13 |  |  | [cA5 Read body](client/layer-4/read-body.md) reads the answer |  |  |  |  |  |
| 14 |  |  | [cA2 Is success](client/layer-4/is-success.md) checks the status; a refusal goes to [cA8 Read refusal](client/layer-4/read-refusal.md) and ends in [cS7 Refused](client/states/refused.md) |  |  |  |  |  |
| 15 |  |  | [cS6 Withdrawn](client/states/withdrawn.md) shows it is gone |  |  |  |  |  |

## Without a person

Two things happen on the server by themselves: a deadline passes, and decided submissions are cleaned up.

| # | Person | App | Client | Endpoint | Intent | Built from atomic parts | Atomic | Host |
|---|---|---|---|---|---|---|---|---|
| 1 |  |  |  |  |  | [B3 Deadline timer](middleware/layer-3/deadline-timer.md) fires when the deadline, or its extension, passes |  |  |
| 2 |  |  |  |  |  | [B9 Store](middleware/layer-3/store.md) gives the submission, if it is still pending |  |  |
| 3 |  |  |  |  |  | [B6 Poll evaluation](middleware/layer-3/poll-evaluation.md) decides with the deadline passed |  |  |
| 4 |  |  |  |  |  |  | [A3 Deadline](middleware/layer-4/deadline.md) says whether the deadline and the extension have passed |  |
| 5 |  |  |  |  |  | executed: [B5 Execution](middleware/layer-3/execution.md) carries it out, see [Carrying out](#carrying-out) |  |  |
| 6 |  |  |  |  |  | pending, with time left in the extension: [B3 Deadline timer](middleware/layer-3/deadline-timer.md) starts again |  |  |
| 7 |  |  |  |  |  | otherwise: [B7 Rejection](middleware/layer-3/rejection.md) rejects it, see [Rejecting](#rejecting) |  |  |
| 8 |  |  |  |  |  | [B2 Cleanup](middleware/layer-3/cleanup.md) looks at the store every tick |  |  |
| 9 |  |  |  |  |  |  | [A4 Expiry](middleware/layer-4/expiry.md) says whether a decided submission was last polled more than a minute ago |  |
| 10 |  |  |  |  |  | [B9 Store](middleware/layer-3/store.md) removes the expired submissions |  |  |

## A request

Every move of the client reaches the server this way, and comes back to the client as an answer.

| # | Person | App | Client | Endpoint | Intent | Built from atomic parts | Atomic | Host |
|---|---|---|---|---|---|---|---|---|
| 1 |  |  | [cB1 Ask](client/layer-3/ask.md) adds the headers of the move |  |  |  |  |  |
| 2 |  | [cH1 Send](client/app/send.md) sends the request |  |  |  |  |  |  |
| 3 |  |  |  | [E1 QAP endpoint](middleware/layer-1/qap-endpoint.md) receives it on its route |  |  |  |  |
| 4 |  |  |  |  |  |  | [A19 Is UUID](middleware/layer-4/is-uuid.md) refuses an `id` in the route that is not a UUID: 400 |  |
| 5 |  |  |  |  |  |  | [A5 Intent](middleware/layer-4/intent.md) reads `X-QAP-Intent`; without one, see [A plain request](#a-plain-request) |  |
| 6 |  |  |  |  |  |  |  | [H2 Current user](middleware/host/current-user.md) says who asks; nobody: 401 |
| 7 |  |  |  |  |  |  | [A6 Permission](middleware/layer-4/permission.md) resolves the permissions of the route, only for discover and submit |  |
| 8 |  |  |  |  | the intent handles it, as in the flow |  |  |  |
| 9 |  |  | [cA5 Read body](client/layer-4/read-body.md) reads the answer |  |  |  |  |  |

## An answer

Every answer that holds a submission becomes a state of the client this way.

| # | Person | App | Client | Endpoint | Intent | Built from atomic parts | Atomic | Host |
|---|---|---|---|---|---|---|---|---|
| 1 |  |  | [cB4 Read state](client/layer-3/read-state.md) reads the answer |  |  |  |  |  |
| 2 |  |  | [cA2 Is success](client/layer-4/is-success.md) checks the status; a refusal goes to [cA8 Read refusal](client/layer-4/read-refusal.md) and ends in [cS7 Refused](client/states/refused.md) |  |  |  |  |  |
| 3 |  |  | [cA3 Is view](client/layer-4/is-view.md) checks it is a view; otherwise [cA4 Not QAP](client/layer-4/not-qap.md) |  |  |  |  |  |
| 4 |  |  | [cB5 Read view](client/layer-3/read-view.md) looks at the resolution |  |  |  |  |  |
| 5 |  |  | pending: [cB3 Pending state](client/layer-3/pending-state.md), with the moves of a pending submission |  |  |  |  |  |
| 6 |  |  | decided: [cB2 Decided state](client/layer-3/decided-state.md), with the result |  |  |  |  |  |
| 7 |  |  | [cA6 Read facts](client/layer-4/read-facts.md) reads the facts |  |  |  |  |  |
| 8 |  |  | ends in [cS3 Pending](client/states/pending.md), [cS4 Executed](client/states/executed.md) or [cS5 Rejected](client/states/rejected.md) |  |  |  |  |  |

## Carrying out

An approved submission is carried out once, as its author, and recorded in the audit trail.

| # | Person | App | Client | Endpoint | Intent | Built from atomic parts | Atomic | Host |
|---|---|---|---|---|---|---|---|---|
| 1 |  |  |  |  |  | [B5 Execution](middleware/layer-3/execution.md) carries out the submission |  |  |
| 2 |  |  |  |  |  | [B8 Resolution](middleware/layer-3/resolution.md) marks it executed, unless it was decided or voted on in the meantime |  |  |
| 3 |  |  |  |  |  | [B1 Carry out](middleware/layer-3/carry-out.md) runs the stored request |  |  |
| 4 |  |  |  |  |  |  |  | [H3 Endpoint](middleware/host/endpoint.md) checks the body again, when it has Check |
| 5 |  |  |  |  |  |  | [A7 Problems](middleware/layer-4/problems.md) turns the problems into a result with status 422, and nothing is executed |  |
| 6 |  |  |  |  |  |  |  | [H3 Endpoint](middleware/host/endpoint.md) reads the entity before, when it has Read and the route an id |
| 7 |  |  |  |  |  |  |  | [H3 Endpoint](middleware/host/endpoint.md) executes the request |
| 8 |  |  |  |  |  |  |  | [H3 Endpoint](middleware/host/endpoint.md) reads the entity after |
| 9 |  |  |  |  |  |  | [A18 Changes](middleware/layer-4/changes.md) compares both |  |
| 10 |  |  |  |  |  | [B9 Store](middleware/layer-3/store.md) keeps the result |  |  |
| 11 |  |  |  |  |  |  | [A1 Audit summary](middleware/layer-4/audit-summary.md) summarizes it, with the changes unless the policy is anonymous |  |
| 12 |  |  |  |  |  |  |  | [H1 Audit trail](middleware/host/audit-trail.md) records the summary |

## Rejecting

A rejected submission is never carried out; only the decision is recorded.

| # | Person | App | Client | Endpoint | Intent | Built from atomic parts | Atomic | Host |
|---|---|---|---|---|---|---|---|---|
| 1 |  |  |  |  |  | [B7 Rejection](middleware/layer-3/rejection.md) rejects the submission |  |  |
| 2 |  |  |  |  |  | [B8 Resolution](middleware/layer-3/resolution.md) marks it rejected, unless it was decided or voted on in the meantime |  |  |
| 3 |  |  |  |  |  |  | [A1 Audit summary](middleware/layer-4/audit-summary.md) summarizes it |  |
| 4 |  |  |  |  |  |  |  | [H1 Audit trail](middleware/host/audit-trail.md) records the summary |

## A plain request

A request without `X-QAP-Intent`, on a route that allows one, is not a submission: the endpoint runs right away.

| # | Person | App | Client | Endpoint | Intent | Built from atomic parts | Atomic | Host |
|---|---|---|---|---|---|---|---|---|
| 1 |  |  |  | [E1 QAP endpoint](middleware/layer-1/qap-endpoint.md) finds no intent |  |  |  |  |
| 2 |  |  |  |  |  |  |  | [H2 Current user](middleware/host/current-user.md) says who asks, when the route needs a logged-in user |
| 3 |  |  |  |  |  |  | [A8 Request copy](middleware/layer-4/request-copy.md) copies the request |  |
| 4 |  |  |  |  |  |  |  | [H3 Endpoint](middleware/host/endpoint.md) checks the body, when it has Check |
| 5 |  |  |  |  |  |  | [A7 Problems](middleware/layer-4/problems.md) turns the problems into the answer: 422 |  |
| 6 |  |  |  |  |  |  |  | [H3 Endpoint](middleware/host/endpoint.md) executes the request |
| 7 |  |  |  |  |  |  | [A17 Write result](middleware/layer-4/write-result.md) turns the result into the answer |  |
