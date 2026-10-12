import type { WireDiscovery } from './types'

export function isDiscovery(body: unknown): body is WireDiscovery {
  const { permission, policies } = (body ?? {}) as Partial<WireDiscovery>
  return (
    typeof permission === 'string' &&
    Array.isArray(policies) &&
    policies.every((policy) => typeof policy?.id === 'string' && typeof policy.name === 'string' && typeof policy.favour_threshold === 'number')
  )
}
