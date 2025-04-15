FROM golang:1.23.5 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o api-benchmark main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/api-benchmark .

# Copia arquivos usados no runtime, como data/results.json
COPY --from=builder /app/data ./data

# Expõe a porta usada no backend
EXPOSE 8080

# Comando para rodar o backend
CMD ["./api-benchmark"]
