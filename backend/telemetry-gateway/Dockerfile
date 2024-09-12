FROM golang:1.22 AS builder
ENV GO111MODULE=on
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o telemetry-gateway

# Stage save binary in lightweight docker image
FROM alpine:3.20
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/telemetry-gateway /app/telemetry-gateway
RUN chmod +x /app/telemetry-gateway

# Run the binary
ENTRYPOINT ["/app/telemetry-gateway"]