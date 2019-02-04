FROM golang:1.11 as builder
COPY . $GOPATH/src/github.com/filatovw/Wattx-challenge-top-coins/
WORKDIR $GOPATH/src/github.com/filatovw/Wattx-challenge-top-coins/
RUN go install -v ./...
RUN go build -o /usr/bin/app ./api


FROM filatovw/services-base:latest
WORKDIR /usr/bin/
COPY --from=builder /usr/bin/app .
COPY ./api/docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]