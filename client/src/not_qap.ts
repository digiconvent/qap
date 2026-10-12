import type { Answer, Refused } from './types'

export function notQap(answer: Answer): Refused {
  return { kind: 'refused', status: answer.status, message: 'not a QAP answer', hint: undefined, acceptablePolicies: [], problems: {} }
}
