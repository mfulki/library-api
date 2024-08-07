FROM golang:1.22.3 AS builder
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/app


FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y ca-certificates


COPY --from=builder /app/main /app/main


ENV GIN_MODE=release

EXPOSE 8080

CMD ["/app/main"]