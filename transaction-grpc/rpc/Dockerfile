FROM --platform=linux/amd64 golang:1.20-alpine AS builder

WORKDIR "/app"
RUN apk add --no-cache git ca-certificates

ENV go env -w GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go mod tidy -compat=1.20 && go build -ldflags="-s -w" -o . rpc/transaction.go


FROM --platform=linux/amd64 alpine:3.16

RUN apk update upgrade
RUN apk --no-cache add ca-certificates bash

WORKDIR /root/

COPY --from=builder /app .
RUN chmod +x transaction

EXPOSE $PORT
CMD ["./transaction", "-f", "rpc/etc/transaction.yaml"]