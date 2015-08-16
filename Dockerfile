FROM golang
MAINTAINER Chekunov Artem <scorp.dev.null@gmail.com>


ADD . /go/src/github.com/sc0rp1us/dotviz-server
ADD entrypoint.sh /tmp/entrypoint.sh
WORKDIR /go/src/github.com/sc0rp1us/dotviz-server

RUN apt-get update && apt-get install -y graphviz
RUN go get github.com/sc0rp1us/dotviz-server
RUN go install github.com/sc0rp1us/dotviz-server

ENTRYPOINT /tmp/entrypoint.sh

EXPOSE 1234
