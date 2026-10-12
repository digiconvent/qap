import type { EndpointOptions, Pending, SubmissionState, VoteChoice, WireView } from './types'
import { readFacts } from './read_facts'
import { poll } from './poll'
import { vote } from './vote'
import { comment } from './comment'
import { withdraw } from './withdraw'
import { listen } from './listen'

export function pendingState(options: EndpointOptions, view: WireView): Pending {
  const facts = readFacts(view)
  const state: Pending = {
    kind: 'pending',
    ...facts,
    poll: () => poll(options, facts.id),
    vote: (choice: VoteChoice) => vote(options, facts.id, choice),
    comment: (text: string) => comment(options, facts.id, text),
    withdraw: () => withdraw(options, facts.id),
    listen: (onChange: (state: SubmissionState) => void, everyMs = 5000) => listen(options, facts.id, state, onChange, everyMs),
  }
  return state
}
