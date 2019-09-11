FROM golang:1.10.0
ADD . /go/src/go-sbh
WORKDIR /go/src/go-sbh
RUN go get
RUN go build -o go-sbh .
CMD ["./go-sbh"]
