FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN apk update && apk upgrade && apk add --no-cache bash git openssh && go get -u gopkg.in/gomail.v2 && go build src/server.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 5252
CMD ["./server"]