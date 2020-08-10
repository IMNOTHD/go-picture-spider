FROM golang:latest

MAINTAINER imnothd "ideshenghe@gmail.com"

WORKDIR $GOPATH/src/github.com/IMNOTHD/go-picture-spider
COPY . $GOPATH/src/github.com/IMNOTHD/go-picture-spider

# RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build .

EXPOSE 80

ENTRYPOINT ["./go-picture-spider"]