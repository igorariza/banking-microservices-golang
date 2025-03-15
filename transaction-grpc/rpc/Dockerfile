FROM --platform=linux/amd64 golang:1.22-alpine AS builder

WORKDIR "/app"
RUN apk add --no-cache git ca-certificates

ENV go env -w GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go mod tidy  && go build -ldflags="-s -w" -o . rpc/transactionapi.go


FROM --platform=linux/amd64 alpine:3.16

RUN apk update upgrade
RUN apk --no-cache add ca-certificates bash

WORKDIR /root/

COPY --from=builder /app .
RUN chmod +x transaction

EXPOSE $PORT
CMD ["./transaction", "-f", "rpc/etc/transactionapi.yaml"]
