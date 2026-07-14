import http from "k6/http";
import { check, sleep } from "k6";

import {
    BASE_URL,
    CLIENT_ID,
} from "../lib/config.js";

import {
    BURST,
    TOTAL_REQUESTS,
    REFILL_WAIT_MS,
    EXPECTED_REFILL_TOKENS,
} from "../lib/tokenBucketConfig.js";

export const options = {
    vus: 1,
    iterations: 1,
};

export default function () {

    let allowed = 0;
    let denied = 0;

    for (let i = 0; i < TOTAL_REQUESTS; i++) {

        const response = http.get(
            `${BASE_URL}/check?client_id=${CLIENT_ID}`
        );

        const body = response.json();

        if (body.allowed) {
            allowed++;
        } else {
            denied++;
        }

    }

    check(null, {
        "allowed requests equal burst": () => allowed === BURST,

        "exactly one request denied": () =>
            denied === TOTAL_REQUESTS - BURST,
    });

    console.log("");

    console.log("==================================");
    console.log("Token Bucket Validation Summary");
    console.log("==================================");

    console.log(`Expected Allowed : ${BURST}`);
    console.log(`Actual Allowed   : ${allowed}`);

    console.log(
        `Expected Denied  : ${TOTAL_REQUESTS - BURST}`
    );
    console.log(`Actual Denied    : ${denied}`);

    console.log("==================================");
    console.log("");

    console.log("");
    console.log("Starting refill validation...");
    console.log("");

    sleep(REFILL_WAIT_MS / 1000);

    let refillAllowed = 0;
    let refillDenied = 0;

    for (
        let i = 0;
        i < EXPECTED_REFILL_TOKENS + 1;
        i++
    ) {
        const response = http.get(
            `${BASE_URL}/check?client_id=${CLIENT_ID}`
        );

        const body = response.json();

        if (body.allowed) {
            refillAllowed++;
        } else {
            refillDenied++;
        }
    }

    check(null, {
        "refill regenerated expected tokens": () =>
            refillAllowed === EXPECTED_REFILL_TOKENS,

        "refill denied exactly one request": () =>
            refillDenied === 1,
    });

    console.log("");

    console.log("==================================");
    console.log("Refill Validation Summary");
    console.log("==================================");

    console.log(
        `Expected Allowed : ${EXPECTED_REFILL_TOKENS}`
    );
    console.log(
        `Actual Allowed   : ${refillAllowed}`
    );

    console.log("Expected Denied  : 1");
    console.log(
        `Actual Denied    : ${refillDenied}`
    );

    console.log("==================================");
    console.log("");

}