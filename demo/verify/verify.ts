import { scenarios } from '../web/src/stories/scenarios'
import { openLane } from '../web/src/play/open_lane'
import { perform } from '../web/src/play/perform'
import { endpointFor } from '../web/src/play/endpoint_for'
import { check } from '../web/src/play/check'
import { explain } from '../web/src/render/explain'
import type { Scenario, Shared } from '../web/src/types'

const base = process.env.DEMO_URL ?? 'http://localhost:4280'
const original = globalThis.fetch

function fetchFromDemo(input: string | URL | Request, init?: RequestInit) {
  const url = typeof input === 'string' && input.startsWith('/') ? base + input : input
  return original(url, init)
}

function sleep(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, Math.max(0, ms)))
}

async function play(scenario: Scenario): Promise<number> {
  await fetch('/demo/reset', { method: 'POST' })
  const lanes = scenario.actors.map((actor) => openLane(scenario, actor))
  const shared: Shared = { id: undefined, deadlineAt: undefined }
  let failures = 0
  console.log(`\n${scenario.id}. ${scenario.title}`)
  for (const step of scenario.steps) {
    if (step.action.do === 'wait') {
      await sleep((shared.deadlineAt ?? Date.now()) + 1500 - Date.now())
      console.log(`  ⏳ ${step.says}`)
      continue
    }
    const lane = lanes.find((l) => l.actor === step.actor)!
    const next = await perform(lane, step.action, scenario, shared, endpointFor(lane, step))
    if ('counts' in next) {
      shared.id ??= next.id
      if (scenario.deadlineSeconds) shared.deadlineAt ??= (next.when + scenario.deadlineSeconds) * 1000
    }
    lane.state = next
    const problems = check(next, step.expect)
    if (problems.length) failures++
    console.log(`  ${problems.length ? '✗' : '✓'} ${step.says} → ${explain(next)}${problems.length ? `  (${problems.join('; ')})` : ''}`)
  }
  return failures
}

globalThis.fetch = fetchFromDemo as typeof fetch
let failures = 0
for (const scenario of scenarios) failures += await play(scenario)
console.log(failures ? `\n${failures} steps did not go as the stories say` : '\nevery story played out as written')
process.exit(failures ? 1 : 0)
