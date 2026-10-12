import type { Refused } from '../../../../client/src'

export function notAvailable(message: string): Refused {
  return { kind: 'refused', status: 0, message, hint: 'the state machine does not offer this move', acceptablePolicies: [], problems: {} }
}
