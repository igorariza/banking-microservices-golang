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
