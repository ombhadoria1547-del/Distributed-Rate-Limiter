# Performance Report

## Purpose

This document summarizes the performance characteristics of the Distributed Rate Limiter under concurrent load.

The objective of these benchmarks was to measure throughput, latency, saturation behavior, and overall system stability after correctness had already been verified.

---

# Test Environment

| Component | Specification |
|------------|---------------|
| CPU | Intel Core i5 (13th Gen) |
| Memory | 16 GB RAM |
| Operating System | Windows |
| Runtime | Docker |
| Backend | Go + Gin |
| Database | Redis 8 |
| Benchmark Tool | k6 |

---

# Benchmark Configuration

The following benchmark configuration was used throughout all performance tests.

- Test Duration: **10 seconds**
- Virtual Users (VUs): **10 → 70**
- Benchmark Tool: **k6**
- Error Rate Target: **0%**
- Benchmark Mode Enabled (`BENCHMARK_MODE`)
- All benchmark runs executed using the same benchmark suite with only the Virtual User count varied.

---

# Benchmark Results

| Virtual Users | Throughput (req/s) | Avg Latency | p95 Latency | Max Latency |
|---------------|-------------------:|------------:|------------:|------------:|
| 10 | 7,820.44 | 1.22 ms | 1.71 ms | 6.52 ms |
| 20 | 11,556.20 | 1.67 ms | 2.62 ms | 10.70 ms |
| 30 | 14,019.40 | 2.08 ms | 3.21 ms | 16.31 ms |
| 40 | 15,759.04 | 2.47 ms | 3.86 ms | 15.83 ms |
| 50 | 17,169.20 | 2.85 ms | 4.42 ms | 19.42 ms |
| 60 | **17,650.25** | 3.33 ms | 5.34 ms | 26.86 ms |
| 70 | 16,736.15 | 4.11 ms | 6.86 ms | 29.27 ms |

All benchmark runs completed with **0% errors**.

---

# Performance Analysis

## Throughput

Throughput increased steadily as concurrency increased from 10 to 60 Virtual Users.

The highest sustained throughput achieved during testing was:

**≈17.65k requests/second**

Increasing concurrency beyond this point did not improve throughput.

---

## Latency

Latency increased gradually as concurrency increased, which is expected for a saturated system.

Even at peak throughput:

- Average latency remained approximately **3.33 ms**
- p95 latency remained approximately **5.34 ms**

This indicates that the service maintained low response times while processing high request volumes.

---

## Saturation Analysis

The service reached its saturation point at approximately **60 Virtual Users**.

Beyond this point:

- Throughput began to decrease.
- Latency continued increasing.

This behavior indicates that the application had reached the processing capacity of the development hardware, with additional concurrency contributing primarily to queueing delay rather than increased throughput.

---

# Engineering Observations

During benchmarking, an unexpected throughput limitation was observed.

Initial investigation considered:

- Redis
- Lua script execution
- Network transport

Systematic debugging identified the actual bottleneck as **Gin's synchronous request logger**, enabled through `gin.Default()`.

To eliminate this overhead without modifying application logic, a configuration-driven benchmark mode (`BENCHMARK_MODE`) was introduced.

Benchmark mode starts the server using:

- `gin.New()`
- `gin.Recovery()`

while retaining `gin.Default()` for normal development.

This allowed production behavior and benchmark behavior to be switched through configuration rather than source code changes.

---

# Summary

- Peak Throughput: **≈17.65k requests/sec**
- Peak Concurrency: **60 Virtual Users**
- Average Latency at Peak: **3.33 ms**
- p95 Latency at Peak: **5.34 ms**
- Maximum Observed Latency: **29.27 ms**
- Error Rate: **0%**
- Saturation Point: **≈60 Virtual Users**

---

# Conclusion

The Distributed Rate Limiter demonstrated stable performance under concurrent load, sustaining approximately **17.65k requests per second** on the test hardware while maintaining low latency and zero observed request errors.

Performance testing also led to the discovery and resolution of a middleware-related bottleneck, improving throughput through configuration rather than changes to the core rate-limiting algorithms.

The benchmark results indicate that the service comfortably exceeds the original design target of **500+ requests per second**, achieving more than **30×** the required throughput while preserving correctness under concurrent access.