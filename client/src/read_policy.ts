import type { PolicySettings, WirePolicy } from './types'

export function readPolicy(policy: WirePolicy): PolicySettings {
  return {
    id: policy.id,
    name: policy.name,
    favourThreshold: policy.favour_threshold,
    againstThreshold: policy.against_threshold,
    quorum: policy.quorum,
    deadlineSeconds: policy.deadline_seconds,
    extensionSeconds: policy.extension_seconds ?? null,
    withholdPermitted: policy.withhold_permitted,
  }
}
