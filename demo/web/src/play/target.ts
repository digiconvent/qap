import type { Pending } from '../../../../client/src'
import type { LaneState, Shared } from '../types'

export function target(state: LaneState, shared: Shared): Pending | null {
  if (state.kind === 'pending') return state
  if (state.kind !== 'listed') return null
  return state.submissions.find((submission) => submission.id === shared.id) ?? state.submissions[0] ?? null
}
