import type { Answer, EndpointOptions } from './types'
import { readBody } from './read_body'

export async function ask(options: EndpointOptions, headers: Record<string, string>, body?: string): Promise<Answer> {
  const { send, path, method } = options
  const response = await send(path, { method, headers, body: body || undefined })
  return { status: response.status, headers: response.headers, body: await readBody(response) }
}
