FROM golang:1.14.2
COPY . $GOPATH
WORKDIR $GOPATH/src/sum
RUN go build

CMD ["./sum"]