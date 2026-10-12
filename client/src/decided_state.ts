import type { EndpointOptions, Executed, Rejected, WireView } from './types'
import { readFacts } from './read_facts'
import { poll } from './poll'

export function decidedState(options: EndpointOptions, view: WireView): Executed | Rejected {
  const facts = readFacts(view)
  return {
    kind: view.submission.resolution === 'executed' ? 'executed' : 'rejected',
    ...facts,
    result: view.submission.result,
    poll: () => poll(options, facts.id),
  }
}
