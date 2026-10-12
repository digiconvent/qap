import type { Action } from '../types'

export function label(action: Action): string {
  if (action.do === 'submit') return 'submit · ' + (action.policy || 'no policy')
  if (action.do === 'vote') return 'vote ' + action.choice
  return action.do
}
