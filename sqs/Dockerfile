# Estágio de compilação
FROM golang:alpine AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia o código fonte para o diretório de trabalho no container
COPY . .

# Compila o código fonte
RUN go build -o main .

# Estágio de produção
FROM alpine:latest

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia o binário compilado do estágio anterior
COPY --from=builder /app/main .

# Comando padrão para executar a aplicação
CMD ["./main"]
