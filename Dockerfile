FROM golang:latest

ENV GO111MODULE on

RUN apt-get update -qq && \
  apt-get -q -y install mysql-client

RUN mkdir -p /go/src/sandbox-api
WORKDIR /go/src/sandbox-api

ADD . /go/src/sandbox-api
RUN go get -u
RUN go build

CMD ["./sandbox-api"]