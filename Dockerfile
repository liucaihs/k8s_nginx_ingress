FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/eddycjy/awesome-project
COPY . $GOPATH/src/github.com/eddycjy/awesome-project
RUN go build .

EXPOSE 9001
ENTRYPOINT ["./awesome-project"]