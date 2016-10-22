## Prod/Staging

FROM alpine:latest

ENV GOPATH /go

RUN apk add --update git go && \
    git clone https://github.com/MarcosSegovia/MyWins.git /go/src/github.com/MarcosSegovia/MyWins &&\
    cd /go/src/github.com/MarcosSegovia/MyWins/src &&\
    go get && go build -o /mywins &&\
    apk del go git &&\
    mv /go/src/github.com/MarcosSegovia/MyWins/files /files &&\
    rm -rf /go

EXPOSE 8080
CMD ["/mywins"]

## Local Development

#FROM alpine:latest
#
#ENV GOPATH /go
#
#COPY . /go/src/github.com/MarcosSegovia/MyWins
#
#RUN apk add --update git go && \
#    cd /go/src/github.com/MarcosSegovia/MyWins/src &&\
#    go get github.com/gorilla/mux && go build -o /mywins *.go &&\
#    apk del go git &&\
#    mv /go/src/github.com/MarcosSegovia/MyWins/files /files &&\
#    rm -rf /go
#
#EXPOSE 8080
#CMD ["/mywins"]
