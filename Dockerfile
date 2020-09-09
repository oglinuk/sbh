FROM golang:1.15.0
ADD . /go/src/go-sbh
WORKDIR /go/src/go-sbh
RUN go get
RUN go build -o go-sbh .
EXPOSE 9001
CMD ["./go-sbh"]