# Imagem base
FROM golang:latest

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia o código fonte para o diretório de trabalho no container
COPY . .

# Compila o código fonte
RUN go build -o main .

# Comando padrão para executar a aplicação
CMD ["./main"]
