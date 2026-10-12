import type { Endpoint, EndpointOptions } from './types'
import { discover } from './discover'
import { submit } from './submit'
import { list } from './list'
import { follow } from './follow'

export function openEndpoint(options: EndpointOptions): Endpoint {
  return {
    discover: () => discover(options),
    submit: (policyId: string, body?: string) => submit(options, policyId, body),
    list: () => list(options),
    follow: (id: string) => follow(options, id),
  }
}
