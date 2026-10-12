import type { WireView } from './types'

const resolutions = ['pending', 'executed', 'rejected']

export function isView(body: unknown): body is WireView {
  const { submission, poll } = (body ?? {}) as Partial<WireView>
  return (
    typeof submission === 'object' && submission !== null &&
    typeof submission.id === 'string' &&
    resolutions.includes(submission.resolution) &&
    typeof submission.policy === 'object' && submission.policy !== null &&
    typeof poll === 'object' && poll !== null &&
    typeof poll.favour === 'number' && typeof poll.against === 'number' && typeof poll.withhold === 'number' &&
    (poll.comments === null || Array.isArray(poll.comments))
  )
}
