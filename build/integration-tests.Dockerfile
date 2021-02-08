FROM golang:1.15.2 as build

WORKDIR /go/src

COPY . ${CODE_DIR}

RUN go test -i ./tests/integration/...

CMD go test -v ./tests/integration/...
