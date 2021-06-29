# Start from the latest golang base image
FROM golang:latest as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY snowflake_server.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .


FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
