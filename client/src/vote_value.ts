import type { VoteChoice } from './types'

export function voteValue(choice: VoteChoice): string {
  if (choice === 'favour') return '1'
  if (choice === 'against') return '-1'
  return '0'
}
