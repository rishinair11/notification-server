FROM alpine:latest

RUN apk add --update ca-certificates

ADD build/bin/notification-server-go /usr/local/bin/notification-server-go

EXPOSE 5252

ENTRYPOINT ["/usr/local/bin/notification-server-go"]