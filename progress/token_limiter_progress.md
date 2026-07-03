# 🚀 Distributed Rate Limiter Project Tracker (v2)

---

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

⏳ Current Milestone

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

⬜ Pending

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

⬜ Pending

---

### Redis Lua Scripting

Objectives:

* Atomic Operations
* Redis EVAL
* Race Condition Prevention

Deliverable:

Redis-backed Token Bucket

Status:

⬜ Pending

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

⬜ Pending

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

⬜ Pending

---

### Sliding Window

Objectives:

* Sorted Sets
* Window Management
* Alternative Algorithm

Status:

⬜ Pending

---

### Strategy Pattern

Objectives:

* Token Bucket
* Sliding Window

Through common interface.

Status:

⬜ Pending

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

Future additions:

```text
cmd/
internal/
scripts/
loadtest/
Dockerfile
docker-compose.yml
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

⬜ Gin Hello World

⬜ Redis Connectivity

⬜ Token Bucket

⬜ Lua Integration

⬜ Docker Compose

⬜ Admin API

⬜ Sliding Window

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

---

# 🎯 CURRENT MILESTONE

## HTTP Fundamentals + Gin Framework

Objectives:

* Understand HTTP
* Learn Request / Response Flow
* Learn Status Codes
* Learn Headers
* Learn REST Basics
* Learn Gin Framework
* Build First HTTP Server

Deliverable:

```text
GET /ping
```

Expected Response:

```json
{
    "message": "pong"
}
```

Completion Criteria:

* Understand HTTP basics
* Build Gin server
* Successfully return JSON response
* Understand routing fundamentals

Status:

⏳ In Progress

---

# 📈 RESUME READINESS TRACKER

```text
Foundation
██████████ 100%

Backend Basics
██░░░░░░░░ 20%

Rate Limiter Core
░░░░░░░░░░ 0%

Advanced Features
░░░░░░░░░░ 0%

Deployment
░░░░░░░░░░ 0%

Documentation
░░░░░░░░░░ 0%
```

Current Estimated Progress:

~15%

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
