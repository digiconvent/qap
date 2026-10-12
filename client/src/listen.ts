import type { EndpointOptions, SubmissionState } from './types'
import { poll } from './poll'

export function listen(options: EndpointOptions, id: string, current: SubmissionState, onChange: (state: SubmissionState) => void, everyMs: number): () => void {
  let last = JSON.stringify(current)
  let stopped = false
  let timer = setTimeout(check, everyMs)

  async function check() {
    const next = await poll(options, id).catch(() => null)
    if (stopped) return
    if (next !== null) {
      const seen = JSON.stringify(next)
      if (seen !== last) {
        last = seen
        onChange(next)
      }
      if (next.kind !== 'pending') return
    }
    timer = setTimeout(check, everyMs)
  }

  return () => {
    stopped = true
    clearTimeout(timer)
  }
}
