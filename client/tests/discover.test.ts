import { describe, expect, it } from 'vitest'
import { openEndpoint, version } from '../src'
import { fakeSend } from './fake_send'

const speaksThisVersion = { 'X-QAP-Version': version }

describe('discover', () => {
  it('asks the endpoint which policies the user may use, and reads their settings', async () => {
    const fake = fakeSend()
    fake.answer(200, {
      endpoint: '/api/iam/user/create',
      permission: 'iam.user.create',
      policies: [{ id: 'peer review', name: 'peer review', favour_threshold: 2, against_threshold: null, quorum: null, deadline_seconds: 60, extension_seconds: 30, withhold_permitted: false }],
    }, speaksThisVersion)
    const discovered = await openEndpoint({ path: '/api/iam/user/create', method: 'POST', send: fake.send }).discover()

    expect(fake.requests[0]).toMatchObject({ url: '/api/iam/user/create', method: 'POST', headers: { 'x-qap-intent': 'discover' } })
    expect(discovered.kind).toBe('discovered')
    if (discovered.kind !== 'discovered') return
    expect(discovered.permission).toBe('iam.user.create')
    expect(discovered.version).toBe(version)
    expect(discovered.policies).toEqual([{ id: 'peer review', name: 'peer review', favourThreshold: 2, againstThreshold: null, quorum: null, deadlineSeconds: 60, extensionSeconds: 30, withholdPermitted: false }])
  })

  it('reads a missing extension as null', async () => {
    const fake = fakeSend()
    fake.answer(200, { endpoint: '/x', permission: 'p', policies: [{ id: 'fast track', name: 'fast track', favour_threshold: 1, against_threshold: null, quorum: null, deadline_seconds: null, withhold_permitted: false }] }, speaksThisVersion)
    const discovered = await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).discover()

    if (discovered.kind !== 'discovered') throw new Error('expected discovered')
    expect(discovered.policies[0].extensionSeconds).toBeNull()
  })

  it('is refused when the server speaks another version, or none', async () => {
    for (const [headers, server] of [[{ 'X-QAP-Version': '9.9.9' }, '9.9.9'], [undefined, 'no version']] as const) {
      const fake = fakeSend()
      fake.answer(200, { endpoint: '/x', permission: 'p', policies: [] }, headers)
      const discovered = await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).discover()

      expect(discovered).toEqual({ kind: 'refused', status: 200, message: `QAP version mismatch: the server speaks ${server}, this client ${version}`, hint: undefined, acceptablePolicies: [], problems: {} })
    }
  })

  it('submits under one of the discovered policies', async () => {
    const fake = fakeSend()
    fake.answer(200, { endpoint: '/x', permission: 'p', policies: [{ id: 'fast track', name: 'fast track', favour_threshold: 1, against_threshold: null, quorum: null, deadline_seconds: null, withhold_permitted: false }] }, speaksThisVersion)
    fake.answer(202, { submission: { id: 'qap_9', endpoint: '/x', method: 'POST', body: '', permission: 'p', policy: { id: 'fast track', name: 'fast track' }, author: 'samwise', when: 1, resolution: 'executed', result: { status: 204, headers: {}, body: null } }, poll: { favour: 1, against: 0, withhold: 0, comments: [], resolution: 'executed' } })
    const discovered = await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).discover()
    if (discovered.kind !== 'discovered') throw new Error('expected discovered')
    const submitted = await discovered.submit(discovered.policies[0])

    expect(fake.requests[1].headers['x-qap-policy-id']).toBe('fast track')
    expect(submitted.kind).toBe('executed')
  })
})
