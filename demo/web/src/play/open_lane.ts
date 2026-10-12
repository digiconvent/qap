import { openEndpoint } from '../../../../client/src'
import type { Lane, Scenario } from '../types'
import { sendAs } from './send_as'

export function openLane(scenario: Scenario, actor: string): Lane {
  const send = sendAs(actor)
  const endpoint = openEndpoint({ path: scenario.path, method: scenario.method, send })
  return { actor, endpoint, send, state: { kind: 'idle' }, history: [] }
}
