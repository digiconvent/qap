import { describe, expect, it } from 'vitest'
import { openEndpoint, type Send } from '../src'

describe('send', () => {
  it('calls send on its own, so the browser fetch can be passed as it is', async () => {
    let receiver: unknown = 'not called'
    const send = async function (this: unknown) {
      receiver = this
      return new Response(JSON.stringify({ endpoint: '/x', permission: 'p', policies: [] }), { status: 200 })
    } as Send
    await openEndpoint({ path: '/x', method: 'POST', send }).discover()

    expect(receiver).toBeUndefined()
  })
})
