import { describe, expect, it, vi } from 'vitest'
import { openEndpoint, type Send, type SubmissionState } from '../src'
import { fakeSend } from './fake_send'
import { view } from './views'

const refusedWithoutBody = { kind: 'refused', status: 500, message: '', hint: undefined, acceptablePolicies: [], problems: {} }

describe('answers at the edges', () => {
  it('reads a refusal without a body as an empty message', async () => {
    const fake = fakeSend()
    fake.answer(500)
    expect(await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).discover()).toEqual(refusedWithoutBody)
  })

  it('reads a refusal that is not json as its plain text', async () => {
    const send: Send = async () => new Response('the server fell over', { status: 502 })
    expect(await openEndpoint({ path: '/x', method: 'POST', send }).list()).toEqual({ kind: 'refused', status: 502, message: 'the server fell over', hint: undefined, acceptablePolicies: [], problems: {} })
  })

  it('refuses a list that the server refuses', async () => {
    const fake = fakeSend()
    fake.answer(401, { error: 'unauthorized' })
    expect(await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).list()).toMatchObject({ kind: 'refused', status: 401, message: 'unauthorized' })
  })

  it('reads a discovery without a body as not a QAP answer', async () => {
    const fake = fakeSend()
    fake.answer(200)
    expect(await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).discover()).toMatchObject({ kind: 'refused', message: 'not a QAP answer' })
  })

  it('reads comments sent as null as no comments', async () => {
    const fake = fakeSend()
    const withoutComments = view('pending')
    withoutComments.poll.comments = null as never
    fake.answer(200, withoutComments)
    expect(await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).follow('qap_1')).toMatchObject({ kind: 'pending', comments: [] })
  })

  it('polls a decided submission again', async () => {
    const fake = fakeSend()
    fake.answer(200, view('executed'))
    fake.answer(200, view('executed'))
    const decided = await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).follow('qap_1')
    if (decided.kind !== 'executed') throw new Error('expected executed')
    expect((await decided.poll()).kind).toBe('executed')
    expect(fake.requests[1].headers).toMatchObject({ 'x-qap-intent': 'poll', 'x-qap-id': 'qap_1' })
  })

  it('ignores the answer to a poll that was under way when the listening stopped', async () => {
    vi.useFakeTimers()
    let answer: (response: Response) => void = () => {}
    const send: Send = vi.fn()
      .mockResolvedValueOnce(new Response(JSON.stringify(view('pending')), { status: 202 }))
      .mockImplementationOnce(() => new Promise<Response>((resolve) => { answer = resolve }))
    const proposed = await openEndpoint({ path: '/x', method: 'POST', send }).submit('peer review')
    if (proposed.kind !== 'pending') throw new Error('expected pending')
    const seen: SubmissionState[] = []
    const stop = proposed.listen((next) => seen.push(next), 1000)

    await vi.advanceTimersByTimeAsync(1000)
    stop()
    answer(new Response(JSON.stringify(view('executed')), { status: 200 }))
    await vi.advanceTimersByTimeAsync(5000)
    vi.useRealTimers()
    expect(seen).toHaveLength(0)
  })
})
