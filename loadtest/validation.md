Note: This document summarizes the validation process. The corresponding automated test scripts are available under loadtest/scripts/.

# Validation Report

## Purpose

This document records the verification process used to validate the correctness of the Distributed Rate Limiter.

The objective of these tests was not to measure performance, but to verify that both implemented algorithms behaved exactly as expected under normal and concurrent workloads.

---

# Validation Methodology

The validation process was divided into five independent stages:

1. Baseline Verification
2. Token Bucket Validation
3. Sliding Window Validation
4. Redis State Inspection
5. Load Testing Validation

Each stage verified a different aspect of the system.

---

# 1. Baseline Verification

## Objective

Verify that the service behaves correctly under normal request loads before introducing concurrency.

## Verification

- Verified `/check` endpoint returns the expected HTTP responses.
- Verified Allow and Deny decisions.
- Verified response headers.
- Verified default configuration fallback.
- Verified custom client configuration through the Admin API.

## Result

✅ Passed

---

# 2. Token Bucket Validation

## Objective

Verify the correctness of the Token Bucket implementation.

### Burst Validation

Verified that:

- requests are allowed until bucket capacity is exhausted
- the first request exceeding the burst capacity receives HTTP 429
- bucket capacity is never exceeded

Result:

✅ Passed

---

### Refill Mathematics Validation

Verified the refill equation:

```

New Tokens = Old Tokens + (Elapsed Time × Refill Rate)

```

The Redis bucket state (`tokens`, `last_refill`) matched the expected mathematical calculations.

Result:

✅ Passed

---

# 3. Sliding Window Validation

## Objective

Verify the correctness of the Sliding Window implementation.

### Capacity Validation

Verified that:

- requests are accepted until window capacity is reached
- requests exceeding the configured limit are rejected

Result:

✅ Passed

---

### Expiration Validation

Verified that expired timestamps are removed correctly using:

```

ZREMRANGEBYSCORE

```

Confirmed that new requests are accepted once older entries expire and the window advances.

Result:

✅ Passed

---

# 4. Redis State Inspection

## Objective

Verify that Redis stores the expected state throughout execution.

The following data structures were inspected during testing:

### Token Bucket

Redis Hash

Verified:

- tokens
- last_refill

### Sliding Window

Redis Sorted Set

Verified:

- request timestamps
- expired entry removal
- current window contents

### Client Configuration

Redis Hash

Verified:

- rate
- burst
- algorithm

Result:

✅ Passed

---

# 5. Load Testing Validation

Correctness was verified again while the system was under concurrent load.

Verified:

- Token Bucket burst enforcement
- Token Bucket refill mathematics
- Sliding Window capacity enforcement
- Sliding Window expiration
- Redis state consistency
- Zero double-spend violations

Result:

✅ Passed

---

# Overall Validation Summary

The Distributed Rate Limiter successfully passed every verification stage.

Validated features:

- Baseline functionality
- Token Bucket correctness
- Sliding Window correctness
- Redis state consistency
- Concurrent correctness
- Zero double-spend behavior

Overall Status:

✅ Validation Successful