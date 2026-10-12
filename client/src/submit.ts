import type { EndpointOptions, SubmissionState } from './types'
import { ask } from './ask'
import { readState } from './read_state'

export async function submit(options: EndpointOptions, policyId: string, body?: string): Promise<SubmissionState> {
  return readState(options, await ask(options, { 'X-QAP-Intent': 'submit', 'X-QAP-Policy-ID': policyId }, body))
}
