FROM alpine:latest
MAINTAINER Marcos Segovia <velozmarkdrea@gmail.com>

ENV GOPATH /go

ENV app_env prod
ENV DB_DBNAME mywins
ENV DB_WINS_COLLECTION wins
ENV DB_FAILS_COLLECTION fails

COPY . /go/src/github.com/MarcosSegovia/MyWins

RUN apk add --update git go make &&\
    go get github.com/Masterminds/glide &&\
    cd /go/src/github.com/Masterminds/glide &&\
    make install &&\
    cd /go/src/github.com/MarcosSegovia/MyWins &&\
    glide install &&\
    cd /go/src/github.com/MarcosSegovia/MyWins/src &&\
    go build -o /mywins *.go &&\
    mv /go/src/github.com/MarcosSegovia/MyWins/files /files &&\
    apk del go git &&\
    rm -rf /go

EXPOSE 8080
EXPOSE 8081
CMD ["/mywins"]
