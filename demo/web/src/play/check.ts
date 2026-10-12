import type { Expectation, LaneState } from '../types'

export function check(state: LaneState, expect: Expectation): string[] {
  const problems: string[] = []
  const compare = (what: string, expected: unknown, actual: unknown) => {
    if (expected !== undefined && JSON.stringify(expected) !== JSON.stringify(actual)) problems.push(`${what}: expected ${JSON.stringify(expected)}, got ${JSON.stringify(actual)}`)
  }
  compare('state', expect.kind, state.kind)
  if (state.kind === 'refused') {
    compare('status', expect.status, state.status)
    compare('message', expect.message, state.message)
    compare('acceptable policies', expect.acceptable, state.acceptablePolicies.length)
  }
  if (state.kind === 'pending' || state.kind === 'executed' || state.kind === 'rejected') {
    compare('favour', expect.favour, state.counts.favour)
    compare('against', expect.against, state.counts.against)
    compare('withhold', expect.withhold, state.counts.withhold)
  }
  if (state.kind === 'executed' && expect.contents !== undefined) compare('answer visible', expect.contents, state.result?.body != null)
  if (state.kind === 'discovered') compare('policies', expect.policies, state.policies.map((p) => p.name))
  if (state.kind === 'listed') compare('listed', expect.listed, state.submissions.length)
  return problems
}
