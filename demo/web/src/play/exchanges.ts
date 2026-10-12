export interface Exchange {
  method: string
  url: string
  headers: [string, string][]
  body: string
  status: number
  statusText: string
  responseHeaders: [string, string][]
  responseBody: string
}

const recorded: Exchange[] = []

export function record(exchange: Exchange) {
  recorded.push(exchange)
}

export function takeExchanges(): Exchange[] {
  return recorded.splice(0, recorded.length)
}
