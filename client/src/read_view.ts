import type { EndpointOptions, SubmissionState, WireView } from './types'
import { pendingState } from './pending_state'
import { decidedState } from './decided_state'

export function readView(options: EndpointOptions, view: WireView): SubmissionState {
  if (view.submission.resolution === 'pending') return pendingState(options, view)
  return decidedState(options, view)
}
