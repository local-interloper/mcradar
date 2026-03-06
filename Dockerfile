# Builder
FROM golang:1.26 AS builder

WORKDIR /workspace

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./mc-radar

# Runner
FROM alpine:latest AS runner

WORKDIR /app

COPY --from=builder /workspace/mc-radar .

CMD ["./mc-radar"]