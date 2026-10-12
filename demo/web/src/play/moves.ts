import type { Action, LaneState, Shared } from '../types'
import { target } from './target'

const reviewing: Action[] = [
  { do: 'vote', choice: 'favour' },
  { do: 'vote', choice: 'against' },
  { do: 'vote', choice: 'withhold' },
  { do: 'comment', text: '' },
  { do: 'withdraw' },
]

export function moves(state: LaneState, shared: Shared): Action[] {
  if (state.kind === 'discovered') return state.policies.map((p) => ({ do: 'submit', policy: p.id }))
  if (state.kind === 'refused') return state.acceptablePolicies.map((p) => ({ do: 'submit', policy: p.id }))
  if (state.kind === 'listed') return target(state, shared) ? reviewing : []
  if (state.kind === 'pending') return [{ do: 'poll' }, ...reviewing]
  if (state.kind === 'executed' || state.kind === 'rejected') return [{ do: 'poll' }]
  return []
}
