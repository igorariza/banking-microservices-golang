FROM --platform=linux/amd64 golang:1.20.4-alpine3.16 AS builder

WORKDIR "/app"
RUN apk add --no-cache git ca-certificates

ENV go env -w GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOPRIVATE=banking-system/*

COPY . .
# Only for testing purposes
COPY .envrc .envrc
# Only for testing purposes 
RUN source .envrc
RUN echo "go mod tidy -compat=1.20"
RUN go mod tidy && go build

FROM --platform=linux/amd64 alpine:3.16

RUN apk update upgrade
RUN apk --no-cache add ca-certificates bash

WORKDIR /root/

COPY --from=builder /app .
RUN chmod +x account-service

EXPOSE $PORT

ENTRYPOINT ["./account-service"]
