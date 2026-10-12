import type { Discovered, EndpointOptions, PolicyReference, Refused } from './types'
import { ask } from './ask'
import { isDiscovery } from './is_discovery'
import { isSuccess } from './is_success'
import { notQap } from './not_qap'
import { readRefusal } from './read_refusal'
import { readPolicy } from './read_policy'
import { submit } from './submit'
import { version } from './version'
import { versionMismatch } from './version_mismatch'

export async function discover(options: EndpointOptions): Promise<Discovered | Refused> {
  const answer = await ask(options, { 'X-QAP-Intent': 'discover' })
  if (!isSuccess(answer)) return readRefusal(answer)
  if (!isDiscovery(answer.body)) return notQap(answer)
  const mismatch = versionMismatch(answer)
  if (mismatch) return mismatch
  const body = answer.body
  return {
    kind: 'discovered',
    endpoint: body.endpoint,
    permission: body.permission,
    version,
    policies: body.policies.map(readPolicy),
    submit: (policy: PolicyReference, requestBody?: string) => submit(options, policy.id, requestBody),
  }
}
