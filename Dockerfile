FROM golang:1.22-alpine

# Agregar dependencias necesarias para compilar
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

# Asegurar que las variables estén disponibles
ENV GIN_MODE=release

CMD ["./main"] 