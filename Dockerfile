# Builder
FROM golang:1.23.2-alpine3.20 as builder

RUN apk update && apk upgrade && \
    apk --update add git make bash build-base

WORKDIR /app

COPY . .

RUN make build

# Distribution

FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 8080

COPY --from=builder /app/bin /app/

CMD /app/case-study/library
