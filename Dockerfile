FROM golang:1.19.13-alpine3.18 AS builder
RUN apk add build-base
WORKDIR /build
COPY ./ .
RUN go build -tags netgo -ldflags="-w -s" -o bing-bong main.go

FROM alpine
LABEL MAINTAINER=github.com/amtoaer/bing-bong
WORKDIR /data
RUN apk add -U tzdata
COPY --from=builder /build/bing-bong /usr/bin/bing-bong
RUN chmod +x /usr/bin/bing-bong
VOLUME /data
ENTRYPOINT ["/usr/bin/bing-bong"]