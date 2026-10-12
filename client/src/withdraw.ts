import type { EndpointOptions, SubmissionState } from './types'
import { ask } from './ask'
import { isSuccess } from './is_success'
import { readRefusal } from './read_refusal'

export async function withdraw(options: EndpointOptions, id: string): Promise<SubmissionState> {
  const answer = await ask(options, { 'X-QAP-Intent': 'delete-submission', 'X-QAP-ID': id })
  if (!isSuccess(answer)) return readRefusal(answer)
  return { kind: 'withdrawn', id }
}
