FROM golang:latest

WORKDIR /go/src/github.com/Dimitriy14/commuticators/grpc/client

ADD . .

RUN go get google.golang.org/grpc

CMD ["go", "run", "."]
