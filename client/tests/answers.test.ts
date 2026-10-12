import { describe, expect, it } from 'vitest'
import { openEndpoint } from '../src'
import { fakeSend } from './fake_send'
import { view } from './views'

const notQap = { kind: 'refused', status: 200, message: 'not a QAP answer', hint: undefined, acceptablePolicies: [], problems: {} }

function opened() {
  const fake = fakeSend()
  return { fake, endpoint: openEndpoint({ path: '/x', method: 'POST', send: fake.send }) }
}

describe('a success that is not a QAP answer', () => {
  it('is refused when following a submission, as when an unguarded route answers', async () => {
    for (const body of [undefined, 'reached', { ok: true }, { submission: { id: 'qap_1' } }]) {
      const { fake, endpoint } = opened()
      fake.answer(200, body)
      expect(await endpoint.follow('qap_1')).toEqual(notQap)
    }
  })

  it('is refused when its resolution is one this client does not know, instead of being read as rejected', async () => {
    const { fake, endpoint } = opened()
    fake.answer(200, view('expired'))
    expect(await endpoint.follow('qap_1')).toEqual(notQap)
  })

  it('is refused when discovering', async () => {
    const { fake, endpoint } = opened()
    fake.answer(200, { permission: 'p', policies: [{ id: 'peer review' }] })
    expect(await endpoint.discover()).toEqual(notQap)
  })

  it('is refused when listing, also when the list holds a decided submission', async () => {
    for (const body of [{ submissions: [] }, [view('executed')]]) {
      const { fake, endpoint } = opened()
      fake.answer(200, body)
      expect(await endpoint.list()).toEqual(notQap)
    }
  })
})
