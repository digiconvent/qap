import type { EndpointOptions, Listed, Pending, Refused } from './types'
import { ask } from './ask'
import { isSuccess } from './is_success'
import { isView } from './is_view'
import { notQap } from './not_qap'
import { readRefusal } from './read_refusal'
import { pendingState } from './pending_state'

export async function list(options: EndpointOptions): Promise<Listed | Refused> {
  const answer = await ask(options, { 'X-QAP-Intent': 'list' })
  if (!isSuccess(answer)) return readRefusal(answer)
  const views = answer.body
  if (!Array.isArray(views) || !views.every((view) => isView(view) && view.submission.resolution === 'pending')) return notQap(answer)
  const submissions: Pending[] = views.map((view) => pendingState(options, view))
  return { kind: 'listed', submissions }
}
