import type { Discovered, Endpoint, Listed, Send, SubmissionState, VoteChoice } from '../../../client/src'

export type Action =
  | { do: 'discover' }
  | { do: 'list' }
  | { do: 'follow' }
  | { do: 'submit'; policy: string; body?: string }
  | { do: 'poll' }
  | { do: 'vote'; choice: VoteChoice }
  | { do: 'comment'; text: string }
  | { do: 'withdraw' }
  | { do: 'wait' }

export interface Expectation {
  kind: string
  status?: number
  message?: string
  favour?: number
  against?: number
  withhold?: number
  contents?: boolean
  policies?: string[]
  listed?: number
  acceptable?: number
}

export interface Route {
  method: string
  path: string
}

export interface Step {
  actor?: string
  route?: Route
  says: string
  why?: string
  action: Action
  expect: Expectation
}

export interface Scenario {
  id: string
  title: string
  summary: string
  method: string
  path: string
  body: string
  what: string
  rule: string
  deadlineSeconds?: number
  actors: string[]
  steps: Step[]
}

export type LaneState = { kind: 'idle' } | Discovered | Listed | SubmissionState

export interface Lane {
  actor: string
  endpoint: Endpoint
  send: Send
  state: LaneState
  history: string[]
}

export interface Shared {
  id: string | undefined
  deadlineAt: number | undefined
}

export type Outcome = { step: number; problems: string[] }
