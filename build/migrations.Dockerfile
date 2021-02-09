FROM golang:1.15.2

WORKDIR "/migrations"

COPY ./migrations .
COPY ./scripts/wait-for-it.sh .

RUN go get -u github.com/pressly/goose/cmd/goose
RUN chmod +x "./wait-for-it.sh"

ENTRYPOINT ["./wait-for-it.sh"]
