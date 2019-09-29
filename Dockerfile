FROM golang:alpine
WORKDIR /app
COPY . .
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh && \
    go get -u gopkg.in/gomail.v2 && \
    go build src/server.go
EXPOSE 5252
CMD ["./server"]