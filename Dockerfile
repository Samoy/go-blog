FROM golang:latest

WORKDIR $GOPATH/src/github.com/samoy/go-blog
COPY . $GOPATH/src/github.com/samoy/go-blog
RUN go build .
EXPOSE 8000
ENTRYPOINT ["./go-blog"]