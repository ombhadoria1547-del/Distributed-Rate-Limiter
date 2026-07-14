import http from "k6/http";

import {
    BASE_URL,
    CLIENT_ID,
} from "../lib/config.js";

const DEBUG = __ENV.DEBUG === "true";

export const options = {
    vus: Number(__ENV.VUS) || 10,

    duration: __ENV.DURATION || "10s",
};

export default function () {
    const response = http.get(
        `${BASE_URL}/check?client_id=${CLIENT_ID}`
    );

    if (DEBUG && response.status === 429) {
        console.log("429 received");
    }

    if (DEBUG && response.error) {
        console.log("Transport Error:", response.error);
    }

    if (
        DEBUG &&
        response.status !== 200 &&
        response.status !== 429
    ) {
        console.log(
            "Unexpected Status:",
            response.status
        );
    }
}