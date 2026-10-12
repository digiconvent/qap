import type { Action } from '../types'

export function verb(action: Action): string {
  switch (action.do) {
    case 'discover': return 'ask which policies apply'
    case 'list': return 'see what waits'
    case 'follow': return 'open by id'
    case 'submit': return 'propose under ' + (action.policy || 'no policy')
    case 'poll': return 'check status'
    case 'vote': return { favour: 'vote for', against: 'vote against', withhold: 'abstain' }[action.choice]
    case 'comment': return 'comment'
    case 'withdraw': return 'withdraw'
    case 'wait': return 'wait'
  }
}
