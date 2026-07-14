import http from "k6/http";
import { check, sleep } from "k6";

import {
    BASE_URL,
    CLIENT_ID,
} from "../lib/config.js";

import {
    LIMIT,
    TOTAL_REQUESTS,
    WINDOW_SIZE,
    WINDOW_WAIT_SECONDS,
} from "../lib/slidingWindowConfig.js";

export default function () {

    let allowed = 0;
    let denied = 0;

    for (let i = 0; i < TOTAL_REQUESTS; i++) {

        const response = http.get(
            `${BASE_URL}/check?client_id=${CLIENT_ID}`
        );

        if (response.error) {
            console.log("Transport Error:", response.error);
        }

        if (response.status !== 200 && response.status !== 429) {
            console.log("Unexpected Status:", response.status);
        }

        const body = response.json();

        if (body.allowed) {
            allowed++;
        } else {
            denied++;
        }

    }

    check(null, {
        "allowed requests equal limit": () =>
            allowed === LIMIT,

        "exactly one request denied": () =>
            denied === TOTAL_REQUESTS - LIMIT,
    });

    console.log("");

    console.log("==================================");
    console.log("Sliding Window Capacity Validation");
    console.log("==================================");

    console.log(`Expected Allowed : ${LIMIT}`);
    console.log(`Actual Allowed   : ${allowed}`);

    console.log(
        `Expected Denied  : ${TOTAL_REQUESTS - LIMIT}`
    );
    console.log(`Actual Denied    : ${denied}`);

    console.log("==================================");
    console.log("");

    console.log("");
    console.log("Starting expiration validation...");
    console.log("");

    sleep(WINDOW_WAIT_SECONDS);

    let expirationAllowed = 0;
    let expirationDenied = 0;

    for (let i = 0; i < TOTAL_REQUESTS; i++) {

        const response = http.get(
            `${BASE_URL}/check?client_id=${CLIENT_ID}`
        );

        const body = response.json();

        if (body.allowed) {
            expirationAllowed++;
        } else {
            expirationDenied++;
        }

    }

    check(null, {
        "expiration allowed requests equal limit": () =>
            expirationAllowed === LIMIT,

        "expiration denied exactly one request": () =>
            expirationDenied === TOTAL_REQUESTS - LIMIT,
    });


    console.log("");

    console.log("==================================");
    console.log("Sliding Window Expiration Validation");
    console.log("==================================");

    console.log(`Expected Allowed : ${LIMIT}`);
    console.log(`Actual Allowed   : ${expirationAllowed}`);

    console.log(
        `Expected Denied  : ${TOTAL_REQUESTS - LIMIT}`
    );
    console.log(`Actual Denied    : ${expirationDenied}`);

    console.log("==================================");
    console.log("");



}