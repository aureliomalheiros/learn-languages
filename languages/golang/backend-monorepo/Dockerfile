FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:3.21.3

WORKDIR /app

COPY --from=builder /app/main .

CMD ["/app/main"]
