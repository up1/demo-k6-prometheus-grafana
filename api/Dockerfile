FROM golang:1.24-alpine AS step1
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o /app/main main.go

FROM alpine:3.21
COPY --from=step1 /app/main /api/main
RUN apk add --no-cache curl
EXPOSE 8080
CMD ["/api/main"]
