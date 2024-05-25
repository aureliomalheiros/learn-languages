# Usar a imagem base do Go
FROM golang:1.21-alpine

# Instalar dependências necessárias
RUN apk add --no-cache gcc musl-dev sqlite

# Criar diretório de trabalho
WORKDIR /app

# Copiar o código do servidor para o diretório de trabalho
COPY . .

# Compilar o código Go
RUN go build -o server server.go

# Expor a porta que o servidor vai usar
EXPOSE 8080

# Comando para rodar o servidor
CMD ["./server"]
