import type { SubmissionFacts } from '../../../client/src'
import type { Lane, LaneState, Scenario, Shared, Step } from './types'
import { scenarios } from './stories/scenarios'
import { groups, names } from './stories/cast'
import { openLane } from './play/open_lane'
import { perform } from './play/perform'
import { endpointFor } from './play/endpoint_for'
import { moves } from './play/moves'
import { check } from './play/check'
import { verb } from './render/verb'
import { call } from './render/call'
import { explain } from './render/explain'
import { escape } from './render/escape'
import { takeExchanges, type Exchange } from './play/exchanges'
import { renderExchange } from './render/render_exchange'

interface Done {
  before: LaneState
  after: LaneState
  problems: string[]
  exchanges: Exchange[]
}

interface Proposal {
  facts: SubmissionFacts
  status: string
}

interface Decision {
  resolution: string
  policy: { name: string }
  favour: string[] | null
  against: string[] | null
  withhold: string[] | null
  comments?: { author: string; text: string }[]
}

const play = {
  scenario: scenarios[0] as Scenario,
  lanes: [] as Lane[],
  cursor: 0,
  shared: { id: undefined, deadlineAt: undefined } as Shared,
  done: [] as Done[],
  proposal: undefined as Proposal | undefined,
  executed: 0,
  decisions: [] as Decision[],
  opened: new Set<number>(),
  waitingUntil: undefined as number | undefined,
  busy: false,
}

const element = (id: string) => document.getElementById(id) as HTMLElement
const nameOf = (actor: string) => names[actor] ?? actor.charAt(0).toUpperCase() + actor.slice(1)
const sleep = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms))

async function load(scenario: Scenario) {
  await fetch('/demo/reset', { method: 'POST' })
  Object.assign(play, {
    scenario,
    lanes: scenario.actors.map((actor) => openLane(scenario, actor)),
    cursor: 0,
    shared: { id: undefined, deadlineAt: undefined },
    done: [],
    proposal: undefined,
    executed: 0,
    decisions: [],
    opened: new Set<number>(),
    waitingUntil: undefined,
  })
  render()
}

function remember(after: LaneState) {
  if ('counts' in after) {
    play.shared.id ??= after.id
    if (play.scenario.deadlineSeconds && play.shared.deadlineAt === undefined) play.shared.deadlineAt = (after.when + play.scenario.deadlineSeconds) * 1000
    play.proposal = { facts: after, status: after.kind }
  }
  if (after.kind === 'withdrawn' && play.proposal) play.proposal.status = 'withdrawn'
}

async function refreshRecord() {
  const record = await (await fetch('/demo/record')).json()
  play.executed = record.executed.length
  play.decisions = record.audited
}

async function waitForDeadline() {
  play.waitingUntil = (play.shared.deadlineAt ?? Date.now()) + 1500
  render()
  while (Date.now() < play.waitingUntil) {
    await sleep(250)
    element('proposal').innerHTML = renderProposal()
  }
  play.waitingUntil = undefined
}

async function step() {
  const current = play.scenario.steps[play.cursor]
  if (!current || play.busy) return
  play.busy = true
  if (current.action.do === 'wait') {
    await waitForDeadline()
    play.done.push({ before: { kind: 'idle' }, after: { kind: 'idle' }, problems: [], exchanges: [] })
  } else {
    const lane = play.lanes.find((l) => l.actor === current.actor) as Lane
    const before = lane.state
    takeExchanges()
    const after = await perform(lane, current.action, play.scenario, play.shared, endpointFor(lane, current))
    lane.state = after
    remember(after)
    play.done.push({ before, after, problems: check(after, current.expect), exchanges: takeExchanges() })
  }
  await refreshRecord()
  play.cursor++
  play.busy = false
  render()
}

async function playAll() {
  while (play.cursor < play.scenario.steps.length) {
    await step()
    await sleep(1200)
  }
}

function badge(kind: string): string {
  return `<span class="badge badge-${kind}">${kind === 'idle' ? 'not started' : kind}</span>`
}

function renderClock(): string {
  const { deadlineAt } = play.shared
  if (!deadlineAt || play.proposal?.status !== 'pending') return ''
  const left = Math.max(0, Math.ceil((deadlineAt - Date.now()) / 1000))
  return `<p class="clock">⏳ ${left ? `${left} seconds left to decide` : 'the deadline has passed'}</p>`
}

function renderDecision(decision: Decision): string {
  const who = (label: string, list: string[] | null) => (list?.length ? `<span>${label}: ${list.map(nameOf).map(escape).join(', ')}</span>` : '')
  const comments = (decision.comments ?? []).map((c) => `<li>${escape(nameOf(c.author))}: “${escape(c.text)}”</li>`).join('')
  return `<div class="decision">${badge(decision.resolution)} under ${escape(decision.policy.name)}
    <p class="who">${who('for', decision.favour)}${who('against', decision.against)}${who('abstained', decision.withhold)}</p>
    ${comments ? `<ul>${comments}</ul>` : ''}</div>`
}

function renderProposal(): string {
  const { scenario, proposal } = play
  const status = proposal ? badge(proposal.status) : '<span class="badge">not proposed yet</span>'
  const counts = proposal
    ? `<div class="tally"><span><b>${proposal.facts.counts.favour}</b> for</span><span><b>${proposal.facts.counts.against}</b> against</span><span><b>${proposal.facts.counts.withhold}</b> abstained</span></div>`
    : ''
  const by = proposal ? ` · proposed by <b>${escape(nameOf(proposal.facts.author))}</b>` : ''
  const ran = play.executed ? '<p class="ran">✓ the server carried out the request</p>' : ''
  const chronicle = play.decisions.length ? '<h3>written in the record</h3>' + play.decisions.map(renderDecision).join('') : ''
  return `<div class="proposal-head"><div><h3>the proposal</h3><p class="what">${escape(scenario.what)}${by}</p>
    <p class="muted"><code>${escape(scenario.method)} ${escape(scenario.path)}</code>${scenario.body ? ` <code>${escape(scenario.body)}</code>` : ''}</p></div>${status}</div>
    <p class="rule">${escape(scenario.rule)}</p>${counts}${renderClock()}${ran}${chronicle}`
}

function renderHeader(lane: Lane): string {
  const available = moves(lane.state, play.shared)
  return `<div class="cell head"><h2>${escape(nameOf(lane.actor))}</h2><p class="muted">${escape(groups[lane.actor] ?? '')}</p>
    <p>${badge(lane.state.kind)}</p><p class="explain">${escape(explain(lane.state))}</p>
    <p class="can">${available.length ? 'can now: ' + available.map((a) => `<span class="chip">${escape(verb(a))}</span>`).join('') : 'can start from the endpoint: ask, propose, see what waits'}</p></div>`
}

function renderPayload(index: number, done: Done): string {
  const open = index === play.cursor - 1 || play.opened.has(index) ? 'open' : ''
  const body = done.exchanges.length ? done.exchanges.map(renderExchange).join('') : '<p class="muted">no request was sent: the client does not offer this move in this state</p>'
  return `<details class="payload" data-step="${index}" ${open}><summary>payload</summary>${body}</details>`
}

function renderRoute(s: Step): string {
  return s.route ? `<p class="muted">on <code>${escape(s.route.method)} ${escape(s.route.path)}</code></p>` : ''
}

function renderCard(s: Step, index: number): string {
  const done = play.done[index]
  if (!done) {
    const current = index === play.cursor
    const lane = play.lanes.find((l) => l.actor === s.actor) as Lane
    return `<div class="card ${current ? 'current' : 'future'}"><p class="says">${escape(s.says)}</p>${renderRoute(s)}
      <p><code>${escape(call(s.action, current ? lane.state : { kind: 'idle' }))}</code></p>
      ${current ? '<button class="run">run this step</button>' : ''}</div>`
  }
  const ok = done.problems.length === 0
  return `<div class="card done ${ok ? '' : 'bad'}"><p class="says">${escape(s.says)}</p>${renderRoute(s)}
    <p><code>${escape(call(s.action, done.before))}</code></p>
    <p class="answer">${badge(done.after.kind)} ${escape(explain(done.after))}</p>
    ${s.why ? `<p class="why">${escape(s.why)}</p>` : ''}
    ${renderPayload(index, done)}
    ${ok ? '' : `<p class="fail">✗ this did not go as the story says: ${escape(done.problems.join('; '))}</p>`}</div>`
}

function renderTimeCard(s: Step, index: number): string {
  const state = index < play.cursor ? 'done' : index === play.cursor ? 'current' : 'future'
  const action = state === 'current' ? (play.waitingUntil ? '<p class="clock">waiting…</p>' : '<button class="run">let the time pass</button>') : ''
  return `<div class="card time ${state}"><p class="says">⏳ ${escape(s.says)}</p>${state === 'done' && s.why ? `<p class="why">${escape(s.why)}</p>` : ''}${action}</div>`
}

function renderGrid(): string {
  const { scenario, lanes } = play
  const columns = `grid-template-columns: 36px repeat(${lanes.length}, minmax(230px, 1fr))`
  const head = '<div class="cell corner"></div>' + lanes.map(renderHeader).join('')
  const rows = scenario.steps.map((s, index) => {
    const cls = index < play.cursor ? 'past' : index === play.cursor ? 'now' : 'later'
    const number = `<div class="cell number ${cls}">${index + 1}</div>`
    if (!s.actor) return number + `<div class="cell ${cls}" style="grid-column: span ${lanes.length}">${renderTimeCard(s, index)}</div>`
    return number + lanes.map((lane) => `<div class="cell ${cls}">${lane.actor === s.actor ? renderCard(s, index) : ''}</div>`).join('')
  }).join('')
  return `<div class="grid" style="${columns}">${head}${rows}</div>`
}

function render() {
  const { scenario } = play
  element('summary').textContent = scenario.summary
  element('proposal').innerHTML = renderProposal()
  element('lanes').innerHTML = renderGrid()
  const total = scenario.steps.length
  const finished = play.cursor >= total
  const failed = play.done.some((d) => d.problems.length)
  ;(element('next') as HTMLButtonElement).disabled = finished || play.busy
  ;(element('play') as HTMLButtonElement).disabled = finished || play.busy
  element('verdict').textContent = finished ? (failed ? '✗ the story did not play out as written' : 'the end') : `step ${play.cursor + 1} of ${total}`
  element('verdict').className = finished && failed ? 'fail' : ''
  element('lanes').querySelector('.card.current')?.scrollIntoView({ block: 'nearest', behavior: 'smooth' })
}

function onLanesClick(event: Event) {
  if ((event.target as HTMLElement).closest('button.run')) step()
}

function onPayloadToggle(event: Event) {
  const details = event.target as HTMLDetailsElement
  if (!details.matches('details.payload')) return
  const index = Number(details.dataset.step)
  if (details.open) play.opened.add(index)
  else play.opened.delete(index)
}

function onScenarioChange(event: Event) {
  load(scenarios.find((s) => s.id === (event.target as HTMLSelectElement).value) as Scenario)
}

function restart() {
  load(play.scenario)
}

const picker = element('scenario') as HTMLSelectElement
picker.innerHTML = scenarios.map((s) => `<option value="${s.id}">${s.id}. ${escape(s.title)}</option>`).join('')
picker.addEventListener('change', onScenarioChange)
element('next').addEventListener('click', step)
element('play').addEventListener('click', playAll)
element('restart').addEventListener('click', restart)
element('lanes').addEventListener('click', onLanesClick)
element('lanes').addEventListener('toggle', onPayloadToggle, true)
load(play.scenario)
