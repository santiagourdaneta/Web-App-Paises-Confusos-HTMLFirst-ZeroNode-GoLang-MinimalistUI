# ETAPA 1: Construcción
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
# Compilamos un binario estático (sin dependencias externas)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# ETAPA 2: Ejecución (Ultra ligera)
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/index.html .
# Si tienes carpeta static, quita el comentario de abajo:
# COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./main"]
