# Token Bucket Rate Limiter Service — Full Build Guide
*For: 3rd-year ECE student, NIT Jaipur — internship interviews, 15–18 day timeline*

---

## 0. TL;DR

- **Stack: Go + Gin + Redis + Docker.** Not C++. Reasons below — read it before you argue with me.
- Day 6 checkpoint: token bucket algorithm working, persisted in Redis, concurrency-safe, one endpoint, resume-postable.
- Day 20 checkpoint: sliding window mode, admin API, rate-limit headers, load test report, README with a diagram, optionally a dashboard.
- This is a genuinely good project. It is not a CRUD app. Treat "why is this hard" as your interview script.

---

## PART 0 — TECH STACK DECISION

### Comparison

| Criterion | **Option 1: C++ (Drogon/Crow) + Redis + Docker** | **Option 2: Java + Spring Boot + Redis + Docker** | **Option 3: Go + Gin + Redis + Docker** |
|---|---|---|---|
| Learning curve (given you know DSA, not backend) | Steep. You'll fight the framework, JSON parsing, Redis client bindings, and build tooling (CMake) all at once. | Moderate. Spring Boot hides a lot behind annotations — fast to get running, slower to *understand what's happening*. | Shallow. Gin is ~an hour to learn. Go's stdlib (`net/http`, `sync`, `context`) is small and readable. |
| Resume value | High *if finished well*. Rare among interns. Signals systems-level comfort. | Medium-high. Extremely common stack; less differentiating. | High. Go is the de facto language for infra/backend tooling (rate limiters, proxies, schedulers are *literally what Go was built for*). |
| Development speed (for a backend beginner) | Slowest. Expect to burn 3–4 days just on framework + JSON + Redis client setup before writing any rate-limit logic. | Medium. Spring Boot's auto-config gets you a working server fast, but debugging "why didn't my bean load" costs time you don't have. | Fastest. Minimal boilerplate, one binary, one `go.mod`, concurrency primitives are built into the language you'll be using anyway. |
| Interview value (concurrency + systems questions) | Very high — pointers, manual memory, mutexes are impressive if you can defend every line. | Medium — JVM memory model questions are less commonly asked of interns. | High — goroutines/channels map *directly* onto "how did you handle concurrent requests," which is the centerpiece question for this project. |
| Beginner difficulty | Hard. C++ error messages, template soup in JSON libraries, and CMake/build config will eat your 18 days. | Medium. Annotations are magic until they aren't; when Spring breaks, beginners lose hours. | Easy. `go build`, single binary, minimal ceremony, very good docs. |
| Risk of not finishing | **High** | Medium | **Low** |

### Verdict: **Go + Gin + Redis + Docker**

You said you prefer C++ — I'm overriding that, and here's the actual engineering reasoning, not just "Go is trendy":

1. **You have 15–18 days and you're new to backend.** The bottleneck in this project isn't the algorithm (token bucket is ~20 lines of logic you already understand from DSA). The bottleneck is *backend plumbing*: HTTP routing, JSON (de)serialization, a Redis client, concurrency primitives, Docker. In C++, every one of those is a separate fight. In Go, all of them are either stdlib or a single well-documented import.
2. **Go's concurrency model *is* the interview story.** This project's most interesting technical question is "how did you prevent two concurrent requests from double-spending the same token bucket?" Go gives you goroutines, channels, `sync.Mutex`, and `context` as first-class citizens — you'll naturally produce idiomatic, defensible code. In C++ you'd spend real time on `std::mutex` boilerplate and thread pool setup that has nothing to do with rate limiting itself.
3. **Go is what real rate limiters are written in.** Reverse proxies, API gateways, and rate limiters (Envoy uses C++, but most rate-limit *services* like `ratelimit` by Lyft — used at real companies — are written in Go). You can honestly say "I used the language this class of infra software is typically built in."
4. **Redis integration is a one-liner in Go** (`go-redis`), vs. non-trivial client setup in C++.
5. **You still get a legitimate systems-flavored project.** Go isn't "easy mode that looks bad on a resume" — it's the standard choice for this exact problem domain at companies like Uber, Cloudflare (in part), and Lyft.

**If you still want C++ despite the risk:** Use **Drogon** (has an ORM, JSON, async, and a Redis client built in — Crow/Pistache will cost you extra glue code). Expect to lose ~3 days of your 18 to framework friction, which means cutting the sliding-window mode or the dashboard. I will not recommend this path, but if you go this route anyway, tell me and I'll re-cut the roadmap around Drogon.

**Skip Java/Spring Boot** for this project specifically — it's a fine stack, but it's the *least* differentiating of the three for an intern resume, and Spring's auto-configuration magic tends to cost beginners debugging time disproportionate to the learning value.

---

## PART 1 — PROJECT ARCHITECTURE

### High-level architecture

```
                         ┌─────────────────────────────┐
                         │        Client / Caller        │
                         └───────────────┬───────────────┘
                                          │ HTTP
                                          ▼
                    ┌─────────────────────────────────────┐
                    │           Go + Gin HTTP Server        │
                    │  ┌─────────────┐   ┌───────────────┐ │
                    │  │  /check     │   │  /admin/*     │ │
                    │  │  (rate      │   │  (config CRUD)│ │
                    │  │  limit API) │   │               │ │
                    │  └──────┬──────┘   └───────┬───────┘ │
                    │         │                  │         │
                    │  ┌──────▼──────────────────▼──────┐  │
                    │  │     Rate Limiter Engine         │  │
                    │  │  - TokenBucket strategy          │  │
                    │  │  - SlidingWindow strategy         │  │
                    │  │  (Strategy pattern, interface)   │  │
                    │  └──────────────┬───────────────────┘  │
                    └─────────────────┼──────────────────────┘
                                       │ Lua script (atomic)
                                       ▼
                    ┌─────────────────────────────────────┐
                    │              Redis                    │
                    │  client:{id}:bucket   -> tokens,ts    │
                    │  client:{id}:config   -> rps,burst    │
                    │  client:{id}:window   -> sorted set    │
                    └─────────────────────────────────────┘
```

### Components

1. **HTTP layer (Gin):** routes, request validation, response headers.
2. **Rate limiter engine:** pure logic, no HTTP awareness — takes `(clientID, algorithm)` → `(allowed bool, remaining int, resetAt time)`. This separation is what lets you support two algorithms cleanly and what an interviewer will specifically probe ("why is your engine decoupled from your handler?").
3. **Redis:** source of truth for both bucket state and per-client config. Chosen because it gives you atomicity primitives (Lua scripting, `MULTI/EXEC`) and TTLs for free — exactly what a hand-rolled in-memory store would force you to reinvent.
4. **Admin API:** plain REST CRUD over client configs, stored in Redis (or Postgres if you want a relational option — not necessary, adds scope).
5. **Load tester:** a small Go program (or `k6`/`hey`/`vegeta` — see Part 2) that fires N req/sec and asserts allow/deny counts match theory.

### Request lifecycle (`/check?client_id=X`)

```
1. Gin receives request → extracts client_id
2. Engine loads that client's config from Redis (limit, burst, algorithm)
   - if none, use default config
3. Engine executes an atomic Lua script in Redis:
     a. fetch current token count + last refill timestamp
     b. compute tokens to add = (now - last_refill) * rate
     c. new_tokens = min(burst, current + tokens_to_add)
     d. if new_tokens >= 1: new_tokens -= 1, ALLOW
        else: DENY
     e. write back new_tokens + timestamp
4. Handler sets response headers:
     X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset
5. Return 200 {"decision":"ALLOW"} or 429 {"decision":"DENY"}
```

**Why the Lua script matters (this is your concurrency answer):** Redis executes a Lua script as a single atomic operation — no other command can interleave mid-script, even under concurrent load. That's what prevents two simultaneous requests from both reading "3 tokens left" and both being allowed (double-spend). This is *the* central design decision of the whole project — know it cold.

### Redis schema

```
Key: rl:cfg:{client_id}          Type: Hash
  fields: rate (float, tokens/sec), burst (int), algo ("token_bucket"|"sliding_window")

Key: rl:bucket:{client_id}       Type: Hash   (token bucket state)
  fields: tokens (float), last_refill (unix_ms)

Key: rl:window:{client_id}       Type: Sorted Set (sliding window state)
  members: request timestamps (score = timestamp), used with ZREMRANGEBYSCORE to expire old entries
```

### Admin API design

```
POST   /admin/clients                body: {client_id, rate, burst, algo}   → create config
GET    /admin/clients/{id}                                                    → read config
PUT    /admin/clients/{id}            body: {rate, burst, algo}              → update config
DELETE /admin/clients/{id}                                                    → delete config
GET    /admin/clients                                                        → list all
```

### Rate limiting API design

```
GET /check?client_id=abc
  200 {"decision":"ALLOW","remaining":4}   headers: X-RateLimit-*
  429 {"decision":"DENY","retry_after":2}  headers: X-RateLimit-*
```

### Persistence strategy

Redis with `appendonly yes` (AOF persistence) — bucket state and configs survive a Redis restart. This is a config flag, not code you write, but you must be able to explain the tradeoff (AOF = durability + slightly slower writes vs. RDB snapshotting = faster but can lose the last few seconds).

### Concurrency strategy

- **Within Redis:** atomicity via Lua scripting (see above) — this is your primary safety mechanism, not app-level locking.
- **Within Go:** Gin handles each request on its own goroutine automatically; you don't need to manage threads yourself. Where you *do* need care: any shared in-process state (e.g., a metrics counter for the dashboard) needs a `sync.Mutex` or `sync/atomic`.
- Explicitly avoid: "check-then-act" logic split across two separate Redis calls (e.g., `GET` tokens, then `SET` tokens in a second round trip) — that's a classic race condition and a common mistake (see Part 4).

### Error handling strategy

- Redis unreachable → **fail open or fail closed** is a decision you must make and defend. For an internship project, fail-closed (deny) is the safer/more defensible answer ("availability of the protected service matters more than a few false denies"), but document that you *considered* fail-open and why you chose otherwise.
- Invalid client_id / missing config → fall back to a default rate limit, don't 500.
- Malformed admin payloads → 400 with a clear validation message.

---

## PART 2 — CONCEPTS YOU MUST LEARN

Only what's actually required. Ordered by when you'll need it.

| # | Concept | What it is | Why needed | Difficulty | Est. time | Learn for THIS project |
|---|---|---|---|---|---|---|
| 1 | HTTP fundamentals | Request/response, methods, status codes, headers | Every API decision depends on this | Easy | 2–3 hrs | Just enough to know GET vs POST, 200/400/429/500, and what a header is |
| 2 | REST API design | Resource-oriented URL/verb conventions | Your admin API needs to look "correct" to an interviewer | Easy | 2 hrs | CRUD verb mapping, status code conventions |
| 3 | Go basics | Syntax, structs, interfaces, error handling (`if err != nil`) | You're building the whole thing in it | Easy–Medium | 2 days | Enough to write handlers, structs, interfaces — skip generics/advanced tooling |
| 4 | Goroutines & channels | Go's concurrency primitives | Concurrent request handling, and to *talk about* concurrency in interviews | Medium | 1 day | `sync.Mutex`, `sync/atomic`, basic goroutine model — skip channels unless you use them for the dashboard |
| 5 | Gin framework | Go HTTP router/framework | Your web layer | Easy | 3–4 hrs | Routing, middleware, JSON binding — that's it |
| 6 | Redis basics | In-memory data store w/ persistence | Your state store | Medium | 1 day | Strings, hashes, sorted sets, TTL, `MULTI/EXEC` |
| 7 | Redis Lua scripting (`EVAL`) | Atomic multi-step ops inside Redis | **This is your concurrency-safety mechanism** — non-negotiable | Medium | 1 day | Writing one Lua script for token bucket, one for sliding window |
| 8 | Rate limiting algorithms | Token bucket & sliding window theory | The actual core logic | Medium | 1 day | Token bucket refill math, sliding window log/counter tradeoffs |
| 9 | Race conditions & atomicity | Concurrent access bugs | You must be able to explain *why* your design avoids them | Medium | half day (mostly conceptual, paired with #7) | Check-then-act problem, atomic ops, idempotency |
| 10 | Docker basics | Containerization | Packaging + demo-ability | Easy | 4–5 hrs | `Dockerfile`, `docker-compose.yml` for app+Redis — skip Kubernetes entirely |
| 11 | Load testing | Simulating concurrent traffic | Your 500+ req/sec proof | Easy | half day | Learn one tool (`vegeta` or `hey`), not several |
| 12 | Deployment (basic) | Getting it publicly reachable | Demo strategy for recruiters | Easy | half day | Render/Railway/Fly.io free tier — skip AWS/GCP infra setup, not worth the time here |

**Explicitly skip:** Kubernetes, gRPC, message queues (Kafka/RabbitMQ), microservice orchestration, service mesh, OAuth/full auth systems, ORMs, GraphQL. None of these serve this project's story and all of them will eat days you don't have.

**Best resources** (pick one source per concept, don't rabbit-hole comparing five):
- Go: *Tour of Go* (go.dev/tour) + *Go by Example* (gobyexample.com)
- Gin: official Gin README + a single "Gin CRUD API" YouTube walkthrough (search "Gin framework tutorial")
- Redis: *Redis University RU101* (free) or the official Redis docs "Data types" page
- Redis Lua/EVAL: official Redis docs page on "Scripting with Lua" — short, sufficient
- Token bucket / sliding window: Cloudflare's engineering blog post "How we built rate limiting capable of scaling to millions of domains" (conceptual, well-written) + Redis's own "Rate limiting patterns" doc
- Docker: *Docker for Beginners* (docker-curriculum.com)
- Load testing: `vegeta` README (github.com/tsenart/vegeta) — 10-minute read, does everything you need

---

## PART 3 — DAY-BY-DAY ROADMAP (18 days)

Assume ~3–4 focused hours/day on weekdays, more on weekends. Adjust to your actual schedule, but keep the Day-6 and Day-18/20 checkpoints fixed.

### Days 1–3: Foundation

**Day 1** — *Objective: environment + Go basics*
- Learn: Go syntax, structs, interfaces, error handling
- Build: nothing yet — write 2–3 throwaway Go scripts
- Deliverable: comfortable writing a basic Go program
- Hours: 3–4

**Day 2** — *Objective: HTTP + Gin*
- Learn: HTTP fundamentals, Gin routing/middleware/JSON binding
- Build: a "hello world" Gin server with 2 dummy routes
- Deliverable: server that responds to `GET /ping`
- Hours: 3

**Day 3** — *Objective: Redis basics*
- Learn: Redis data types, TTL, running Redis via Docker
- Build: connect Go to Redis (`go-redis`), write/read a test key
- Deliverable: Go app that can SET/GET from a Redis container
- Hours: 3–4

### Days 4–6: MVP (resume-postable by Day 6)

**Day 4** — *Objective: token bucket logic*
- Learn: token bucket refill math, atomicity/race conditions
- Build: token bucket algorithm as a plain Go function (in-memory first, no Redis yet) — prove correctness with a quick manual test
- Deliverable: working token bucket function, unit-tested by hand
- Hours: 4

**Day 5** — *Objective: move state into Redis atomically*
- Learn: Redis Lua scripting (`EVAL`)
- Build: rewrite token bucket logic as a Lua script executed via `EVAL`; wire it into your `/check` endpoint
- Deliverable: `GET /check?client_id=X` returns ALLOW/DENY backed by Redis, race-condition-safe
- Hours: 4–5

**Day 6** — *Objective: MVP polish + headers*
- Build: `X-RateLimit-Limit/Remaining/Reset` headers, default config fallback, basic error handling, `docker-compose.yml` (app + Redis)
- Deliverable: **Resume checkpoint** — a working, persistent, concurrency-safe, containerized rate limiter with one algorithm and proper headers. Push to GitHub with a README covering what's done.
- Hours: 4

> **By end of Day 6 you can honestly write:** *"Built a token-bucket rate limiter service in Go, backed by Redis with atomic Lua scripting for concurrency safety, containerized with Docker."* That's a real, defensible resume line.

### Days 7–12: Advanced functionality

**Day 7** — *Objective: admin API*
- Build: `POST/GET/PUT/DELETE /admin/clients` — full config CRUD, stored in Redis hashes
- Deliverable: per-client configurable rate/burst working end-to-end
- Hours: 4

**Day 8** — *Objective: sliding window theory*
- Learn: sliding window log vs. sliding window counter tradeoffs
- Build: design (on paper) the sorted-set-based sliding window approach
- Deliverable: written design note (this doubles as interview prep)
- Hours: 3

**Day 9–10** — *Objective: implement sliding window*
- Build: sliding window Lua script using Redis sorted sets (`ZADD`, `ZREMRANGEBYSCORE`, `ZCARD`), selectable per client via the `algo` config field
- Deliverable: `/check` supports both algorithms, switchable per client
- Hours: 4–5/day

**Day 11** — *Objective: strategy pattern cleanup*
- Build: refactor engine behind a common Go interface (`Limiter.Allow(clientID) (bool, Meta)`) implemented by both algorithms — makes the code interview-clean
- Deliverable: clean, swappable algorithm architecture
- Hours: 3

**Day 12** — *Objective: robustness*
- Build: fail-closed/fail-open handling for Redis downtime, input validation, structured error responses, basic logging
- Deliverable: service that degrades predictably under failure
- Hours: 3–4

### Days 13–15: Load testing & deployment

**Day 13** — *Objective: load testing setup*
- Learn: `vegeta` (or `hey`)
- Build: load test scripts hitting `/check` at 500+ req/sec with a mix of clients
- Deliverable: raw load test output (latency percentiles, allow/deny counts)
- Hours: 3–4

**Day 14** — *Objective: correctness verification*
- Build: a small verification script comparing expected vs. actual allow/deny counts under load (prove your bucket math holds under concurrency — this is your "proof" artifact)
- Deliverable: a report/markdown doc with graphs or tables showing correctness under load
- Hours: 4

**Day 15** — *Objective: deployment*
- Build: deploy to Render/Railway/Fly.io free tier (app + managed/containerized Redis)
- Deliverable: a live, public URL you can hand to a recruiter
- Hours: 3

### Days 16–18: Polish, docs, dashboard, interview prep

**Day 16** — *Objective: dashboard (stretch goal)*
- Build: a minimal HTML/JS page (or Go template) hitting a `/metrics` endpoint showing live allow/deny counts per client
- Deliverable: dashboard, or skip if behind schedule (see Part 5)
- Hours: 4

**Day 17** — *Objective: documentation*
- Build: full README (architecture diagram, setup instructions, API docs, load test results, design decisions + tradeoffs)
- Deliverable: polished GitHub repo
- Hours: 3–4

**Day 18** — *Objective: interview prep*
- Build: nothing — review Part 6 of this doc, rehearse explaining every design decision out loud
- Deliverable: you, confidently explaining the Lua atomicity trick without notes
- Hours: 2–3

---

## PART 4 — IMPLEMENTATION ORDER (with reasoning)

1. **Basic HTTP server (Gin, no logic)** — Dependencies: none. Why first: validates your dev environment before anything else. Common mistake: skipping this and debugging Gin + Redis + algorithm bugs simultaneously. Difficulty: trivial.
2. **Redis connectivity** — Dependencies: Docker running Redis. Why here: nothing downstream works without it. Common mistake: not setting a connection timeout, so a dead Redis hangs your whole server. Difficulty: easy.
3. **Token bucket logic, in-memory first** — Dependencies: none (pure function). Why: validate the *algorithm* is correct before adding the complexity of Redis/atomicity. Common mistake: trying to write the Lua script and the algorithm logic simultaneously — decouple them. Difficulty: medium (the refill-rate math trips people up — work through it on paper with real numbers first).
4. **Move token bucket into Redis via Lua** — Dependencies: #2, #3. Why here, not earlier: you need a *correct* algorithm before you worry about making it atomic. Common mistake: doing a `GET` then a separate `SET` from Go instead of one atomic Lua script — this reintroduces the exact race condition the project is supposed to solve. Difficulty: medium-high (Lua syntax is unfamiliar, but the script itself is short).
5. **Rate-limit headers** — Dependencies: #4. Why here: trivial once the engine returns remaining/reset values. Common mistake: computing `Reset` incorrectly (it should reflect when the bucket will next have a token, not a fixed window). Difficulty: easy.
6. **Admin API (config CRUD)** — Dependencies: #2. Why after MVP, not before: config CRUD is easy but not differentiating — get the harder algorithmic core working and provable first, so if you run short on time later, this doesn't block your resume checkpoint. Difficulty: easy.
7. **Sliding window algorithm** — Dependencies: #4 (reuse the Lua/atomicity pattern), #6 (needs per-client `algo` selection). Common mistake: implementing sliding window as a simple counter reset every N seconds — that's *fixed* window, not sliding, and interviewers will immediately catch this distinction. Difficulty: medium.
8. **Interface/strategy refactor** — Dependencies: #7. Why after both algorithms exist: premature abstraction before you have two concrete implementations to generalize from is wasted effort. Difficulty: easy-medium.
9. **Failure handling (fail open/closed, validation)** — Dependencies: everything above works on the happy path first. Difficulty: easy.
10. **Load testing** — Dependencies: a fully working service. Common mistake: load testing too early and mistaking "my laptop is slow" bugs for "my algorithm is wrong" bugs. Difficulty: easy (tooling), medium (interpreting results correctly).
11. **Deployment** — Dependencies: Docker Compose working locally first. Difficulty: easy-medium (mostly platform-specific fiddling).
12. **Dashboard, docs, polish** — Last, because they don't block correctness and are the first things to cut if you're behind.

---

## PART 5 — MVP vs. ADVANCED vs. STRETCH vs. OVERKILL

Brutally practical, as requested.

### 1. Must-have (resume-worthy — do not skip these)
- Token bucket algorithm, correct refill math
- Redis-backed persistence (survives restart)
- Atomic concurrency handling via Lua scripting (this is the single most important thing you build — it's the whole "why this isn't a CRUD app" answer)
- `/check` endpoint returning ALLOW/DENY
- Rate-limit response headers
- Per-client configurable limits
- Docker Compose (app + Redis)
- A load test proving correctness at 500+ req/sec
- A clear README

### 2. Good-to-have (do if on schedule)
- Sliding window mode, selectable per client
- Full admin CRUD API
- Fail-open/fail-closed handling with justification
- Deployed live demo (Render/Railway/Fly.io)
- Basic structured logging

### 3. Stretch goals (only after everything above is solid)
- Live metrics dashboard
- Distributed mode (multiple app instances sharing Redis — note: if Redis is already your shared store, you get *most* of this "for free" once your app is stateless; true distributed-mode complexity is really about multiple *Redis* nodes, which is overkill — see below)

### 4. Overkill for an internship project — do NOT do these
- True distributed Redis (Cluster mode, Sentinel, replication) — massively out of scope for 18 days and not expected of an intern
- Kubernetes deployment
- Custom auth/JWT system for the admin API (an API key check is enough)
- gRPC instead of REST
- Multiple rate-limiting algorithms beyond the two requested (e.g., leaky bucket, fixed window as a *third* mode) — diminishing resume return per hour spent
- A frontend framework (React etc.) for the dashboard — plain HTML+JS or a Go template is enough and saves you a day

### If you're behind schedule, cut in this exact order:
1. Dashboard (cut first — lowest resume ROI per hour)
2. Deployment (keep it running locally + Docker Compose is still legitimate to demo)
3. Sliding window mode (keep token bucket only — still a complete, defensible project)
4. Full admin CRUD (keep just config *creation*, drop update/delete)

**Never cut:** the Lua-script atomicity, the load test, or the README. Those three are what separate this from a toy project.

---

## PART 6 — INTERVIEW PREPARATION

For each topic: ideal answer, likely follow-ups, and what to have ready for a deep dive.

### Token Bucket
**Q: Explain how your token bucket algorithm works.**
A bucket holds up to `burst` tokens. Tokens refill continuously at `rate` tokens/sec. Each request consumes one token if available (ALLOW) or is rejected if the bucket is empty (DENY). I store `tokens` and `last_refill_timestamp` per client in Redis; on each request I compute `elapsed = now - last_refill`, add `elapsed * rate` tokens (capped at `burst`), then decide allow/deny.
- *Follow-up:* Why token bucket over a simple fixed counter? → Token bucket allows controlled bursts (up to `burst` size) while still enforcing a long-run average rate; a fixed counter resets abruptly at window boundaries and allows a burst of 2x the limit right at the boundary edge.
- *Deep dive:* What happens with clock skew across instances? → I rely on Redis's server time (via the Lua script using `redis.call('TIME')`) rather than each app instance's local clock, precisely to avoid skew between app servers.

### Sliding Window
**Q: Why is sliding window different from fixed window, and how did you implement it?**
Fixed window resets a counter every N seconds — this allows a burst of up to 2x the limit around the boundary. Sliding window avoids that by tracking actual request timestamps (I used a Redis sorted set, score = timestamp) and counting only requests within the trailing N-second window, trimming old entries with `ZREMRANGEBYSCORE`.
- *Follow-up:* What's the memory cost vs. token bucket? → Sliding window log stores one entry per request in the window, vs. token bucket's O(1) state — a real tradeoff I'd mention explicitly.
- *Deep dive:* Sliding window counter (approximate, weighted average of two fixed windows) vs. sliding window log (exact, what I built) — be ready to explain both exist and why you picked the exact version for correctness given the project's explicit requirement.

### Redis
**Q: Why Redis instead of an in-memory Go map or a SQL database?**
In-memory map: fast but doesn't survive restarts and doesn't work across multiple app instances. SQL: durable but too slow for a per-request hot path and doesn't offer atomic scripting the way Redis does. Redis gives sub-millisecond latency, built-in TTL, and — critically — atomic Lua execution, which solves my concurrency requirement directly.
- *Follow-up:* What if Redis goes down? → covered under error handling — fail-closed by default, explain the tradeoff.

### Concurrency / Race Conditions / Thread Safety
**Q: How do you prevent two concurrent requests from double-spending tokens?**
The classic bug is "check-then-act" split across two network round trips: `GET` tokens, decide in application code, then `SET` tokens back — between the GET and SET, another request can read the same stale value. I eliminated this by doing the entire read-decide-write sequence inside a single Redis Lua script, which Redis executes atomically (single-threaded execution model), so no interleaving is possible.
- *Follow-up:* Could you have used a distributed lock instead? → Yes (e.g., Redlock), but that adds latency (lock acquire/release round trips) and complexity for no benefit here, since Lua scripting already gives atomicity natively.
- *Deep dive:* What about concurrency *inside* your Go process (not Redis)? → Gin handles each request on its own goroutine; any shared in-process state (like dashboard metrics) is protected with `sync/atomic` counters, not mutex-heavy code, to keep the hot path fast.

### Persistence
**Q: How does state survive a restart?**
Redis persistence — I use AOF (append-only file), which logs every write operation and replays it on restart, favoring durability over the faster-but-lossier RDB snapshot approach, since losing a few seconds of rate-limit state could let clients briefly exceed limits.

### API Design
**Q: Walk me through your API design choices.**
REST-ish resource model for admin (`/admin/clients/{id}` with standard verbs), a single purpose-built `/check` endpoint for the hot path (not RESTful-CRUD, since it's an action, not a resource), and standard rate-limit headers so this could be a drop-in reference for the `X-RateLimit-*` convention used by GitHub/Twitter/Stripe APIs.

### Load Testing
**Q: How did you validate correctness under load, not just throughput?**
I didn't just measure requests/sec — I asserted *correctness*: for a client configured with rate=10/sec, burst=20, fired 500 req/sec for 10 seconds, and verified the allow count matched the theoretical maximum (burst + rate×time) within a small margin, proving no double-spending occurred under real concurrency.
- *Follow-up:* What tool did you use and why? → name your tool (`vegeta`/`hey`), and that you chose it for simplicity over something like JMeter/Locust, which were unnecessary for this scope.

### Scaling
**Q: How would this scale to multiple instances / much higher traffic?**
Because all state lives in Redis (not in-process), the Go app itself is stateless — I can run N instances behind a load balancer with no code changes, since they all hit the same Redis. The actual bottleneck at very high scale becomes Redis itself, which is where you'd introduce Redis Cluster/sharding by client_id — I'd call this out as a known next step rather than something I implemented, given the 18-day scope.

---

## PART 7 — RESUME POSITIONING

**Project title:** *Distributed Rate Limiter Service (Token Bucket & Sliding Window, Go + Redis)*

**Resume bullets (pick 2–3, don't cram all of them):**
- Designed and built a rate-limiting microservice in Go supporting Token Bucket and Sliding Window algorithms, using atomic Redis Lua scripting to guarantee correctness under concurrent access.
- Engineered per-client configurable rate limits via an admin REST API, with state persisted in Redis (AOF) to survive service restarts.
- Load-tested the service at 500+ concurrent requests/sec using Vegeta, verifying zero double-spend violations and sub-X-ms p99 latency.
- Containerized the full stack with Docker Compose and deployed a public demo on [Render/Railway/Fly.io].

**Technologies to list:** Go, Gin, Redis (Lua scripting/EVAL, AOF persistence), Docker, Docker Compose, REST API design, concurrency (goroutines, `sync/atomic`), Vegeta (load testing).

**Quantifiable achievements to mention (fill in with your real numbers):**
- "500+ req/sec sustained, p99 latency under X ms"
- "Zero rate-limit violations across N load-test runs"
- "Supports M concurrent client configurations"

**How to present it to recruiters:** Lead with the *problem*, not the tech list — "Most rate limiters are a library import; I wanted to understand what happens when you build the shared, concurrent, persistent version yourself." Then be ready to draw the Lua-atomicity diagram on a whiteboard from memory — that single explanation carries most of the interview value of this project.

**"What makes this significantly better than a CRUD app?"**
A CRUD app tests whether you can wire a database to a form. This project tests whether you understand *concurrent, shared, stateful systems* — the same category of problem behind real infrastructure (API gateways, distributed locks, quota systems). It has a genuine correctness property (no double-spending under concurrency) that you can *prove* with a load test, not just "it looks like it works." That's a fundamentally different, harder claim to defend — and defending it well is exactly what separates an intern candidate who understands backend systems from one who's only shipped tutorials.

---

## PART 8 — GITHUB & DEPLOYMENT

### Folder structure
```
rate-limiter/
├── cmd/
│   └── server/main.go
├── internal/
│   ├── limiter/          # algorithm interface + token bucket + sliding window
│   ├── handler/          # Gin route handlers
│   ├── config/           # admin config CRUD logic
│   └── redisclient/      # Redis connection + Lua script loading
├── scripts/
│   ├── token_bucket.lua
│   └── sliding_window.lua
├── loadtest/
│   └── vegeta-targets.txt, results.md
├── docker-compose.yml
├── Dockerfile
├── go.mod / go.sum
└── README.md
```

### Git commit milestones (make these real commits, not one giant dump)
1. `chore: project scaffold, Gin hello-world`
2. `feat: redis connectivity`
3. `feat: token bucket algorithm (in-memory)`
4. `feat: token bucket via atomic Lua script + Redis persistence`
5. `feat: rate-limit response headers`
6. `feat: admin CRUD API for client configs`
7. `feat: sliding window algorithm`
8. `refactor: extract limiter strategy interface`
9. `feat: fail-open/closed handling + validation`
10. `test: load testing suite + correctness report`
11. `chore: docker-compose, deployment config`
12. `feat: metrics dashboard` (if time allows)
13. `docs: full README + architecture diagram`

A clean commit history like this is itself something interviewers glance at — it shows incremental, deliberate engineering rather than one dump.

### README structure
1. One-paragraph problem statement (why this isn't a CRUD app)
2. Architecture diagram (reuse the ASCII one above, or redraw cleanly)
3. Quickstart (`docker-compose up`)
4. API reference (`/check`, `/admin/*`)
5. Design decisions & tradeoffs (algorithm choice, atomicity approach, fail-open/closed)
6. Load test results (numbers + how to reproduce)
7. What I'd do differently at larger scale (Redis Cluster, etc.) — shows maturity

### Deployment options
- **Render or Railway free tier:** easiest for a Go app + Redis add-on, minimal config, good enough for a live demo link.
- **Fly.io:** slightly more control, still free-tier friendly, good if you want practice with a `fly.toml`.
- Skip AWS/GCP/Azure for this — the IAM/VPC/security-group setup alone could eat 2 of your remaining days for zero resume differentiation over Render/Railway.

### Demo strategy for recruiters/interviewers
- Have the live URL in your README and resume.
- Prepare a 60-second live demo: hit `/check` a few times past the limit in a terminal (`curl` in a loop) and show the 429 kick in — visceral and fast, better than describing it.
- Have the load-test report open as a backup in case the live demo network hiccups during an interview call.

---

## Final note

The single highest-leverage thing in this entire project is being able to explain, unprompted and precisely, **why a Lua script inside Redis prevents the race condition that a naive GET-then-SET implementation would have.** Everything else — sliding window, dashboard, deployment — is set dressing around that one idea. If you only have time to nail one explanation before an interview, nail that one.
