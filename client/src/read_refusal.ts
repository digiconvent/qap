import type { Answer, PolicyReference, Refused } from './types'

export function readRefusal(answer: Answer): Refused {
  const body = (answer.body ?? {}) as { error?: string; hint?: string; acceptable_policies?: PolicyReference[]; problems?: Record<string, string[]> }
  return {
    kind: 'refused',
    status: answer.status,
    message: body.error ?? '',
    hint: body.hint,
    acceptablePolicies: (body.acceptable_policies ?? []).map(({ id, name }) => ({ id, name })),
    problems: body.problems ?? {},
  }
}
