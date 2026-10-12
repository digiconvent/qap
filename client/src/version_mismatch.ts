import type { Answer, Refused } from './types'
import { version } from './version'

export function versionMismatch(answer: Answer): Refused | null {
  const server = answer.headers.get('X-QAP-Version')
  if (server === version) return null
  return {
    kind: 'refused',
    status: answer.status,
    message: `QAP version mismatch: the server speaks ${server ?? 'no version'}, this client ${version}`,
    hint: undefined,
    acceptablePolicies: [],
    problems: {},
  }
}
