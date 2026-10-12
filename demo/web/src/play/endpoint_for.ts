import { openEndpoint, type Endpoint } from '../../../../client/src'
import type { Lane, Step } from '../types'

export function endpointFor(lane: Lane, step: Step): Endpoint {
  if (!step.route) return lane.endpoint
  return openEndpoint({ path: step.route.path, method: step.route.method, send: lane.send })
}
