FROM golang:alpine as builder

WORKDIR /gateway-build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o gateway main.go

FROM alpine:latest

RUN apk add ca-certificates

WORKDIR /gateway

COPY --from=builder /gateway-build/gateway .

EXPOSE 10000
CMD [ "./gateway" ]
