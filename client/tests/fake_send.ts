import type { Send } from '../src'

export interface SentRequest {
  url: string
  method: string
  headers: Record<string, string>
  body: string | undefined
}

export interface FakeSend {
  send: Send
  requests: SentRequest[]
  answer(status: number, body?: unknown, headers?: Record<string, string>): void
}

export function fakeSend(): FakeSend {
  const requests: SentRequest[] = []
  const answers: Response[] = []
  const send: Send = async (url, init) => {
    requests.push({
      url,
      method: init.method ?? 'GET',
      headers: Object.fromEntries(new Headers(init.headers).entries()),
      body: init.body === undefined || init.body === null ? undefined : String(init.body),
    })
    const response = answers.shift()
    if (!response) throw new Error('no answer prepared for ' + url)
    return response
  }
  const answer = (status: number, body?: unknown, headers?: Record<string, string>) => {
    answers.push(new Response(body === undefined ? null : JSON.stringify(body), { status, headers }))
  }
  return { send, requests, answer }
}
