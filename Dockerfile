FROM golang:1.24.4-alpine3.22 AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /api ./cmd/api

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /api /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Run as non-root user by default (UID=65532)
USER nonroot

ENTRYPOINT ["/api"]
