import type { Answer, EndpointOptions, SubmissionState } from './types'
import { isSuccess } from './is_success'
import { isView } from './is_view'
import { notQap } from './not_qap'
import { readRefusal } from './read_refusal'
import { readView } from './read_view'

export function readState(options: EndpointOptions, answer: Answer): SubmissionState {
  if (!isSuccess(answer)) return readRefusal(answer)
  if (!isView(answer.body)) return notQap(answer)
  return readView(options, answer.body)
}
