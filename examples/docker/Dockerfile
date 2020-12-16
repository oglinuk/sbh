FROM golang:1.15.0
ADD . /go/src/sbh
WORKDIR /go/src/sbh
RUN go get
RUN go build -o sbh .
EXPOSE 9001