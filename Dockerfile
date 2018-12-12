FROM golang:latest

RUN mkdir -p /go/src/sandbox-api
WORKDIR /go/src/sandbox-api

ADD . /go/src/sandbox-api
RUN go build

CMD ["./sandbox-api"]