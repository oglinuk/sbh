FROM golang:1.17.5
WORKDIR /go/src/sbh
COPY . .
RUN go mod download
RUN go build -o sbh
EXPOSE 9001
