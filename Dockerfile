## Prod/Staging

FROM alpine:latest

ENV GOPATH /go

RUN apk add --update git go make && \
    git clone https://github.com/MarcosSegovia/MyWins.git /go/src/github.com/MarcosSegovia/MyWins &&\
    go get github.com/Masterminds/glide &&\
    cd /go/src/github.com/Masterminds/glide &&\
    make install &&\
    cd /go/src/github.com/MarcosSegovia/MyWins &&\
    glide install &&\
    cd /go/src/github.com/MarcosSegovia/MyWins/src &&\
    go build -o /mywins *.go &&\
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
#RUN apk add --update git go make &&\
#    go get github.com/Masterminds/glide &&\
#    cd /go/src/github.com/Masterminds/glide &&\
#    make install &&\
#    cd /go/src/github.com/MarcosSegovia/MyWins &&\
#    glide install &&\
#    cd /go/src/github.com/MarcosSegovia/MyWins/src &&\
#    go build -o /mywins *.go &&\
#    mv /go/src/github.com/MarcosSegovia/MyWins/files /files &&\
#    apk del go git &&\
#    rm -rf /go
#
#EXPOSE 8080
#CMD ["/mywins"]
