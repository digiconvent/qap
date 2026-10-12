export function view(resolution: string, overrides: Record<string, unknown> = {}) {
  return {
    submission: {
      id: 'qap_1',
      endpoint: '/api/iam/user/create',
      method: 'POST',
      body: '{"email":"frodo@shire.test"}',
      permission: 'iam.user.create',
      policy: { id: 'peer review', name: 'peer review' },
      author: 'samwise',
      when: 1760000000,
      resolution,
      result: resolution === 'executed' ? { status: 204, headers: {}, body: null } : null,
      ...overrides,
    },
    poll: { favour: 1, against: 0, withhold: 0, comments: [], resolution },
  }
}
