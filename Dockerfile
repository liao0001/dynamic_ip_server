FROM golang:1.15.1-alpine3.12 AS be-builder
ENV GO111MODULE on

ADD . /app

RUN GOOS=linux go build -mod=vendor -a -o serv main.go

 docker run --rm -v /d/workspace/dev/src/github.com/liao0001/dynamic_ip_server:/usr/src/myapp -w /d/workspace/dev/src/github.com/liao0001/dynamic_ip_server -e GOOS=windows -e GOARCH=386 golang:1.15.1-alpine3.12 go build -v
 docker run --rm -v "$PWD":/usr/src/myapp -w "$PWD" -e GOOS=windows -e GOARCH=386 golang:latest go build -v