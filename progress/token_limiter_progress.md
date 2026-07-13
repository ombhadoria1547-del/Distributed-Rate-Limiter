# 🚀 Distributed Rate Limiter Project Tracker (v2)

----

# 📌 PROJECT CHARTER

## Project Title

**Distributed Rate Limiter Service**
(Token Bucket + Sliding Window)

---

## Project Goal

Build a production-style backend systems project capable of enforcing rate limits across multiple clients while maintaining correctness under concurrent access.

The project should demonstrate:

* Backend Development
* REST API Design
* Redis Integration
* Concurrency Handling
* Distributed State Management
* Rate Limiting Algorithms
* Docker Containerization
* Load Testing
* Deployment
* Technical Documentation

---

## Why This Project Exists

Most student projects are CRUD applications.

This project focuses on:

* Shared State
* Concurrency
* Atomic Operations
* System Design
* Infrastructure Concepts

The objective is not simply to implement a rate limiter.

The objective is to understand how production systems prevent abuse, manage shared resources, and maintain correctness under load.

---

## Why Go + Gin + Redis + Docker

### Go

Chosen because:

* Simple backend development
* Excellent concurrency model
* Minimal boilerplate
* Widely used in infrastructure software

### Gin

Chosen because:

* Lightweight
* Easy routing
* Fast setup
* Industry adoption

### Redis

Chosen because:

* Extremely fast
* Persistence support
* Atomic operations
* Lua scripting support

### Docker

Chosen because:

* Consistent environments
* Easy deployment
* Industry standard

---

## Resume Target

By completion, this project should be strong enough to appear as a major backend project on internship resumes.

Expected Resume Bullet:

> Built a distributed rate limiter service in Go using Token Bucket and Sliding Window algorithms, Redis-backed persistence, atomic Lua scripting, Docker containerization, and load testing under concurrent traffic.

---

## MVP Checkpoint (Day 6)

By Day 6 the project should support:

* Token Bucket
* Redis Persistence
* Atomic Lua Script
* /check Endpoint
* Docker Compose
* Rate Limit Headers

At this point the project becomes resume-worthy.

---

## Final Checkpoint (Day 18)

Project should include:

* Token Bucket
* Sliding Window
* Admin API
* Load Testing
* Deployment
* Documentation
* Interview Preparation

---

# 🏗️ SYSTEM ARCHITECTURE

## High Level Flow

Client

↓

HTTP Request

↓

Gin Server

↓

Rate Limiter Engine

↓

Redis

↓

ALLOW / DENY Response

---

## Core Components

### HTTP Layer

Responsibilities:

* Routes
* Validation
* Headers
* JSON Responses

Endpoints:

```text
/check
/admin/*
```

---

### Rate Limiter Engine

Responsibilities:

* Token Bucket Logic
* Sliding Window Logic
* Strategy Selection

Must remain independent from HTTP layer.

---

### Redis

Responsibilities:

* Bucket State
* Client Configuration
* Sliding Window Data
* Persistence

Redis acts as the source of truth.

---

### Lua Scripts

Responsibilities:

* Atomic Updates
* Race Condition Prevention
* Concurrency Safety

Most important technical concept in the entire project.

---

### Load Testing Layer

Responsibilities:

* Performance Validation
* Correctness Validation
* Stress Testing

---

# 🗺️ MASTER ROADMAP

---

## Phase 1 — Foundations

### Git & GitHub

Objectives:

* Version Control
* Branching
* Merging
* Rebase
* Commit History

Deliverable:

Professional Git Workflow

Status:

✅ Completed

---

### Go Fundamentals

Objectives:

* Variables
* Functions
* Arrays
* Slices
* Maps
* Structs
* Pointers
* Methods
* Interfaces
* Error Handling
* Packages

Deliverable:

Comfortable writing Go programs independently.

Status:

✅ Completed

---

### HTTP Fundamentals + Gin

Objectives:

* Request / Response Cycle
* HTTP Methods
* Status Codes
* Headers
* REST Basics
* Gin Routing
* JSON Responses

Deliverable:

```text
GET /ping
```

Status:

✅ Completed

---

### Redis Fundamentals

Objectives:

* Redis Data Types
* Redis Persistence
* Redis Containers
* Go ↔ Redis Connectivity

Deliverable:

Go application successfully reading and writing Redis data.

Status:

✅ Completed

---

# ⚙️ Phase 2 — MVP Development

Goal:

Build first resume-worthy version.

---

### Token Bucket

Objectives:

* Refill Logic
* Consumption Logic
* Burst Handling

Deliverable:

In-Memory Token Bucket

Status:

✅ Completed

---

### Redis Lua Scripting

Objectives:

* Atomic Operations
* Redis EVAL
* Race Condition Prevention

Deliverable:

Redis-backed Token Bucket

Status:

✅ Completed

---

### MVP Completion

Requirements:

* Token Bucket
* Redis Persistence
* Atomic Updates
* Docker Compose
* /check Endpoint
* Rate Limit Headers

Deliverable:

Resume-Ready MVP

Status:

✅ Completed

---

# 🚀 Phase 3 — Advanced Features

### Admin API

Endpoints:

```text
POST   /admin/clients
GET    /admin/clients
PUT    /admin/clients
DELETE /admin/clients
```

Status:

✅ Completed

---

### Sliding Window

Objectives:

* Sorted Sets
* Window Management
* Alternative Algorithm

Status:

✅ Completed

---

### Strategy Pattern

Objectives:

* Token Bucket
* Sliding Window

Through common interface.

Status:

🟡 In Progress (Next Up)

---

### Robustness

Objectives:

* Validation
* Logging
* Fail Open / Fail Closed

Status:

⬜ Pending

---

# 📊 Phase 4 — Validation

### Load Testing

Objectives:

* Throughput Validation
* Concurrency Validation
* Correctness Validation

Target:

500+ Requests / Second

Status:

⬜ Pending

---

### Correctness Verification

Objectives:

* Verify Allow Counts
* Verify Deny Counts
* Validate Token Bucket Mathematics

Status:

⬜ Pending

---

# 🌍 Phase 5 — Deployment

Objectives:

* Public URL
* Deployment Workflow
* Production Demonstration

Potential Platforms:

* Render
* Railway
* Fly.io

Status:

⬜ Pending

---

# 📚 Phase 6 — Documentation & Interview Preparation

### Documentation

Deliverables:

* README
* Architecture Diagram
* API Documentation
* Load Test Results
* Design Decisions

Status:

⬜ Pending

---

### Interview Preparation

Objectives:

Be able to explain:

* Token Bucket
* Sliding Window
* Redis
* Lua Scripts
* Concurrency
* Race Conditions
* Load Testing

Without looking at code.

Status:

⬜ Pending

---

# 📂 PROJECT WORKSPACE

Current Structure:

```text
token-limiter/
│
├── roadmap/
│   └── token-bucket-rate-limiter-roadmap.md
│
├── progress/
│   └── token_limiter_progress.md
│
├── docs/
│
├── notes/
│
├── experiments/
│
├── source/
│
└── README.md
```

Current Additions (as of Day 10):

```text
Dockerfile
docker-compose.yml
scripts/
    sliding_window.lua
source/
    redis_window.go
```

Future additions:

```text
cmd/
internal/
loadtest/
```

As project progresses, workspace structure must remain synchronized with roadmap milestones.

---

# 🔄 GIT POLICY

Every meaningful milestone must be committed.

Workflow:

```bash
git add .
git commit -m "meaningful message"
git push
```

Never leave major completed work uncommitted.

---

## Planned Milestone Commits

```text
✅ Git Learning

✅ Gin Hello World

✅ Redis Connectivity

✅ Token Bucket

✅ Lua Integration

✅ Docker Compose

✅ Admin API

✅ Sliding Window

⬜ Load Testing

⬜ Deployment

⬜ Documentation
```

---

# ✅ COMPLETED MILESTONES

## Git & GitHub Fundamentals

Status:

✅ Completed

Hours:

~3 Hours

Topics Covered:

* init
* add
* commit
* push
* pull
* fetch
* stash
* log
* branching
* merging
* rebase

Outcome:

Comfortable with professional Git workflow.

---

## Go Fundamentals

Status:

✅ Completed

Hours:

~5–6 Hours

Topics Covered:

* Variables
* Arrays
* Slices
* Maps
* Functions
* Structs
* Pointers
* Methods
* Interfaces
* Error Handling
* Packages
* Exported / Unexported Names

Outcome:

Comfortable understanding and writing Go code required for this project.


## Day 3 — Environment Setup & Repository Initialization

Date:

2026-07-04

Hours Spent:

~3 Hours

Topics Learned:

* Go Installation
* Go Environment Setup
* Go Modules
* Repository Structure Planning
* Git Commit Workflow
* GitHub Repository Management

Files Created:

* README.md
* go.mod
* main.go

Folders Created:

* roadmap/
* progress/
* docs/
* notes/
* experiments/
* source/

Problems Faced:

* Broken Go installation
* Incorrect PATH configuration
* Git Bash unable to locate Go executable
* Go command not recognized by terminal

Resolution:

* Reinstalled Go
* Verified installation using go version
* Successfully initialized Go module
* Executed first Go program
* Verified GitHub repository synchronization

Key Learnings:

* Difference between git add, git commit and git push
* Purpose of go.mod
* Structure of a professional Go project
* Repository initialization workflow
* Basic Go compilation and execution workflow

Git Commit Created:

Yes

Commit:

chore: initialize project structure

Outcome:

✅ Go installed and working

✅ Repository created and linked with GitHub

✅ Project structure initialized

✅ Go module initialized

✅ First Go program executed successfully

✅ Initial project commit pushed

Current Workspace:

```text
Distributed-Rate-Limiter/
│
├── README.md
├── go.mod
├── main.go
│
├── roadmap/
│   └── token-bucket-rate-limiter-roadmap.md
│
├── progress/
│   └── token_limiter_progress.md
│
├── docs/
│
├── notes/
│
├── experiments/
│
└── source/
```

---

## Day 4 — HTTP Fundamentals + Gin Framework

Date:

2026-07-07

Hours Spent:

~3–4 Hours

Topics Learned:

* HTTP Request / Response Flow
* HTTP Methods (GET, POST, PUT, DELETE)
* Status Codes (200, 400, 404, 429, 500)
* Headers (conceptual understanding)
* JSON structure (keys, values, objects)
* Installing and using Gin (`go get github.com/gin-gonic/gin`)
* Building a Gin server with `gin.Default()` and `router.Run(":8080")`
* Query Parameters (`c.Query("name")`)
* Path Parameters (`c.Param("id")`)
* Difference between query params and path params
* Manual testing via browser

Files Created:

* main.go (updated with Gin routes)

Endpoints Built:

```text
GET /ping
GET /hello
GET /user?name=om
GET /user/:id
```

Problems Faced:

* None blocking — completed without major issues

Key Learnings:

* `router := gin.Default()` creates the server instance
* `router.GET(...)` registers a route handler for incoming GET requests
* `c.JSON(...)` sends a JSON response back to the client
* Query parameters (`?name=om`) and path parameters (`/user/123`) are handled differently but serve a similar purpose
* This exact query-parameter mechanism (`?client_id=abc`) is what the future `/check` endpoint will rely on

Git Commit Created:

Yes

Commit:

feat: gin hello world

Outcome:

✅ Understood HTTP request/response flow

✅ Understood GET/POST and status codes

✅ Understood JSON structure

✅ Installed Gin

✅ Built a running API with 4 working endpoints

✅ Handled query parameters and path parameters

✅ Tested all endpoints successfully in browser

✅ Commit pushed: "feat: gin hello world"

Next Objective:

Redis Fundamentals — Redis Data Types, Persistence, Containers, Go ↔ Redis Connectivity

---

## Day 5 — Redis Fundamentals

Date:

2026-07-07

Hours Spent:

~3 Hours

Topics Learned:

* What Redis actually is — conceptually, a `map[string]string` living outside the application
* Why Redis exists: application memory disappears on restart, so token counts can't live only in a Go variable
* Why multiple Go server instances need a shared store instead of separate in-process maps
* Redis Data Types (conceptual only, not deep): String, Hash, Sorted Set
* String → simple key/value (e.g. `username = om`)
* Hash → grouped fields under one key (e.g. `user:1 → name, age`)
* Sorted Set → values with scores (previewed for future Sliding Window use)
* Why Redis is run in Docker instead of installed directly (environment consistency, matches real company practice)
* Docker commands conceptually: `docker pull redis`, `docker run redis`, `docker ps`
* Redis client for Go: `github.com/redis/go-redis/v9`
* `client := redis.NewClient(...)` creates the connection
* Performing a `SET` and `GET` from a Go program and printing the retrieved value
* How today's `username = om` exercise maps directly onto the future rate limiter schema (`client:{id} → tokens, last_refill`) — same mechanism, different data

Files Created:

* main.go (updated with Redis connection, SET, and GET logic)

Problems Faced:

* None blocking — completed without major issues

Key Learnings:

* Redis is an in-memory database used to store state outside the Go application
* A rate limiter needs Redis because token counts must survive requests and server restarts
* Go maps are insufficient because they vanish on restart and can't be shared across multiple server instances
* The Go Redis client (`go-redis/v9`) handles the connection; `SET`/`GET` round-trip correctly from Go to Redis and back
* Today's exercise (`username = om`) is functionally identical to the future bucket state (`tokens`, `last_refill`) — only the data changes, not the mechanism

Git Commit Created:

Yes

Commit:

feat: redis connectivity

Outcome:

✅ Redis container running locally

✅ Go application successfully connected to Redis

✅ Performed SET from Go

✅ Performed GET from Go

✅ Retrieved value printed and verified correct

✅ Understood why Redis exists and why Go maps are insufficient

✅ Commit pushed: "feat: redis connectivity"

Next Objective:

Token Bucket — Refill Logic, Consumption Logic, Burst Handling (In-Memory Token Bucket)

---

## Day 6 — Token Bucket (In-Memory Implementation)

Topics Learned:

* Why Rate Limiting exists — API abuse, resource exhaustion, denial of service, fair resource sharing, protecting backend systems
* Why simply adding more servers does not solve the abuse problem
* Why a simple request counter and Fixed Window approaches are insufficient
* Why Token Bucket was invented — smooth traffic flow, burst handling, average rate enforcement
* Core Token Bucket concepts: Bucket, Tokens, Capacity, Refill Rate, Consumption, Allow, Deny, Burst Capacity
* Mathematics behind Token Bucket: elapsed time, refill formula, why bucket size never exceeds capacity, why refill is calculated lazily instead of every second
* Designing the in-memory data structure: Current Tokens, Maximum Capacity, Refill Rate, Last Refill Timestamp — and why each field is required
* Translating the design into a working Go implementation

Files Created:

* source/tokenbucket.go (Token Bucket struct + refill/consume logic)
* main.go (updated to exercise the Token Bucket manually)

Problems Faced:

* None blocking — completed without major issues

Key Learnings:

* A Token Bucket is fully described by four values: current tokens, max capacity, refill rate, and last refill timestamp
* Refill is calculated lazily — based on elapsed time since the last refill — rather than running a background timer every second
* Consumption logic must recalculate refill first, then attempt to deduct a token, so state is always current before an Allow/Deny decision is made
* Bucket size must always be capped at maximum capacity, even after a long idle period
* Validating an algorithm with manual test scenarios (full bucket, drain to empty, wait-and-refill, burst) builds real trust in the logic before adding Redis/Lua complexity
* This in-memory version is deliberately disposable — it exists to prove correctness before the same logic is ported into a Redis-backed, Lua-atomic implementation

Git Commit Created:

Yes

Commit:

feat: token bucket algorithm (in-memory)

Outcome:

✅ Working in-memory Token Bucket implementation

✅ Refill logic implemented and verified

✅ Consumption logic implemented and verified

✅ Burst handling verified (bucket never exceeds capacity)

✅ Manual test scenarios run and validated

✅ Progress tracker updated

✅ Roadmap synchronized

✅ Commit pushed: "feat: token bucket algorithm (in-memory)"

Next Objective:

Redis Lua Scripting — Atomic Operations, Redis EVAL, Race Condition Prevention (Redis-backed Token Bucket)

---

## Day 7 — Redis-backed Token Bucket (Atomic Lua Scripting)

Date:

2026-07-11

Hours Spent:

~4–5 Hours

Topics Learned:

* Why the in-memory Token Bucket struct is no longer sufficient — Go memory doesn't survive restarts and can't be shared across multiple server instances
* Race Conditions & Atomic Operations — why a GET → modify → SET sequence is dangerous, the check-then-act bug, and how it leads to double-spending tokens
* What "atomic" actually means and why splitting logic into multiple separate Redis commands is unsafe under concurrent access
* Redis Lua Scripting fundamentals — what Lua is, why Redis supports it, the `EVAL` command, and why Redis guarantees a Lua script runs as a single atomic unit
* Designing the Redis-backed Token Bucket schema — converting the Go struct into a Redis Hash (keys, fields, token storage, timestamp storage) and how one client's bucket is separated from another's
* Writing the full Lua script: read bucket → calculate elapsed time → refill → cap at capacity → consume token → save state → return result
* Connecting Go to Lua — executing scripts with `Eval()`, passing parameters from Go into Lua, reading the returned result, and wiring it into the `/check` endpoint
* Verification & concurrency validation — manually confirming persistence, refill correctness, consumption correctness, capacity capping, and that concurrent requests don't corrupt bucket state

Files Created:

* source/tokenbucket.lua (atomic refill + consume Lua script)
* source/redis_bucket.go (Go ↔ Redis Lua integration layer)
* main.go (updated `/check` endpoint to use the Redis-backed, Lua-atomic Token Bucket)

Problems Faced:

* None blocking — completed without major issues

Key Learnings:

* A race condition occurs whenever two requests can interleave between reading and writing shared state — Redis Lua scripting eliminates this by executing the entire refill-and-consume sequence as one atomic operation
* Lua is preferred over issuing multiple separate Redis commands because Redis guarantees a script runs to completion without any other command executing in between
* The bucket's state (current tokens, max capacity, refill rate, last refill timestamp) maps directly onto a Redis Hash, with each client isolated by its own key
* The exact refill/consume math from the in-memory version (Day 6) ports over unchanged — only the execution environment moves from Go into Lua
* Concurrent requests hitting the same bucket key no longer corrupt state, since Redis serializes Lua script execution
* This milestone marks the transition from a "toy" in-memory limiter to a genuinely distributed, production-style rate limiter

Git Commit Created:

Yes

Commit:

feat: redis-backed token bucket using atomic lua scripting

Outcome:

✅ Understood why the in-memory implementation could not survive concurrency or restarts

✅ Understood race conditions and why atomicity is required

✅ Learned Redis Lua Scripting and the `EVAL` command

✅ Designed the Redis Hash schema for the Token Bucket

✅ Implemented the full refill + consume + cap + save logic as an atomic Lua script

✅ Connected Go to the Lua script via `Eval()` and wired it into `/check`

✅ Verified bucket persistence, refill correctness, and capacity capping

✅ Verified no race conditions occur under concurrent requests

✅ Progress tracker updated

✅ Roadmap synchronized (no structural changes required)

✅ Commit pushed: "feat: redis-backed token bucket using atomic lua scripting"

Next Objective:

MVP Completion — Docker Compose, finalized `/check` Endpoint, Rate Limit Headers (Resume-Ready MVP)

---

## Day 8 — Resume-Ready MVP Completion (Docker Compose + Production API)

Date:

2026-07-12

Hours Spent:

~5–6 Hours

Topics Learned:

* Why "works on my machine" is a real production problem and why companies never install software manually on servers
* Container vs Image vs Docker Compose vs Multi-container Application — the conceptual difference between each
* Dockerfile fundamentals — Base Image, Layers, Build Context, `COPY`, `RUN`, `CMD`, `EXPOSE`, and why every instruction creates a new layer
* How a Go application is converted into a runnable Docker image, and a conceptual preview of multi-stage builds
* Writing the project's own `Dockerfile` — selecting a Go base image, copying source, downloading dependencies, building the binary, and running it
* Docker Compose fundamentals — `docker-compose.yml`, Services, Networks, Volumes, Port Mapping
* Writing `docker-compose.yml` to orchestrate the Go service and the Redis service together
* Container networking — Bridge Networks, Container DNS, Service Discovery, and why `localhost:6379` must become `redis:6379` inside a container network
* Running the full stack with `docker compose up`, `docker compose down`, `docker compose logs`, `docker compose ps`, and verifying both containers start, communicate, and serve `/check` correctly
* Designing a professional, consistent JSON response contract for `/check` (API Contract, Response Schema, Consistency, Metadata) instead of an ad-hoc shape
* Industry-standard HTTP rate-limit headers — `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset`, `Retry-After` — and why real APIs (GitHub, Stripe) put this metadata in headers rather than only in the JSON body
* End-to-end MVP validation — allowed requests, denied requests, refill behavior, persistence across restarts, and persistence across a full Docker restart
* Docker debugging fundamentals — `docker logs`, `docker compose logs`, `docker exec`, `docker inspect`, container exit codes, and restart policies
* Why documentation (README, progress tracker, roadmap) matters as much as code in a professional repository

Files Created:

* Dockerfile (Go base image, dependency download, build, run)
* docker-compose.yml (Go service + Redis service, networked together)
* main.go (updated `/check` response contract + rate-limit headers)
* README.md (updated to reflect Docker Compose usage and current project state)

Problems Faced:

* None blocking — completed without major issues

Key Learnings:

* Docker solves "works on my machine" by packaging the application and its exact environment together into a portable image
* Docker Compose exists on top of Docker because real applications are multi-container (Go + Redis here), and Compose lets both be defined, networked, and started together with one command
* Inside a container network, services address each other by service name (`redis:6379`) rather than `localhost`, because each container has its own isolated network namespace — Container DNS resolves service names to the right container
* A clean, consistent API contract (fixed JSON shape + standard headers) is what separates a "toy endpoint" from something that looks like a real product API
* Rate-limit metadata belongs in headers (`X-RateLimit-*`, `Retry-After`) rather than being buried in the JSON body, matching the convention used by GitHub, Twitter, and Stripe
* Validating an MVP means testing more than the happy path — refill correctness, restart persistence, and Docker-level restart persistence all had to be manually verified before calling this milestone done
* This milestone is the official transition point: the project is no longer "a correct algorithm," it is now a containerized, resume-ready backend service that runs end-to-end with a single `docker compose up`

Git Commit Created:

Yes

Commit:

feat: resume ready MVP using docker compose

Outcome:

✅ Understood why containerization is necessary for production-style backend services

✅ Wrote the project's Dockerfile (Go base image → build → run)

✅ Wrote docker-compose.yml orchestrating the Go service and Redis service

✅ Understood and verified container networking (`redis:6379` instead of `localhost:6379`)

✅ Verified the full stack starts and communicates via `docker compose up`

✅ Designed a clean, professional `/check` JSON response contract

✅ Added standard HTTP rate-limit headers (`X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset`, `Retry-After`)

✅ Verified allowed requests, denied requests, refill, and persistence across restarts (including Docker restarts)

✅ Learned baseline Docker debugging (`docker logs`, `docker compose logs`, `docker exec`, `docker inspect`)

✅ Updated README, progress tracker, and roadmap for synchronization

✅ Progress tracker updated

✅ Roadmap synchronized

✅ Commit pushed: "feat: resume ready MVP using docker compose"

✅ **Resume-Ready MVP officially complete**

Next Objective:

Admin API — `POST/GET/PUT/DELETE /admin/clients`, Redis-backed client configuration CRUD (Phase 3 — Advanced Features)

---

## Day 9 — Admin API (Redis-backed Client Configuration)

Date:

2026-07-13

Hours Spent:

~5–6 Hours

Topics Learned:

* Why production systems separate public-facing APIs from admin APIs, and why admin functionality is modeled as its own set of endpoints rather than bolted onto `/check`
* CRUD (Create, Read, Update, Delete) as a pattern, and Resource-Oriented API Design — designing around a `ClientConfig` resource instead of ad-hoc actions
* HTTP method semantics for CRUD: `POST` (create), `GET` (read), `PUT` (update), `DELETE` (delete)
* Request Body vs Query Parameters vs Path Parameters, and when each is the appropriate way to pass data into an endpoint
* Designing the `ClientConfig` data model — Client ID, Rate, Burst, and an Algorithm field (kept as a preview field for the future Sliding Window milestone)
* Why per-client configuration belongs in Redis rather than in-process memory — same durability/shared-state reasoning as the bucket state itself
* Redis Hashes for structured configuration data, and key design (`rl:cfg:{client_id}`) so each client's config is isolated
* The idea of Default Configuration vs Custom Configuration — a client with no stored config should transparently fall back to sane defaults instead of erroring
* End-to-end integration discipline — building the routes first with in-memory handling, then wiring Redis underneath, then re-verifying `/check` still behaves correctly for both configured and unconfigured clients

Files Created:

* source/config.go (`ClientConfig` struct + default configuration values)
* source/redis_config.go (Redis Hash CRUD logic for client configs — create, read, update, delete)
* main.go (added `POST/GET/PUT/DELETE /admin/clients` routes; updated `/check` to load per-client config from Redis, falling back to defaults when none exists)

Problems Faced:

* None blocking — completed without major issues

Key Learnings:

* Separating admin (configuration) concerns from the hot-path `/check` endpoint keeps the rate limiter engine decoupled from how its configuration is managed — the engine doesn't care whether config came from a default or from an admin-created override
* Modeling `/admin/clients` as a resource (not a bag of actions) is what makes the CRUD verbs map cleanly onto HTTP methods
* A Redis Hash is a natural fit for `ClientConfig` for the same reason it was a natural fit for bucket state on Day 7 — grouped fields under one isolated key per client
* Falling back to default configuration when a Redis key is missing is what keeps `/check` safe for clients that were never explicitly configured — the system degrades gracefully instead of failing
* The project is no longer a single hardcoded rate limit — it is now a configurable, per-client backend service, which is the entry point into Phase 3 — Advanced Features

Git Commit Created:

Yes

Commit:

feat: admin api for client configuration

Outcome:

✅ Understood Admin API design and CRUD resource modeling

✅ Designed the `ClientConfig` data model (Client ID, Rate, Burst, Algorithm)

✅ Implemented all four CRUD endpoints (`POST/GET/PUT/DELETE /admin/clients`)

✅ Implemented Redis Hash-backed storage for client configurations (`rl:cfg:{client_id}`)

✅ Wired `/check` to load per-client config from Redis, falling back to defaults when none exists

✅ Manually tested Create, Read, Update, Delete against `/admin/clients`

✅ Verified `/check?client_id=A` and `/check?client_id=B` respect independent per-client limits

✅ Verified unknown clients continue to use default configuration correctly

✅ Progress tracker updated

✅ Roadmap synchronized

✅ Commit pushed: "feat: admin api for client configuration"

✅ **Admin API milestone officially complete**

Next Objective:

Sliding Window — Sorted Sets, Window Management, Alternative Algorithm (Phase 3 — Advanced Features)

---

## Day 10 — Sliding Window Rate Limiter (Redis Sorted Sets)

Date:

2026-07-14

Hours Spent:

~5–6 Hours

Topics Learned:

* Why Token Bucket alone isn't sufficient, and why Sliding Window was invented as a stricter, more precise alternative
* Conceptual differences between Fixed Window, Token Bucket, and Sliding Window — burst tolerance vs. strictness tradeoffs, and why companies keep more than one algorithm available
* Redis Sorted Sets (ZSET) as a new data structure — Score, Member, and why ordering (not just grouping, as with Hashes) was needed
* Core Sorted Set commands required for this project: `ZADD`, `ZREMRANGEBYSCORE`, `ZCARD`, and basic `ZRANGE` — nothing beyond these
* Manually experimenting in the Redis CLI — creating a Sorted Set, inserting timestamps, removing expired entries, and counting what remains, before writing any Go code
* The Sliding Window algorithm itself: remove expired timestamps → count remaining requests → if count is under the limit, allow and insert the current timestamp → otherwise, deny
* Manually solving Sliding Window request patterns on paper before coding (e.g. 8 requests against a 5-per-10-second limit) to understand the algorithm mathematically first
* Porting the algorithm into a Redis-backed, Lua-atomic implementation — reusing the same atomicity approach proven with Token Bucket on Day 7
* Wiring the Sliding Window path into the existing `/check` endpoint alongside Token Bucket
* Manually comparing Sliding Window vs. Token Bucket behavior side-by-side under identical request patterns, and articulating the tradeoffs between the two

Files Created:

* scripts/sliding_window.lua (atomic expire + count + insert Lua script using Sorted Sets)
* source/redis_window.go (Go ↔ Redis Sorted Set integration layer for Sliding Window)
* main.go (updated to route requests to Sliding Window when configured, alongside the existing Token Bucket path)

Problems Faced:

* None blocking — completed without major issues

Key Learnings:

* Sorted Sets store timestamps as scored members, which makes "remove everything older than X" (`ZREMRANGEBYSCORE`) and "count what's left" (`ZCARD`) a natural fit for a moving time window — something Hashes alone couldn't express
* Sliding Window is fundamentally stricter than Token Bucket: it enforces a hard cap on requests within any rolling window with no burst allowance beyond the limit, whereas Token Bucket explicitly permits bursts up to bucket capacity
* The same atomicity principle from Day 7 (wrap the entire read-modify-write sequence in a single Lua script) applies identically here — expire-then-count-then-insert must happen as one atomic unit to prevent double-counting under concurrent requests
* Comparing both algorithms against identical request patterns made the tradeoff concrete rather than theoretical — Token Bucket allowed a burst that Sliding Window correctly denied, which is exactly the kind of before/after evidence worth bringing to an interview
* The rate limiter engine now genuinely supports two distinct algorithms end-to-end, which sets up the next step: extracting a common Strategy interface so `/check` doesn't need to know which algorithm it's calling

Git Commit Created:

Yes

Commit:

feat: sliding window rate limiter using redis sorted sets

Outcome:

✅ Understood why Sliding Window was invented and when it's preferred over Token Bucket

✅ Learned Redis Sorted Sets (`ZADD`, `ZREMRANGEBYSCORE`, `ZCARD`, basic `ZRANGE`)

✅ Solved Sliding Window request patterns manually before implementing them

✅ Implemented the Sliding Window algorithm as an atomic Lua script using Sorted Sets

✅ Connected Go to the new Lua script via `redis_window.go` and wired it into `/check`

✅ Added new Redis key pattern `rl:window:{clientID}` for per-client request timestamps

✅ Manually verified allow/deny behavior and window movement as time progressed

✅ Manually compared Sliding Window vs. Token Bucket under identical request patterns

✅ Progress tracker updated

✅ Roadmap synchronized (no structural changes required)

✅ Commit pushed: "feat: sliding window rate limiter using redis sorted sets"

✅ **Sliding Window milestone officially complete**

Next Objective:

Strategy Pattern — unify Token Bucket and Sliding Window behind a common rate-limiter interface (Phase 3 — Advanced Features)

---

# 🎯 CURRENT MILESTONE

## Strategy Pattern

Objectives:

* Token Bucket
* Sliding Window

Through a common interface.

Deliverable:

A shared rate-limiter interface/strategy abstraction that `/check` calls without needing to know which algorithm is behind it.

Completion Criteria:

* Common interface defined for both algorithms
* Token Bucket and Sliding Window both implement it
* `/check` refactored to select a strategy rather than branching on algorithm type
* Progress tracker and roadmap updated to reflect completion

Status:

⬜ Pending

---

# 📈 RESUME READINESS TRACKER

```text
Foundation
██████████ 100%

Backend Basics
██████████ 100%

Rate Limiter Core
██████████ 100%

Advanced Features
█████░░░░░ 50%

Deployment
░░░░░░░░░░ 0%

Documentation
█░░░░░░░░░ 10%
```

Current Estimated Progress:

~60%

Note: Rate Limiter Core hit 100% with the completion of the Resume-Ready MVP (Token Bucket + Redis persistence + atomic Lua + Docker Compose + Rate Limit Headers). Advanced Features moved to 50% with the completion of the Sliding Window milestone (Redis Sorted Sets + atomic Lua-based window algorithm) on Day 10, following the Admin API on Day 9 — Strategy Pattern and Robustness remain pending within this phase. Documentation ticked up slightly from the README/progress/roadmap sync done on Day 8, but full documentation (architecture diagram, API reference, load test results) is still pending until Phase 6.

---

# 📝 PROJECT JOURNAL TEMPLATE

For every study/build session:

Date:

Hours Spent:

Topics Learned:

Files Created:

Problems Faced:

Key Learnings:

Next Objective:

Git Commit Created:

Yes / No

---

# 🤖 ASSISTANT CONTINUATION PROTOCOL

Whenever this file is loaded into a new chat:

1. Read roadmap file first.
2. Read this progress file second.
3. Sync roadmap and progress.
4. Identify current milestone.
5. Continue from next unfinished task.
6. Recommend any new files/folders required.
7. Update workspace structure if necessary.
8. Ensure Git workflow is being followed.
9. Update completed milestones.
10. Maintain roadmap alignment.
11. Treat this file as the authoritative source of project progress.

This document is the single source of truth for the entire project lifecycle.