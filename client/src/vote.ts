import type { EndpointOptions, SubmissionState, VoteChoice } from './types'
import { ask } from './ask'
import { readState } from './read_state'
import { voteValue } from './vote_value'

export async function vote(options: EndpointOptions, id: string, choice: VoteChoice): Promise<SubmissionState> {
  return readState(options, await ask(options, { 'X-QAP-Intent': 'vote', 'X-QAP-ID': id, 'X-QAP-Vote': voteValue(choice) }))
}
