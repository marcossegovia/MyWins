FROM golang:1.7-alpine

RUN apk add --update git && rm -rf /var/cache/apk/*

#Dependencies
RUN go get github.com/gorilla/mux

COPY . $GOPATH

RUN go get -d -v github.com/MarcosSegovia/MyWins/src
RUN go install github.com/MarcosSegovia/MyWins/src

WORKDIR $GOPATH/src/github.com/MarcosSegovia/MyWins
RUN go build -o mywins src/*.go

CMD $GOPATH/mywins

EXPOSE 8080


