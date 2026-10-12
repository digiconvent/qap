import { describe, expect, it } from 'vitest'
import { openEndpoint } from '../src'
import { fakeSend } from './fake_send'
import { view } from './views'

describe('submit', () => {
  it('sends the request body under the chosen policy and lands in the pending state', async () => {
    const fake = fakeSend()
    fake.answer(202, view('pending'))
    const submitted = await openEndpoint({ path: '/api/iam/user/create', method: 'POST', send: fake.send }).submit('peer review', '{"email":"frodo@shire.test"}')

    expect(fake.requests[0]).toMatchObject({ method: 'POST', body: '{"email":"frodo@shire.test"}', headers: { 'x-qap-intent': 'submit', 'x-qap-policy-id': 'peer review' } })
    expect(submitted).toMatchObject({ kind: 'pending', id: 'qap_1', author: 'samwise', permission: 'iam.user.create', counts: { favour: 1, against: 0, withhold: 0 }, comments: [] })
  })

  it('lands in the executed state with the result when one vote is enough', async () => {
    const fake = fakeSend()
    fake.answer(202, view('executed'))
    const submitted = await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).submit('fast track')

    expect(submitted).toMatchObject({ kind: 'executed', result: { status: 204, headers: {}, body: null } })
  })

  it('is refused with the acceptable policies when no policy is given', async () => {
    const fake = fakeSend()
    fake.answer(406, { error: 'QAP policy id header required', hint: 'Specify X-QAP-Policy-ID', acceptable_policies: [{ id: 'peer review', name: 'peer review' }] })
    const submitted = await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).submit('')

    expect(submitted).toEqual({ kind: 'refused', status: 406, message: 'QAP policy id header required', hint: 'Specify X-QAP-Policy-ID', acceptablePolicies: [{ id: 'peer review', name: 'peer review' }], problems: {} })
  })
})

describe('submit on a read', () => {
  it('sends no body at all when there is none, so a GET can be proposed', async () => {
    const fake = fakeSend()
    fake.answer(202, view('pending'))
    await openEndpoint({ path: '/api/shire/red-book', method: 'GET', send: fake.send }).submit('keepers', '')

    expect(fake.requests[0]).toMatchObject({ method: 'GET', body: undefined })
  })
})
