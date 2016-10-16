FROM golang:1.7-alpine

RUN apk add --update git && rm -rf /var/cache/apk/*
COPY . $GOPATH
RUN go get github.com/gorilla/mux && \
    go build src/*.go
CMD $GOPATH/main

EXPOSE 8080


