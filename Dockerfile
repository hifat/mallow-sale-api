FROM golang:1.25.0-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /rest ./cmd/rest

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /rest /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER nonroot

ENTRYPOINT ["/rest"]