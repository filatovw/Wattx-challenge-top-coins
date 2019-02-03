FROM golang:1.11

ENV PROTOBUF_ENV linux-x86_64
ENV PROTOBUF_VERSION 3.6.1

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates tzdata git unzip make curl && \
    rm -rf /var/lib/apt/lists/*

RUN go get github.com/golang/protobuf/proto \
    github.com/golang/protobuf/protoc-gen-go \
    google.golang.org/grpc

# Protobuf
ADD https://github.com/google/protobuf/releases/download/v$PROTOBUF_VERSION/protoc-$PROTOBUF_VERSION-$PROTOBUF_ENV.zip /tmp/
RUN unzip /tmp/protoc-$PROTOBUF_VERSION-$PROTOBUF_ENV.zip -d /usr/local/ && rm /tmp/protoc-$PROTOBUF_VERSION-$PROTOBUF_ENV.zip
