import type { EndpointOptions, SubmissionState } from './types'
import { poll } from './poll'

export function follow(options: EndpointOptions, id: string): Promise<SubmissionState> {
  return poll(options, id)
}
