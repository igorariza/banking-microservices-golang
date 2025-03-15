FROM --platform=linux/amd64 golang:1.24.1-alpine3.20 AS builder

WORKDIR "/app"
RUN apk add --no-cache git ca-certificates

ENV go env -w GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -ldflags="-s -w" -o transactionapi rpc/transactionapi.go



FROM --platform=linux/amd64 alpine:3.16

RUN apk update upgrade
RUN apk --no-cache add ca-certificates bash

WORKDIR /root/

COPY --from=builder /app .
RUN chmod +x transactionapi

EXPOSE $PORT
CMD ["./transactionapi", "-f", "etc/transactionapi.yaml"]
# FROM --platform=linux/amd64 golang:1.22-alpine AS builder

# WORKDIR "/app"
# RUN apk add --no-cache git ca-certificates

# ENV go env -w GO111MODULE=on
# ENV CGO_ENABLED=0
# ENV GOOS=linux
# ENV GOARCH=amd64

# # Copiar go.mod y go.sum primero para aprovechar la caché de Docker
# COPY ./go.mod .
# COPY ./go.sum .

# # Descargar dependencias
# RUN go mod download

# # Copiar el resto del código fuente
# COPY . .

# # Construir la aplicación
# RUN go build -ldflags="-s -w" -o transaction rpc/transactionapi.go


# FROM --platform=linux/amd64 alpine:3.16

# RUN apk update upgrade
# RUN apk --no-cache add ca-certificates bash

# WORKDIR /root/

# COPY --from=builder /app .
# RUN chmod +x transaction

# EXPOSE $PORT
# CMD ["./transaction", "-f", "rpc/etc/transactionapi.yaml"]
