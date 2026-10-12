import type { Action, LaneState } from '../types'

export function call(action: Action, before: LaneState): string {
  const holder = before.kind === 'listed' ? 'listed.submissions[0]' : before.kind
  switch (action.do) {
    case 'discover': return 'endpoint.discover()'
    case 'list': return 'endpoint.list()'
    case 'follow': return 'endpoint.follow(id)'
    case 'submit': return before.kind === 'discovered' ? `discovered.submit(policies['${action.policy}'])` : `endpoint.submit('${action.policy}')`
    case 'poll': return `${holder}.poll()`
    case 'vote': return `${holder}.vote('${action.choice}')`
    case 'comment': return `${holder}.comment('${action.text}')`
    case 'withdraw': return `${holder}.withdraw()`
    case 'wait': return ''
  }
}
