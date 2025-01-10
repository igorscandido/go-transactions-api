FROM golang:1.23-alpine

RUN apk add --no-cache git
RUN go install github.com/air-verse/air@latest

WORKDIR /app

# Rodar o air para a API
CMD ["air", "-c", ".air.toml"]
