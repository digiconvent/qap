import type { Exchange } from '../play/exchanges'
import { escape } from './escape'
import { pretty } from './pretty'

export function renderExchange(exchange: Exchange): string {
  const request = [`${exchange.method} ${exchange.url}`, ...exchange.headers.map(([name, value]) => `${name}: ${value}`)].join('\n')
  const response = [`${exchange.status} ${exchange.statusText}`, ...exchange.responseHeaders.map(([name, value]) => `${name}: ${value}`)].join('\n')
  return `<div class="wire">
    <h4>→ request</h4><pre>${escape(request)}${exchange.body ? '\n\n' + escape(pretty(exchange.body)) : ''}</pre>
    <h4>← response</h4><pre>${escape(response)}${exchange.responseBody ? '\n\n' + escape(pretty(exchange.responseBody)) : '\n\n(empty body)'}</pre>
  </div>`
}
