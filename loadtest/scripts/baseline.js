import http from "k6/http";
import { check } from "k6";

export const options = {
    vus: 1,
    iterations: 1,
};

export default function () {
    const response = http.get(
        "http://localhost:8080/check?client_id=baseline-client"
    );

    check(response, {
        "status is 200": (r) => r.status === 200,

        "response contains allowed": (r) =>
            JSON.parse(r.body).allowed !== undefined,

        "response contains remaining": (r) =>
            JSON.parse(r.body).remaining !== undefined,

        "response contains retry_after": (r) =>
            JSON.parse(r.body).retry_after !== undefined,
    });
}