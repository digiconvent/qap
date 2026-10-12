import type { EndpointOptions, SubmissionState } from './types'
import { ask } from './ask'
import { readState } from './read_state'

export async function poll(options: EndpointOptions, id: string): Promise<SubmissionState> {
  return readState(options, await ask(options, { 'X-QAP-Intent': 'poll', 'X-QAP-ID': id }))
}
