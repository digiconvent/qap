import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'
import { openEndpoint, type Pending, type SubmissionState } from '../src'
import { fakeSend } from './fake_send'
import { view } from './views'

async function pending() {
  const fake = fakeSend()
  fake.answer(202, view('pending'))
  const state = await openEndpoint({ path: '/x', method: 'POST', send: fake.send }).submit('peer review')
  if (state.kind !== 'pending') throw new Error('expected pending')
  return { fake, state: state as Pending }
}

function withVotes(favour: number) {
  const changed = view('pending')
  changed.poll.favour = favour
  return changed
}

describe('listening to a pending submission', () => {
  beforeEach(() => vi.useFakeTimers())
  afterEach(() => vi.useRealTimers())

  it('polls every interval, notifies only on a change, and stops once it is decided', async () => {
    const { fake, state } = await pending()
    const seen: SubmissionState[] = []
    fake.answer(200, view('pending'))
    fake.answer(200, withVotes(2))
    fake.answer(200, view('executed'))
    state.listen((next) => seen.push(next), 1000)

    await vi.advanceTimersByTimeAsync(1000)
    expect(seen).toHaveLength(0)
    await vi.advanceTimersByTimeAsync(1000)
    expect(seen.map((s) => s.kind === 'pending' && s.counts.favour)).toEqual([2])
    await vi.advanceTimersByTimeAsync(1000)
    expect(seen.map((s) => s.kind)).toEqual(['pending', 'executed'])
    await vi.advanceTimersByTimeAsync(5000)
    expect(fake.requests).toHaveLength(4)
    expect(fake.requests[1].headers).toMatchObject({ 'x-qap-intent': 'poll', 'x-qap-id': 'qap_1' })
  })

  it('polls every five seconds by default', async () => {
    const { fake, state } = await pending()
    fake.answer(200, view('executed'))
    const seen: SubmissionState[] = []
    state.listen((next) => seen.push(next))

    await vi.advanceTimersByTimeAsync(4999)
    expect(seen).toHaveLength(0)
    await vi.advanceTimersByTimeAsync(1)
    expect(seen.map((s) => s.kind)).toEqual(['executed'])
  })

  it('keeps listening when a poll gets no answer', async () => {
    const { fake, state } = await pending()
    const seen: SubmissionState[] = []
    state.listen((next) => seen.push(next), 1000)

    await vi.advanceTimersByTimeAsync(1000)
    fake.answer(200, view('rejected'))
    await vi.advanceTimersByTimeAsync(1000)
    expect(seen.map((s) => s.kind)).toEqual(['rejected'])
  })

  it('stops polling when told to', async () => {
    const { fake, state } = await pending()
    const seen: SubmissionState[] = []
    fake.answer(200, withVotes(3))
    const stop = state.listen((next) => seen.push(next), 1000)

    await vi.advanceTimersByTimeAsync(999)
    stop()
    await vi.advanceTimersByTimeAsync(5000)
    expect(seen).toHaveLength(0)
    expect(fake.requests).toHaveLength(1)
  })
})
