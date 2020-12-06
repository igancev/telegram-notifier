# docker build -t igancev/telegram-notifier:latest -f Dockerfile .
# docker push igancev/telegram-notifier:latest

FROM golang:1.15-alpine3.12 as golang

WORKDIR /go
COPY . .
RUN go build -o telegram-notifier .

FROM alpine:3.12
COPY --from=golang /go/telegram-notifier /usr/local/bin/telegram-notifier

CMD ["/usr/local/bin/telegram-notifier"]
