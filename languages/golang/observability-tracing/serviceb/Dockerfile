FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o serviceb .

EXPOSE 8081

CMD ["./serviceb"]
