FROM golang:1.24.4-alpine3.22 AS builder

COPY . .

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /rest-api ./cmd/rest
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /rest-api ./cmd/rest

FROM scratch

COPY --from=builder ./rest-api /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# for run in local
# COPY .env .

ENTRYPOINT ["/rest-api"]