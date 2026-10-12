import type { Send } from '../../../../client/src'
import { record } from './exchanges'

export function sendAs(actor: string): Send {
  return async function send(url, init) {
    const headers = new Headers(init.headers)
    headers.set('X-User', actor)
    const response = await fetch(url, { ...init, headers })
    record({
      method: init.method ?? 'GET',
      url,
      headers: [...headers.entries()],
      body: typeof init.body === 'string' ? init.body : '',
      status: response.status,
      statusText: response.statusText,
      responseHeaders: [...response.headers.entries()].filter(([name]) => name === 'content-type'),
      responseBody: await response.clone().text(),
    })
    return response
  }
}
