import type { SubmissionFacts, WireView } from './types'

export function readFacts(view: WireView): SubmissionFacts {
  const { submission, poll } = view
  return {
    id: submission.id,
    endpoint: submission.endpoint,
    method: submission.method,
    body: submission.body,
    permission: submission.permission,
    policy: { id: submission.policy.id, name: submission.policy.name },
    author: submission.author,
    when: submission.when,
    counts: { favour: poll.favour, against: poll.against, withhold: poll.withhold },
    comments: poll.comments ?? [],
  }
}
