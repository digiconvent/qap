import type { EndpointOptions, SubmissionState } from './types'
import { ask } from './ask'
import { readState } from './read_state'

export async function comment(options: EndpointOptions, id: string, text: string): Promise<SubmissionState> {
  return readState(options, await ask(options, { 'X-QAP-Intent': 'comment', 'X-QAP-ID': id, 'X-QAP-Comment': encodeURIComponent(text) }))
}
