export const BURST = 10;

export const TOTAL_REQUESTS = BURST + 1;

export const REFILL_WAIT_MS = 3000;

export const REFILL_RATE = 2;

export const EXPECTED_REFILL_TOKENS =
    (REFILL_WAIT_MS / 1000) * REFILL_RATE;