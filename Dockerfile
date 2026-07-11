FROM golang:1.26.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o rate-limiter

EXPOSE 8080

CMD ["./rate-limiter"]