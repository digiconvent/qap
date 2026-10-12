import { describe, expect, it } from 'vitest'
import { openEndpoint } from '../src'
import { fakeSend } from './fake_send'
import { view } from './views'

describe('list', () => {
  it('lists the pending submissions the user may vote on, each in the pending state', async () => {
    const fake = fakeSend()
    fake.answer(200, [view('pending')])
    const listed = await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).list()

    expect(fake.requests[0].headers['x-qap-intent']).toBe('list')
    expect(listed.kind).toBe('listed')
    if (listed.kind !== 'listed') return
    expect(listed.submissions.map((s) => [s.kind, s.id])).toEqual([['pending', 'qap_1']])
  })
})

describe('follow', () => {
  it('picks up a submission by its id', async () => {
    const fake = fakeSend()
    fake.answer(200, view('pending'))
    const followed = await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).follow('qap_1')

    expect(fake.requests[0].headers).toMatchObject({ 'x-qap-intent': 'poll', 'x-qap-id': 'qap_1' })
    expect(followed.kind).toBe('pending')
  })

  it('is refused for an unknown id', async () => {
    const fake = fakeSend()
    fake.answer(404, { error: 'Submission state does not exist' })
    expect(await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).follow('qap_404')).toMatchObject({ kind: 'refused', status: 404 })
  })
})
