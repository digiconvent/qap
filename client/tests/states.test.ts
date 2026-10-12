import { describe, expect, it } from 'vitest'
import { openEndpoint, type Pending } from '../src'
import { fakeSend } from './fake_send'
import { view } from './views'

async function pending() {
  const fake = fakeSend()
  fake.answer(202, view('pending'))
  const state = await openEndpoint({ path: '/x', method: 'DELETE', send: fake.send }).submit('peer review')
  if (state.kind !== 'pending') throw new Error('expected pending')
  return { fake, state: state as Pending }
}

describe('a pending submission', () => {
  it('votes in favour, against or withhold as 1, -1 and 0, on the method of the endpoint', async () => {
    for (const [choice, header] of [['favour', '1'], ['against', '-1'], ['withhold', '0']] as const) {
      const { fake, state } = await pending()
      fake.answer(200, view('pending'))
      await state.vote(choice)
      expect(fake.requests[1]).toMatchObject({ method: 'DELETE', headers: { 'x-qap-intent': 'vote', 'x-qap-id': 'qap_1', 'x-qap-vote': header } })
    }
  })

  it('moves to the state the vote leads to', async () => {
    const { fake, state } = await pending()
    fake.answer(200, view('rejected', {}))
    expect((await state.vote('against')).kind).toBe('rejected')
  })

  it('polls its current state', async () => {
    const { fake, state } = await pending()
    fake.answer(200, view('executed'))
    const polled = await state.poll()
    expect(fake.requests[1].headers).toMatchObject({ 'x-qap-intent': 'poll', 'x-qap-id': 'qap_1' })
    expect(polled.kind).toBe('executed')
  })

  it('comments with the text percent-encoded in a header, and no body', async () => {
    const { fake, state } = await pending()
    const commented = view('pending')
    commented.poll.comments = [{ author: 'frodo', at: 5, text: 'ready — at last' }] as never
    fake.answer(200, commented)
    const after = await state.comment('ready — at last')
    expect(fake.requests[1]).toMatchObject({ body: undefined, headers: { 'x-qap-intent': 'comment', 'x-qap-id': 'qap_1', 'x-qap-comment': 'ready%20%E2%80%94%20at%20last' } })
    expect(after).toMatchObject({ kind: 'pending', comments: [{ author: 'frodo', at: 5, text: 'ready — at last' }] })
  })

  it('comments on a proposed GET, since the text needs no body', async () => {
    const fake = fakeSend()
    fake.answer(202, view('pending', { method: 'GET' }))
    const state = await openEndpoint({ path: '/x', method: 'GET', send: fake.send }).submit('peer review')
    if (state.kind !== 'pending') throw new Error('expected pending')
    fake.answer(200, view('pending', { method: 'GET' }))
    await state.comment('only Bilbo may read it')
    expect(fake.requests[1]).toMatchObject({ method: 'GET', body: undefined })
  })

  it('is withdrawn by its author', async () => {
    const { fake, state } = await pending()
    fake.answer(200)
    const withdrawn = await state.withdraw()
    expect(fake.requests[1].headers).toMatchObject({ 'x-qap-intent': 'delete-submission', 'x-qap-id': 'qap_1' })
    expect(withdrawn).toEqual({ kind: 'withdrawn', id: 'qap_1' })
  })

  it('is refused when someone else tries to withdraw it', async () => {
    const { fake, state } = await pending()
    fake.answer(403, { error: 'Only the author can delete this submission' })
    expect(await state.withdraw()).toEqual({ kind: 'refused', status: 403, message: 'Only the author can delete this submission', hint: undefined, acceptablePolicies: [], problems: {} })
  })

  it('is refused with the problems when the endpoint cannot carry out the request', async () => {
    const fake = fakeSend()
    fake.answer(422, { error: 'the request cannot be carried out', problems: { title: ['is required'] } })
    const refused = await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).submit('peer review', '{}')
    expect(refused).toEqual({ kind: 'refused', status: 422, message: 'the request cannot be carried out', hint: undefined, acceptablePolicies: [], problems: { title: ['is required'] } })
  })
})

describe('a decided submission', () => {
  it('offers only polling once it is executed or rejected', async () => {
    for (const resolution of ['executed', 'rejected']) {
      const fake = fakeSend()
      fake.answer(202, view(resolution))
      const state = await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).submit('p')
      expect(Object.keys(state).filter((key) => typeof (state as unknown as Record<string, unknown>)[key] === 'function')).toEqual(['poll'])
    }
  })
})
