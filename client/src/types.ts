export type Send = (url: string, init: RequestInit) => Promise<Response>

export interface EndpointOptions {
  path: string
  method: string
  send: Send
}

export interface PolicyReference {
  id: string
  name: string
}

export interface PolicySettings extends PolicyReference {
  favourThreshold: number
  againstThreshold: number | null
  quorum: number | null
  deadlineSeconds: number | null
  extensionSeconds: number | null
  withholdPermitted: boolean
}

export interface Comment {
  author: string
  at: number
  text: string
}

export interface Result {
  status: number
  headers: Record<string, string>
  body: unknown
}

export interface Counts {
  favour: number
  against: number
  withhold: number
}

export type VoteChoice = 'favour' | 'against' | 'withhold'

export interface SubmissionFacts {
  id: string
  endpoint: string
  method: string
  body: string
  permission: string
  policy: PolicyReference
  author: string
  when: number
  counts: Counts
  comments: Comment[]
}

export interface Pending extends SubmissionFacts {
  kind: 'pending'
  poll(): Promise<SubmissionState>
  vote(choice: VoteChoice): Promise<SubmissionState>
  comment(text: string): Promise<SubmissionState>
  withdraw(): Promise<SubmissionState>
  listen(onChange: (state: SubmissionState) => void, everyMs?: number): () => void
}

export interface Executed extends SubmissionFacts {
  kind: 'executed'
  result: Result | null
  poll(): Promise<SubmissionState>
}

export interface Rejected extends SubmissionFacts {
  kind: 'rejected'
  result: Result | null
  poll(): Promise<SubmissionState>
}

export interface Withdrawn {
  kind: 'withdrawn'
  id: string
}

export interface Refused {
  kind: 'refused'
  status: number
  message: string
  hint: string | undefined
  acceptablePolicies: PolicyReference[]
  problems: Record<string, string[]>
}

export type SubmissionState = Pending | Executed | Rejected | Withdrawn | Refused

export interface Discovered {
  kind: 'discovered'
  endpoint: string
  permission: string
  version: string
  policies: PolicySettings[]
  submit(policy: PolicyReference, body?: string): Promise<SubmissionState>
}

export interface Listed {
  kind: 'listed'
  submissions: Pending[]
}

export interface Endpoint {
  discover(): Promise<Discovered | Refused>
  submit(policyId: string, body?: string): Promise<SubmissionState>
  list(): Promise<Listed | Refused>
  follow(id: string): Promise<SubmissionState>
}

export interface WireView {
  submission: {
    id: string
    endpoint: string
    method: string
    body: string
    permission: string
    policy: PolicyReference
    author: string
    when: number
    resolution: string
    result: Result | null
  }
  poll: Counts & { comments: Comment[] | null; resolution: string }
}

export interface WireDiscovery {
  endpoint: string
  permission: string
  policies: WirePolicy[]
}

export interface WirePolicy {
  id: string
  name: string
  favour_threshold: number
  against_threshold: number | null
  quorum: number | null
  deadline_seconds: number | null
  extension_seconds?: number | null
  withhold_permitted: boolean
}

export interface Answer {
  status: number
  headers: Headers
  body: unknown
}
