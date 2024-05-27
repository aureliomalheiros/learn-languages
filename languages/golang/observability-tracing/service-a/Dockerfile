FROM golang:1.21-alpine

# Instala o Git
RUN apk update && apk add --no-cache git

WORKDIR /app

# Copia os arquivos go.mod e go.sum primeiro para aproveitar o cache do Docker
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante dos arquivos do projeto
COPY . .

RUN go mod tidy
RUN go build -o service-a

CMD ["./service-a"]
