import type { Endpoint } from '../../../../client/src'
import type { Action, Lane, LaneState, Scenario, Shared } from '../types'
import { target } from './target'
import { notAvailable } from './not_available'

export async function perform(lane: Lane, action: Action, scenario: Scenario, shared: Shared, endpoint: Endpoint = lane.endpoint): Promise<LaneState> {
  const { state } = lane
  switch (action.do) {
    case 'discover':
      return endpoint.discover()
    case 'list':
      return endpoint.list()
    case 'follow':
      return shared.id ? endpoint.follow(shared.id) : notAvailable('there is no submission to follow yet')
    case 'submit': {
      const policy = state.kind === 'discovered' ? state.policies.find((p) => p.id === action.policy) : undefined
      const body = action.body ?? scenario.body
      if (state.kind === 'discovered' && policy) return state.submit(policy, body)
      return endpoint.submit(action.policy, body)
    }
    case 'wait':
      return state
    case 'poll':
      if (state.kind === 'pending' || state.kind === 'executed' || state.kind === 'rejected') return state.poll()
      return notAvailable('poll is not available while ' + state.kind)
  }
  const pending = target(state, shared)
  if (!pending) return notAvailable(action.do + ' is not available while ' + state.kind)
  if (action.do === 'vote') return pending.vote(action.choice)
  if (action.do === 'comment') return pending.comment(action.text)
  return pending.withdraw()
}
