FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o stress_test main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/stress_test .
ENTRYPOINT ["./stress_test"]
