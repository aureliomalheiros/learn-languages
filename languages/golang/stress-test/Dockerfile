FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o loadtest main.go

ENTRYPOINT ["./loadtest"]
