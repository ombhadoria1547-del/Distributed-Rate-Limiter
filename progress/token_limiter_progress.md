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

✅ Completed

---

### Robustness

Objectives:

* Validation
* Logging
* Fail Open / Fail Closed

Status:

✅ Completed

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

Current Additions (as of Day 11):

```text
Dockerfile
docker-compose.yml
scripts/
    sliding_window.lua
source/
    redis_window.go
    limiter.go
    strategy.go
    validation.go
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

✅ Strategy Pattern & Robustness

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

## Day 11 — Strategy Pattern & Production Robustness

Date:

2026-07-15

Hours Spent:

~5–6 Hours

Topics Learned:

* Why letting `/check` branch on algorithm type (`if algo == token_bucket ... else if algo == sliding_window ...`) is an architectural weakness — every new algorithm would force an edit to already-working code
* The Open/Closed Principle (SOLID) — software should be open for extension but closed for modification, illustrated with the shipping-method analogy (Road/Air/Express/Drone/Pickup all implementing one interface instead of branching in checkout code)
* The Strategy Pattern itself — a behavioral design pattern where multiple algorithms share one common interface and become interchangeable at runtime (`Limiter.Allow()` instead of `if TokenBucket ... else SlidingWindow`)
* Why the pattern was invented historically — separating "what" (the interface) from "how" (the algorithm behind it) to avoid ever-growing if/else and switch chains
* Where Strategy Pattern shows up across the industry — compression algorithms, payment gateways, authentication providers, sorting libraries, caching engines, database drivers, load balancers, routing engines, cloud SDKs
* Refactoring discipline — moving algorithm-specific logic out of the HTTP handler so `/check` only loads config, selects a strategy, invokes the interface, and formats the response, while the algorithms themselves stay HTTP-agnostic (conceptual dependency inversion)
* Input Validation — why malformed input (missing `client_id`, negative rate/burst, unsupported algorithm names, invalid JSON) must be rejected early with clear `400 Bad Request` responses instead of being allowed to corrupt behavior
* Logging fundamentals — why logs are the primary source of truth once a service is running without a debugger attached, and keeping the hot path free of noisy per-request logs while still logging startup and significant errors
* Fail-Open vs. Fail-Closed — the tradeoff between availability (allow requests when Redis is down) and correctness/protection (deny requests when enforcement can't be guaranteed), and choosing Fail-Closed for this project since it's easier to defend in an interview and better protects the service
* Re-verification discipline — confirming that a refactor changes structure, not behavior, by re-testing Token Bucket, Sliding Window, the Admin API, default configuration, and the full Docker stack after the rewrite

Files Created:

* source/limiter.go (`RateLimiter` interface definition — a single `Allow(clientID)` contract both algorithms satisfy)
* source/strategy.go (selector/factory that picks Token Bucket or Sliding Window based on per-client config)
* source/validation.go (request validation logic — missing/invalid `client_id`, negative rate/burst, unsupported algorithm names)
* source/tokenbucket.go, source/redis_bucket.go, source/redis_window.go (updated so both algorithms satisfy the new `RateLimiter` interface instead of being called directly)
* main.go (`/check` refactored to call the `RateLimiter` interface instead of branching on algorithm type; added validation error responses, basic logging, and Fail-Closed handling for Redis failures)

Problems Faced:

* None blocking — completed without major issues

Key Learnings:

* Branching on algorithm type inside `/check` was fine for two algorithms but would not scale — the Strategy Pattern removes that coupling entirely by making `/check` depend only on an interface, not on any specific implementation
* The Open/Closed Principle is the real justification for the refactor: adding a future algorithm (e.g. Fixed Window, Leaky Bucket) now means implementing the interface, not editing `/check`
* A clean interface boundary (`Allow()`) is what makes Token Bucket and Sliding Window genuinely interchangeable at runtime, which is the same idea used by payment gateways, compression libraries, and database drivers industry-wide
* Robustness is a different concern from correctness — the algorithms were already correct, but validation, logging, and a defined Fail-Open/Fail-Closed policy are what separate a "working prototype" from something that behaves predictably in production
* Choosing Fail-Closed over Fail-Open is a defensible, explainable tradeoff (protection over availability) rather than an arbitrary choice — exactly the kind of design decision worth bringing to an interview
* Re-testing everything after the refactor confirmed the golden rule of refactoring: behavior stayed identical while the internal structure got significantly cleaner

Git Commit Created:

Yes

Commit:

refactor: strategy pattern for rate limiter + robustness (validation, logging, fail-closed)

Outcome:

✅ Understood why algorithm-branching in `/check` doesn't scale, and why the Open/Closed Principle applies here

✅ Understood the Strategy Pattern and where it's used across the industry

✅ Defined a common `RateLimiter` interface

✅ Token Bucket implements the `RateLimiter` interface

✅ Sliding Window implements the `RateLimiter` interface

✅ `/check` refactored to call the interface instead of branching on algorithm type

✅ Added request validation with clear `400 Bad Request` error responses

✅ Added basic logging (startup, significant errors, Redis failures)

✅ Chose and implemented a Fail-Closed policy for Redis failures

✅ Re-tested Token Bucket, Sliding Window, Admin API, and default configuration — all unchanged after the refactor

✅ Re-verified full Docker stack (`docker compose up`) still behaves correctly post-refactor

✅ Progress tracker updated

✅ Roadmap synchronized (no structural changes required)

✅ Commit pushed: "refactor: strategy pattern for rate limiter + robustness (validation, logging, fail-closed)"

✅ **Strategy Pattern & Robustness milestone officially complete**

Next Objective:

Load Testing — Throughput Validation, Concurrency Validation, Correctness Validation against a 500+ requests/second target (Phase 4 — Validation)

---

## Day 12 — Load Testing & Correctness Validation (k6)

Date:

2026-07-16

Hours Spent:

~7–8 Hours

Topics Learned:

* Why Load Testing exists as its own engineering discipline — the difference between Functional Testing (does it work?) and Performance Testing (does it work under realistic concurrent load?), and why passing manual/Postman-level tests says nothing about production readiness
* Performance Engineering fundamentals — Throughput, Latency, Capacity, Reliability, Bottlenecks, Resource Utilization, and why "it responds in 20ms" is meaningless without stating the concurrency it was measured under
* k6 as a tool — Virtual Users (VUs), Scenarios, Executors, the JavaScript scripting model, built-in metrics/thresholds, and why k6 was chosen over Vegeta/JMeter/Locust for this project
* Performance metrics in depth — Throughput vs. Latency, why p95/p99 percentiles matter far more than averages, Error Rate, Saturation, and how to read a k6 report as engineering evidence rather than a screenshot
* Correctness Validation as a discipline distinct from performance — why a faster rate limiter that over-allows requests is strictly worse than a slower one that stays mathematically correct, using the "double-spend" framing (two requests must never consume the same token)
* Designing a validation strategy *before* writing any test code — splitting testing into deliberate scenarios (Baseline, Token Bucket Validation, Sliding Window Validation, Sustained Load, Stress Observation) instead of one undirected benchmark
* Building a reusable, version-controlled load-testing suite rather than a one-off script — so the exact same test can be re-run later to compare before/after any optimization
* Hypothesis-driven performance debugging — systematically ruling out Redis, Lua, and transport-layer causes before correctly isolating the real bottleneck
* Middleware as a hidden performance bottleneck — how a synchronous request logger can silently cap throughput far below what the actual rate-limiting logic is capable of
* Configuration-driven application behavior — using environment variables to toggle application behavior (debug logging, benchmark mode) instead of editing source code between runs
* Reading saturation from a throughput/latency curve — identifying the concurrency level at which added load stops increasing throughput and starts only increasing latency

Files Created:

* loadtest/scripts/baseline.js (low-traffic sanity test verifying k6 ↔ Go ↔ Redis connectivity end-to-end)
* loadtest/scripts/token_bucket.js (Token Bucket correctness scenario — burst capacity + refill mathematics)
* loadtest/scripts/sliding_window.js (Sliding Window correctness scenario — capacity + expiration)
* loadtest/scripts/stress.js (sustained/stress load scenario for throughput and saturation testing)
* loadtest/lib/config.js (centralized, reusable benchmark configuration — client IDs, VUs, duration)
* loadtest/lib/token_bucket.js, loadtest/lib/sliding_window.js, loadtest/lib/stress.js (shared benchmark helper logic)
* main.go (added `BENCHMARK_MODE` environment variable — switches between `gin.Default()` for development and `gin.New()` + `gin.Recovery()` for benchmarking, plus a startup log confirming the active mode)
* docker-compose.yml (updated to pass `BENCHMARK_MODE` through as an environment variable, so behavior changes via configuration, not code edits)

Problems Faced:

* Initial benchmark runs produced inconsistent results — throughput would collapse and pause for minutes around ~27k cumulative requests, with the terminal continuing to print well after the test should have ended
* Root cause was not obvious — Redis, the Lua scripts, and the transport layer were all initially suspected and individually investigated
* Systematically ruled out, in order: Redis as the bottleneck, Lua script execution as the bottleneck, and transport/network failures — none of these were the cause
* Eventually isolated the true cause: Gin's default request logger (`gin.Logger()`, enabled via `gin.Default()`) was synchronously logging every single HTTP request, which throttled throughput far below what the rate-limiting logic itself could sustain

Key Learnings:

* A fast algorithm can still be bottlenecked by something completely unrelated to its own logic — the token bucket and sliding window implementations were never the problem; the request logging middleware was
* Hypothesis-driven debugging (systematically eliminating Redis → Lua → transport → middleware) is far more reliable than guessing, and mirrors how real infrastructure incidents get root-caused
* Separating "development mode" from "benchmark mode" via configuration (`BENCHMARK_MODE`) rather than commenting out code is what makes a benchmark trustworthy and repeatable
* Correctness must be validated independently of performance — the Token Bucket and Sliding Window scenarios proved burst capacity, refill math, window expiration, and capacity enforcement were all exactly correct, with zero double-spend behavior observed under concurrent traffic
* Percentile latency (p95, max) tells a very different — and more honest — story than average latency, especially as concurrency increases
* Throughput does not increase indefinitely with concurrency — the service's throughput/latency curve clearly showed a saturation point, beyond which added Virtual Users only increased latency without increasing throughput
* Redis Hash/Sorted Set state (`bucket:{client_id}`, `rl:window:{client_id}`) needed to be reset to a known state before each correctness run to keep experiments clean and reproducible, especially at low refill rates
* A load-testing suite is only genuinely useful if it's reusable — building it as versioned scripts under `loadtest/` (rather than ad-hoc terminal commands) means the exact same benchmark can be re-run after future optimizations to prove impact

Benchmark Results (10-second runs, 0% errors across all runs):

| Virtual Users | Throughput (req/s) | Avg Latency | p95 | Max |
|---|---|---|---|---|
| 10 | 7,820.44 | 1.22 ms | 1.71 ms | 6.52 ms |
| 20 | 11,556.20 | 1.67 ms | 2.62 ms | 10.70 ms |
| 30 | 14,019.40 | 2.08 ms | 3.21 ms | 16.31 ms |
| 40 | 15,759.04 | 2.47 ms | 3.86 ms | 15.83 ms |
| 50 | 17,169.20 | 2.85 ms | 4.42 ms | 19.42 ms |
| 60 | 17,650.25 | 3.33 ms | 5.34 ms | 26.86 ms |
| 70 | 16,736.15 | 4.11 ms | 6.86 ms | 29.27 ms |

Test environment: Intel Core i5 (13th Gen), 16 GB RAM, Windows + Docker, Redis 8, Go + Gin, benchmarked with k6.

Saturation Analysis:

Peak throughput of ≈17.65k req/sec was reached at 60 Virtual Users. Beyond that point throughput slightly decreased while latency continued climbing — the service saturates at ~60 concurrent VUs on this development hardware, comfortably clearing the 500+ req/sec target by more than 30x.

Correctness Validation Performed:

* Token Bucket — burst capacity enforced exactly (429 returned immediately once the configured burst was exhausted); refill mathematics verified against Redis Hash state (`tokens`, `last_refill`) matching `New Tokens = Old Tokens + (Elapsed Time × Refill Rate)` precisely
* Sliding Window — capacity enforced exactly (21st request denied against a 20-request/10-second window); expiration verified, with `ZREMRANGEBYSCORE` correctly evicting expired timestamps and requests being accepted again once the window rolled forward
* Zero double-spend violations observed under concurrent traffic across all scenarios

Git Commit Created:

Yes

Commit:

test: add k6 load testing suite and validation report

Outcome:

✅ Understood why Load Testing and Performance Engineering are mandatory before deployment, distinct from functional testing

✅ Understood k6 (Virtual Users, Scenarios, metrics, thresholds) and why it was chosen over Vegeta for this project

✅ Understood performance metrics — Throughput, Latency, Percentiles (p50/p95/p99), Error Rate, Saturation, Bottlenecks

✅ Understood Correctness Validation as a discipline distinct from and prioritized above raw performance

✅ Designed a full validation strategy (Baseline, Token Bucket, Sliding Window, Sustained Load, Stress Observation) before writing any test code

✅ Built a reusable, version-controlled k6 load-testing suite (`loadtest/scripts`, `loadtest/lib`)

✅ Verified Token Bucket correctness under concurrent load (burst capacity + refill mathematics)

✅ Verified Sliding Window correctness under concurrent load (capacity enforcement + expiration)

✅ Confirmed zero double-spend violations and 0% error rate across all benchmark runs

✅ Diagnosed and fixed a real performance bottleneck (Gin's synchronous request logger) via hypothesis-driven debugging

✅ Added configuration-driven `BENCHMARK_MODE` (Go + Docker Compose) to toggle development vs. benchmark behavior without code changes

✅ Sustained throughput far beyond the 500+ req/sec target — peak ≈17.65k req/sec at 60 VUs, saturation identified beyond that point

✅ Captured full performance metrics (throughput, avg/p95/max latency) across a 10–70 VU sweep for documentation

✅ Progress tracker updated

✅ Roadmap synchronized (no structural changes required)

✅ Commit pushed: "test: add k6 load testing suite and validation report"

✅ **Load Testing & Correctness Validation milestone officially complete**

Next Objective:

Deployment — deploy the containerized service to a public host (Render/Railway/Fly.io) so the project has a live, recruiter-shareable URL (Phase 5 — Deployment)

---

# 🎯 CURRENT MILESTONE

## Deployment

Objectives:

* Public Hosting
* Live Demo URL
* Managed/Containerized Redis in Production

Target:

A publicly reachable, recruiter-shareable live URL running the full Docker Compose stack (Go service + Redis).

Deliverable:

The Distributed Rate Limiter deployed to a free-tier host (Render, Railway, or Fly.io), with `/check` and `/admin/*` reachable over the public internet and behaving identically to the local Docker Compose stack.

Completion Criteria:

* Hosting platform selected (Render / Railway / Fly.io)
* Application deployed with Redis persistence intact
* `/check` and `/admin/*` verified reachable and correct over the public URL
* Live URL documented in README and progress tracker
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
██████████ 100%

Validation (Load Testing & Correctness)
██████████ 100%

Deployment
░░░░░░░░░░ 0%

Documentation
██░░░░░░░░ 20%
```

Current Estimated Progress:

~80%

Note: Rate Limiter Core hit 100% with the completion of the Resume-Ready MVP (Token Bucket + Redis persistence + atomic Lua + Docker Compose + Rate Limit Headers). Advanced Features reached 100% on Day 11 with the completion of the Strategy Pattern refactor and Robustness (validation, logging, Fail-Closed policy). Validation reached 100% on Day 12 with a full k6 load-testing suite, sustained throughput of ≈17.65k req/sec (far beyond the 500+ req/sec target), and mathematically verified correctness (zero double-spend violations) for both Token Bucket and Sliding Window under concurrent load — this closes out Phase 4 — Validation entirely. Documentation moved to 20% with the benchmark results and engineering log captured on Day 12, ready to feed into the final README; full documentation (architecture diagram, complete API reference) is still pending until Phase 6. Next up: Phase 5 — Deployment.

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