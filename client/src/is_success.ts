import type { Answer } from './types'

export function isSuccess(answer: Answer): boolean {
  return answer.status >= 200 && answer.status < 300
}
