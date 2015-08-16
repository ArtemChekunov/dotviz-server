FROM golang
MAINTAINER Chekunov Artem <scorp.dev.null@gmail.com>


ADD . /go/src/github.com/sc0rp1us/dotviz-server
WORKDIR /go/src/github.com/sc0rp1us/dotviz-server

RUN apt-get update && apt-get install -y graphviz
RUN go get github.com/sc0rp1us/dotviz-server
RUN go install github.com/sc0rp1us/dotviz-server

ENTRYPOINT /go/bin/dotviz-server

EXPOSE 1234
