package config

import (
	"fmt"
	"strconv"
)

type listener struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (l *listener) GetAddr() string {
	var (
		host = l.Host
		port = strconv.Itoa(l.Port)
	)
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "8080"
	}
	return fmt.Sprintf("%s:%s", host, port)
}

type Postgres struct {
	listener
	DB       string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type HTTP struct {
	listener
}

type GRPC struct {
	listener
}

func (g GRPC) FullServiceName(service string) string {
	return fmt.Sprintf("consul://%v/services-%v-%v", DefaultConsul, service, g.Port)
}
