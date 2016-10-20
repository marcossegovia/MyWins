FROM alpine:latest

ENV GOPATH /go

RUN apk add --update git go && \
    git clone https://github.com/MarcosSegovia/MyWins.git /go/src/github.com/MarcosSegovia/MyWins &&\
    cd /go/src/github.com/MarcosSegovia/MyWins/src &&\
    go get && go build -o /mywins &&\
    apk del go git &&\
    rm -fr /go

EXPOSE 8080
CMD ["/mywins"]
