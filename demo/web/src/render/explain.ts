import type { LaneState } from '../types'

export function explain(state: LaneState): string {
  switch (state.kind) {
    case 'idle': return 'has not done anything yet'
    case 'discovered': return state.policies.length ? 'may use ' + state.policies.map((p) => p.name).join(' or ') : 'no policy includes them'
    case 'listed': return state.submissions.length === 1 ? 'sees 1 proposal waiting' : `sees ${state.submissions.length || 'no'} proposals waiting`
    case 'pending': return `pending · ${state.counts.favour} for, ${state.counts.against} against`
    case 'executed': return `carried out · ${state.counts.favour} for, ${state.counts.against} against${state.result?.body != null ? ' · the answer is attached' : state.result?.status === 200 ? ' · the answer is hidden from them' : ''}`
    case 'rejected': return `rejected · ${state.counts.favour} for, ${state.counts.against} against`
    case 'withdrawn': return 'the proposal is withdrawn'
    case 'refused': return `refused: ${state.message}${state.status ? ` (${state.status})` : ''}`
  }
}
