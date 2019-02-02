FROM golang:1.11 as builder

COPY vendor $GOPATH/src/github.com/filatovw/app/vendor
COPY api $GOPATH/src/github.com/filatovw/app/api

WORKDIR $GOPATH/src/github.com/filatovw/app/

RUN go install -v ./...
RUN go build -o /usr/bin/app ./app

FROM golang:1.11
WORKDIR /usr/bin/

COPY --from=builder /usr/bin/app .
COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]

EXPOSE 8080