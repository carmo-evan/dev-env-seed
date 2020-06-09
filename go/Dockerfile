FROM golang:alpine

WORKDIR /root

COPY ./ /root

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

CMD ["./start.sh"]