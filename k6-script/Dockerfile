FROM golang:1.24-alpine AS builder
WORKDIR $GOPATH/src/go.k6.io/k6
RUN apk --no-cache add git=~2
RUN CGO_ENABLED=0 go install go.k6.io/xk6/cmd/xk6@latest  \
    && CGO_ENABLED=0 xk6 build \
    --with github.com/grafana/xk6-output-prometheus-remote \
    --output /tmp/k6

# Create image for running k6 with output for Prometheus remote write
FROM alpine:3.21
RUN apk add --no-cache ca-certificates && \
    adduser -D -u 12345 -g 12345 k6
COPY --from=builder /tmp/k6 /usr/bin/k6
USER 12345
WORKDIR /home/k6
COPY . .
ENTRYPOINT ["k6"]